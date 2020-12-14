package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	Ip   string
	Port string
)

func init() {
	cfg, err := ini.Load("../config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径", err)
	} else {
		LoadConfig(cfg)
	}
}

func LoadConfig(file *ini.File) {
	Ip = file.Section("server").Key("Ip").MustString("127.0.0.1")
	Port = file.Section("server").Key("Port").MustString("2380")
}
