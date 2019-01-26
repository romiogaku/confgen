# confgen
`confgen` is simple [confd](https://github.com/kelseyhightower/confd).

`confgen` does not need template resources.
All you need is a configuration template file.

## Support Backend
- etcdv3

## Build
```bash
$ dep ensure
$ go build
```

## Usage
```bash
$ confgen -h
Usage of ./confgen:
  -file string
    	template file path
  -node value
    	list of backend nodes
```

## Quick Start
### Install etcd
https://github.com/etcd-io/etcd/releases

### Run etcd
```
$ etcd
```

### Write key
```
$ ETCDCTL_API=3 etcdctl put /myapp/database/host "localhost"
OK
```
### Run confgen

```bash
$ confgen -file example/.env.tmpl -node localhost:2379
DBHOST = localhost

$ confgen -file example/application.conf.tmpl -node localhost:2379
db {
  host = "localhost"
}

$ confgen -file example/config.toml.tmpl -node localhost:2379
[server]
DBHost = localhost
```

### Output File
```bash
$ confgen -file example/.env.tmpl -node localhost:2379 > .env
```