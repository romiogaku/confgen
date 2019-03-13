package backends

import (
	"fmt"
)

// Config confgen config
type Config struct {
	TemplateFilePath string
	Backend          string
	BackendNodes     Nodes
	Key              string
}

// Nodes is a custom flag Var representing a list of etcd nodes.
type Nodes []string

// String returns the string representation of a node var.
func (n *Nodes) String() string {
	return fmt.Sprintf("%s", *n)
}

// Set appends the node to the etcd node list.
func (n *Nodes) Set(node string) error {
	*n = append(*n, node)
	return nil
}
