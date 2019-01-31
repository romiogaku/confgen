package consul

import (
	"github.com/hashicorp/consul/api"
)

// Client confgen consul client
type Client struct {
	client *api.KV
}

// NewClient return confgen consul client
func NewClient(machines []string) (*Client, error) {
	conf := api.DefaultConfig()
	if len(machines) > 0 {
		conf.Address = machines[0]
	}
	client, err := api.NewClient(conf)
	if err != nil {
		return nil, err
	}
	return &Client{client.KV()}, nil
}

// GetValue return value
func (c *Client) GetValue(key string) (string, error) {
	kv, _, err := c.client.Get(key, nil)
	return string(kv.Value), err
}

// Close etcdv3 client
func (c *Client) Close() {
	//
}
