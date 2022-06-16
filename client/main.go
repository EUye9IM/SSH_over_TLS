package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"

	"golang.org/x/crypto/ssh/terminal"

	"golang.org/x/crypto/ssh"
)

// 读取配置文件后的回调
func onConfig() {
	// 日志
	if Cfg.Debug {
		log.SetFlags(log.Lshortfile)
	} else {
		log.SetFlags(0)
	}
}

func main() {
	// 读取、设定证书
	ca_cert, err := ioutil.ReadFile(Cfg.Certificate)
	if err != nil {
		log.Panicf("Read file %v FAILED: %v", Cfg.Certificate, err)
	}
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(ca_cert)

	tls_cfg := &tls.Config{
		RootCAs: pool,
	}
	if Cfg.Debug {
		tls_cfg.InsecureSkipVerify = true
	}
	// 读取用户名
	if Cfg.Account.Username == "" {
		fmt.Printf("Username: ")
		fmt.Scanln(&Cfg.Account.Username)
		fmt.Printf("Password: ")
		pwbyte, err := terminal.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			log.Panicln("read password ERROR", err)
		}
		Cfg.Account.Password = string(pwbyte)
		fmt.Println("")
	}

	// 连接
	url := Cfg.Host + ":" + strconv.Itoa(Cfg.Port)
	log.Printf("Connecting to: %v", url)
	// tls_conn, err := tls.Dial("tcp", url, tls_cfg)
	tls_dialer := tls.Dialer{Config: tls_cfg}
	net_conn, err := tls_dialer.Dial("tcp", url)
	if err != nil {
		log.Panicf("Connecting FAILED: %v", err)
	}
	log.Printf("Connected SUCCESS")
	defer net_conn.Close()

	//转换

	// 建立 ssh 客户端
	client_config := &ssh.ClientConfig{
		User: Cfg.Account.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(Cfg.Account.Password),
		},
		HostKeyCallback: func(
			hostname string,
			remote net.Addr,
			key ssh.PublicKey,
		) error {
			return nil
		},
	}
	conn, chans, reqs, err := ssh.NewClientConn(net_conn, url, client_config)
	if err != nil {
		log.Printf("Handshake FAILED: %v", err)
		return
	}
	log.Println("Login SUCCESS")
	client := ssh.NewClient(conn, chans, reqs)

	// 建立会话
	session, err := client.NewSession()
	if err != nil {
		log.Printf("Create session FAILED: %v", err)
		return
	}
	defer session.Close()

	// 建立终端
	input := newReader()
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = input
	// 交互
	log.Println(" ==== START SHELL ====")
	defer log.Println(" ===== END SHELL =====")
	session.Shell()
	session.Wait()

}
