# confgen
`confgen` is configuration file generation tool.

`confgen` is similar to [confd](https://github.com/kelseyhightower/confd), but you do not need template resources.
All you need is a configuration template file.

## Support Backend
- etcdv3
- consul
- zookeeper

## Build
```bash
$ make all
```

## Usage
```bash
$ confgen -h
Usage of ./confgen:
  -backend string
    	backend name
  -file string
    	template file path
  -node value
    	list of backend nodes
```

## Quick Start (example: etcd)
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
$ confgen -backend etcd -file example/.env.tmpl -node localhost:2379
DBHOST = localhost

$ confgen -backend etcd -file example/application.conf.tmpl -node localhost:2379
db {
  host = "localhost"
}

$ confgen -backend etcd -file example/config.toml.tmpl -node localhost:2379
[server]
DBHost = localhost
```

### Output File
```bash
$ confgen -backend etcd -file example/.env.tmpl -node localhost:2379 > .env
```
