package main

import (
	"bufio"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

const READ_BUF_SIZE = 1024

// 读取配置文件后的回调
func onConfig() {
	// 日志
	if Cfg.Debug {
		log.SetFlags(log.Lmicroseconds | log.Ldate | log.Lshortfile)
	} else {
		log.SetFlags(0)
	}
}

func main() {
	// 读取、设定证书
	ca_cert, err := ioutil.ReadFile(Cfg.Certificate)
	if err != nil {
		log.Panicf("Read file ERROR: %v", err)
	}
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(ca_cert)
	// cert, err := tls.LoadX509KeyPair(Cfg.Certificate, "")
	// if err != nil {
	// 	log.Panicf("Read file ERROR: %v\n", err)
	// }

	// 连接
	url := Cfg.Host + ":" + strconv.Itoa(Cfg.Port)
	log.Printf("Connecting to: %v", url)
	conn, err := tls.Dial("tcp", url, &tls.Config{
		RootCAs: pool,
	})
	if err != nil {
		log.Panicf("Connecting ERROR: %v", err)
	}
	log.Printf("Connected SUCCESS")

	// 交互
	reader := bufio.NewReader(os.Stdin)
	output := make([]byte, READ_BUF_SIZE)
	for {
		fmt.Printf(Cfg.Prompt)
		input, _ := reader.ReadString('\n')
		conn.Write([]byte(input))
		num, err := conn.Read(output)
		if err != nil {
			log.Panicf("Read ERROR: %v\n", err)
		}
		fmt.Printf("%v", string(output[:num]))
	}

}
