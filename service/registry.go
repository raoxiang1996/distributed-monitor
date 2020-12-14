package service

import "distributed-monitor/etcd"

// 向etcd里注册服务
func Registry() {
	etcd.Init()
	etcd.Insert(etcd.Clictl, "rao", "xiang")
}
