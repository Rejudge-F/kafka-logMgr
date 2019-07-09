package etcd

import (
	"context"
	"fmt"
	"github.com/astaxie/beego/logs"
	etcd_client "go.etcd.io/etcd/clientv3"
	"time"
)

type EtcdClient struct {
	client *etcd_client.Client
}

var (
	etcdClient *EtcdClient
)

func InitEtcd(addr []string, key string) error {
	cli, err := etcd_client.New(etcd_client.Config{
		Endpoints:   addr,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		logs.Error(fmt.Sprintf("%s%s", "InitEtcd Failed", err))
		panic("Init Etcd Failed")
		return err
	}

	etcdClient = &EtcdClient{
		client: cli,
	}
	return nil
}

func GetKey(key string) (s string, err error) {
	s = ""
	err = nil
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := etcdClient.client.Get(ctx, key)
	if err != nil {
		logs.Error(fmt.Sprint("GetKey Failed: %v", err))
		return
	}
	kvs := resp.Kvs
	cancel()
	for _, v := range kvs {
		if string(v.Key) == key {
			return string(v.Value), nil
		}
	}
	return
}
