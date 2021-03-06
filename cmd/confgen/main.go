package main

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"strings"

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
	flag.StringVar(&config.Key, "key", "", "key name")
	flag.Parse()

	// validation
	if parse.Path == "" && config.Key == "" {
		fmt.Fprintln(os.Stderr, "file or key parameter required.")
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

func getValue(key string) string {
	v, err := client.GetValue(key)
	if err != nil {
		panic(err)
	}
	return v
}

func getArrayValue(key string, sep string) []string {
	v := getValue(key)
	return strings.Split(v, sep)
}

func join(values []string, sep string) string {
	ret := strings.Join(values, sep)
	return ret
}

func last(i int, a interface{}) bool {
	return i == reflect.ValueOf(a).Len()-1
}

func main() {
	if config.Key != "" {
		fmt.Print(getValue(config.Key))
		return
	}

	parse.GetValueFuncMap = map[string]interface{}{
		"v":       getValue,
		"explode": getArrayValue,
		"join":    join,
		"last":    last,
	}
	err := parse.Execute(os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	client.Close()
}
