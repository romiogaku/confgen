package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"text/template"
	"time"

	"go.etcd.io/etcd/clientv3"
)

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

// Config confgen config
type Config struct {
	TemplateFilePath string
	BackendNodes     Nodes
}

var config Config

func init() {
	flag.StringVar(&config.TemplateFilePath, "file", "", "template file path")
	flag.Var(&config.BackendNodes, "node", "list of backend nodes")

	flag.Parse()

	// validation
	if config.TemplateFilePath == "" {
		fmt.Fprintln(os.Stderr, "file parameter required.")
		os.Exit(1)
	}
	if len(config.BackendNodes) == 0 {
		fmt.Fprintln(os.Stderr, "node parameter required.")
		os.Exit(1)
	}
}

func getValue(key string) string {
	cfg := clientv3.Config{
		Endpoints:   config.BackendNodes,
		DialTimeout: 5 * time.Second,
	}

	cli, err := clientv3.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		log.Fatal(err)
	}

	value := ""
	for _, ev := range resp.Kvs {
		value = string(ev.Value)
	}
	return value
}

func main() {
	funcMap := template.FuncMap{
		"v": getValue,
	}
	tmpl := template.New(path.Base(config.TemplateFilePath)).Funcs(funcMap)
	tmpl, err := tmpl.ParseFiles(config.TemplateFilePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
