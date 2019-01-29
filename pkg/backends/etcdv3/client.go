package etcdv3

import (
	"context"
	"time"

	"go.etcd.io/etcd/clientv3"
)

// Client confgen etcdv3 client
type Client struct {
	client *clientv3.Client
}

// NewClient return confgen etcdv3 client
func NewClient(machines []string) (*Client, error) {
	cfg := clientv3.Config{
		Endpoints:            machines,
		DialTimeout:          5 * time.Second,
		DialKeepAliveTime:    10 * time.Second,
		DialKeepAliveTimeout: 3 * time.Second,
	}
	cli, err := clientv3.New(cfg)
	if err != nil {
		return &Client{}, err
	}
	return &Client{cli}, nil
}

// GetValue return value
func (c *Client) GetValue(key string) (string, error) {
	value := ""
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := c.client.Get(ctx, key)
	cancel()
	if err != nil {
		return value, err
	}
	for _, ev := range resp.Kvs {
		value = string(ev.Value)
	}
	// TOTO: prefix取得になってないか確認
	return value, nil
}

// Close etcdv3 client
func (c *Client) Close() {
	c.client.Close()
}
