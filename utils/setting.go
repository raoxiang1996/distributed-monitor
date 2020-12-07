package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	ip   string
	port int
)

func init() {
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Printf("配置文件读取错误，请检查文件路径", err)
	}
	LoadServer(cfg)
}

func LoadServer(file *ini.File) {
	ip = file.Section("server").Key("ip").MustString("127.0.0.1")
	port = file.Section("server").Key("port").MustInt(2380)
}
