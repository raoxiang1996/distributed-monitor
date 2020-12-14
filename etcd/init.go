package etcd

import (
	"context"
	"distributed-monitor/utils"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

var (
	dialTimeout    = 10 * time.Second
	requestTimeout = 4 * time.Second
	endpoints      = []string{}
)

var Clictl *clientv3.Client

func Init() {
	url := utils.Ip + ":" + utils.Port
	//url := "127.0.0.1:2380"
	endpoints = append(endpoints, url)
	clientConfig := clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: dialTimeout,
	}
	cli, err := clientv3.New(clientConfig)
	if err != nil {
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err = cli.Status(timeoutCtx, clientConfig.Endpoints[0])
	if err != nil {
		fmt.Println(err, "error checking etcd status: %v", err)
	} else {
		fmt.Println("connect to etcd success")
	}
	Clictl = cli
}
