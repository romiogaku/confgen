package backends

import (
	"errors"

	"github.com/romiogaku/confgen/pkg/backends/consul"
	"github.com/romiogaku/confgen/pkg/backends/etcdv3"
	"github.com/romiogaku/confgen/pkg/backends/zookeeper"
)

// BackendClient interface
type BackendClient interface {
	GetValue(key string) (string, error)
	Close()
}

// New backend client
func New(config Config) (BackendClient, error) {
	switch config.Backend {
	case "etcd":
		return etcdv3.NewClient(config.BackendNodes)
	case "consul":
		return consul.NewClient(config.BackendNodes)
	case "zookeeper":
		return zookeeper.NewClient(config.BackendNodes)
	}
	return nil, errors.New("Invalid backend")
}
