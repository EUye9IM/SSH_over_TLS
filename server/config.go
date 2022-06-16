package main

import (
	"flag"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

const DEFAULT_CONFIG_PATH = "conf.yml"

var Cfg struct {
	Accounts    map[string]string // username: password
	Debug       bool
	Shell_path  string
	Port        int
	Log         string
	Certificate string
	Private_key string
}

// 读取参数中的配置文件路径
func init() {
	var config_path string
	flag.StringVar(&config_path, "c", DEFAULT_CONFIG_PATH, "set path to configuration file")
	flag.Parse()

	if conf_file, err := ioutil.ReadFile(config_path); err != nil {
		panic("Cannot open file: " + config_path + ".\n" + err.Error())
	} else if err := yaml.Unmarshal(conf_file, &Cfg); err != nil {
		panic("File: " + config_path + " unmarshal falied.\n" + err.Error())
	}

	onConfig()
}
