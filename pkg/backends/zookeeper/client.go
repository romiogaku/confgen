package zookeeper

import (
	"errors"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

// Client confgen zookeeper client
type Client struct {
	client *zk.Conn
}

// NewClient return confgen etcdv3 client
func NewClient(machines []string) (*Client, error) {
	cli, _, err := zk.Connect(machines, time.Second, zk.WithLogInfo(false)) //*10)
	if err != nil {
		return &Client{}, err
	}
	return &Client{cli}, nil
}

// GetValue return value
func (c *Client) GetValue(key string) (string, error) {
	exist, _, err := c.client.Exists(key)
	if err != nil {
		return "", err
	}
	if !exist {
		return "", errors.New("key: " + key + " is not exist.")
	}
	v, _, err := c.client.Get(key)
	return string(v), err
}

// Close etcdv3 client
func (c *Client) Close() {
	c.client.Close()
}
