package etcd

import (
	"context"
	"fmt"
	"sync"
	"time"

	"go.etcd.io/etcd/clientv3"
)

var mutex sync.Mutex

func Insert(cli *clientv3.Client, key string, value string) error {
	mutex.Lock()
	defer mutex.Unlock()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err := cli.Put(ctx, key, value)
	cancel()
	if err != nil {
		fmt.Printf("insert to etcd failed, err:%v\n", err)
		return err
	}
	return nil
}

func Update(cli *clientv3.Client, record string) error {
	mutex.Lock()
	defer mutex.Unlock()
	return nil
}

func SearchByPrefix(cli *clientv3.Client, prefix string) ([]string, error) {
	mutex.Lock()
	defer mutex.Unlock()
	return nil, nil
}

func Search(cli *clientv3.Client, key string) (string, error) {
	mutex.Lock()
	defer mutex.Unlock()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, key)
	if err != nil {
		fmt.Printf("search from etcd failed, err:%v\n", err)
		return "", err
	}
	//for _, ev := range resp.Kvs {
	//	fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	//}
	cancel()
	if len(resp.Kvs) > 0 {
		return string(resp.Kvs[0].Value), nil
	}
	return "", nil
}

func Delete(cli *clientv3.Client, record string) error {
	mutex.Lock()
	defer mutex.Unlock()
	return nil
}
