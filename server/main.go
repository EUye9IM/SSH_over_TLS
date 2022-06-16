package main

import (
	"crypto/tls"
	"errors"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strconv"

	"golang.org/x/crypto/ssh"
)

// 读取配置文件后的回调
func onConfig() {
	// 日志
	if !Cfg.Debug && Cfg.Log != "" {
		logFile, err := os.OpenFile(Cfg.Log, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic("Cannot open file: " + Cfg.Log + ".\n" + err.Error())
		}
		log.SetOutput(logFile)
		log.SetFlags(log.Lmicroseconds | log.Ldate)
	} else {
		log.SetFlags(log.Lmicroseconds | log.Ldate | log.Lshortfile)
	}
}

var SHELL_CMD []string

func main() {
	// 检查操作系统
	switch runtime.GOOS {
	case "linux":
		SHELL_CMD = []string{"bash"}
	case "windows":
		SHELL_CMD = []string{"powershell", "-NoLogo", "-NoExit", "-NonInteractive", "cd " + Cfg.Shell_path}

	}
	// 读取、设定密钥、证书
	cert, err := tls.LoadX509KeyPair(Cfg.Certificate, Cfg.Private_key)
	if err != nil {
		log.Panicf("LoadX509KeyPair FAILED: %v", err)
	}
	tls_config := &tls.Config{Certificates: []tls.Certificate{cert}}

	// 端口监听
	listener, err := tls.Listen("tcp", ":"+strconv.Itoa(Cfg.Port), tls_config)
	if err != nil {
		log.Panicf("Listen ERROR: %v", err)
		return
	}
	defer listener.Close()
	log.Printf("Listening at: %v", Cfg.Port)

	// 接受到一个新连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Accept connection ERROR: %v", err)
			continue
		}
		log.Printf("New connection from: %v", conn.RemoteAddr().String())
		go handleConn(conn)
	}
}

// 新连接处理线程
func handleConn(net_conn net.Conn) {
	defer net_conn.Close()

	// 服务安全配置
	server_config := &ssh.ServerConfig{
		PasswordCallback: func(
			c ssh.ConnMetadata, password []byte,
		) (*ssh.Permissions, error) {
			pw, ok := Cfg.Accounts[c.User()]
			if ok && pw == string(password) {
				return nil, nil
			}
			return nil, errors.New("password rejected for " + c.User())
		},
	}
	privateBytes, err := ioutil.ReadFile(Cfg.Private_key)
	if err != nil {
		log.Fatal("Failed to load private key: ", err)
	}

	private, err := ssh.ParsePrivateKey(privateBytes)
	if err != nil {
		log.Fatal("Failed to parse private key: ", err)
	}

	server_config.AddHostKey(private)

	// TLS 基础上建立连接
	conn, chans, reqs, err := ssh.NewServerConn(net_conn, server_config)
	if err != nil {
		log.Printf("Handshake FAILED: %v", err)
		return
	}
	log.Printf("Login SUCCESS: %v", conn.User())

	// 回调设置
	// The incoming Request channel must be serviced.
	go ssh.DiscardRequests(reqs)

	// Service the incoming Channel channel.
	for newChannel := range chans {
		// Channels have a type, depending on the application level
		// protocol intended. In the case of a shell, the type is
		// "session" and ServerShell may be used to present a simple
		// terminal interface.
		if newChannel.ChannelType() != "session" {
			newChannel.Reject(ssh.UnknownChannelType, "unknown channel type")
			log.Printf("Receive unknown channel type: %v", newChannel.ChannelType())
			continue
		}
		channel, requests, err := newChannel.Accept()
		if err != nil {
			log.Printf("Accept channel ERROR: %v", err)
			continue
		}

		// Sessions have out-of-band requests such as "shell",
		// "pty-req" and "env".  Here we handle only the
		// "shell" request.
		go func(in <-chan *ssh.Request) {
			for req := range in {
				req.Reply(req.Type == "shell", nil)
			}
		}(requests)

		go func() {
			defer channel.Close()
			// 打开终端
			cmd := exec.Command(SHELL_CMD[0], SHELL_CMD[1:]...)
			// 运行
			cmd.Stdin = channel
			cmd.Stdout = channel
			cmd.Stderr = channel
			log.Printf("Start shell")
			err = cmd.Run()
			if err != nil {
				log.Printf("Start command ERROR: %v", err)
				return
			}
			log.Printf("End shell")
		}()
	}
}
