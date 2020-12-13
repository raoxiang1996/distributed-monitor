package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	Ip   string
	Port int
)

func init() {
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径", err)
	}
	LoadConfig(cfg)
}

func LoadConfig(file *ini.File) {
	Ip = file.Section("server").Key("ip").MustString("127.0.0.1")
	Port = file.Section("server").Key("port").MustInt(2380)
}
