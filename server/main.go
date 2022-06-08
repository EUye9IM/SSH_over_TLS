package main

import (
	"bufio"
	"crypto/tls"
	"log"
	"net"
	"os"
	"strconv"
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

func main() {
	// 读取、设定密钥、证书
	cert, err := tls.LoadX509KeyPair(Cfg.Certificate, Cfg.Private_key)
	if err != nil {
		log.Panicf("LoadX509KeyPair ERROR: %v", err)
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

func handleConn(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)

	for {
		// Read
		msg, err := r.ReadString('\n')
		if err != nil {
			log.Printf("Read string ERROR: %v", err)
			return
		}
		log.Printf("Read SUCCESS: %v", msg)

		// response
		_, err = conn.Write([]byte(msg))
		if err != nil {
			log.Printf("Write string ERROR: %v", err)
			return
		}
		log.Printf("Write SUCCESS")
	}
}
