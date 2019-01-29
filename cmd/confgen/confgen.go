package main

import (
	"flag"
	"fmt"
	"os"

	b "github.com/romiogaku/confgen/pkg/backends"
	t "github.com/romiogaku/confgen/pkg/template"
)

var parse t.Parse
var config b.Config
var client b.BackendClient

func init() {
	flag.StringVar(&parse.Path, "file", "", "template file path")
	flag.StringVar(&config.Backend, "backend", "", "backend name")
	flag.Var(&config.BackendNodes, "node", "list of backend nodes")
	flag.Parse()

	// validation
	if parse.Path == "" {
		fmt.Fprintln(os.Stderr, "file parameter required.")
		os.Exit(1)
	}
	if config.Backend == "" {
		fmt.Fprintln(os.Stderr, "backend parameter required.")
		os.Exit(1)
	}
	if len(config.BackendNodes) == 0 {
		fmt.Fprintln(os.Stderr, "node parameter required.")
		os.Exit(1)
	}

	// backend client
	c, err := b.New(config)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	client = c
}

func getValue(key string) (string, error) {
	return client.GetValue(key)
}

func main() {
	parse.GetValueFuncMap = map[string]interface{}{"v": getValue}
	err := parse.Execute(os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	client.Close()
}