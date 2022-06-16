package main

import (
	"flag"
)

const DEFAULT_CER_PATH = "pem.cer"

var Cfg struct {
	Account struct {
		Username string
		Password string
	}
	Debug       bool
	Host        string
	Port        int
	Certificate string
}

// 读取参数中的配置文件路径
func init() {
	flag.StringVar(&Cfg.Account.Username, "u", "", "username.")
	flag.StringVar(&Cfg.Account.Password, "p", "", "password")
	flag.StringVar(&Cfg.Certificate, "c", DEFAULT_CER_PATH, "path to certificate file")
	flag.BoolVar(&Cfg.Debug, "d", false, "on debug mode: skipping verify and print code location")
	flag.StringVar(&Cfg.Host, "H", "127.0.0.1", "host")
	flag.IntVar(&Cfg.Port, "P", 443, "port")
	flag.Parse()

	onConfig()
}
