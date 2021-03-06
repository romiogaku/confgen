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
  -backend string
    	backend name
  -file string
    	template file path
  -key string
    	key name
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

### Simple get value
```bash
$ confgen -backend etcd -node dev-confstore01:2379 -key /myapp/database/host
```

## Template Functions
### v
Returns the value as a string where key matches its argument.

Data:

| K   | V     |
|-----|-------|
| key | value |

Template:

```
value: {{v "/key"}}
```

Output:

```
value: value
```

### explode
Split a string by a delimiter. Returns an array of string.

Data:

| K     | V                 |
|-------|-------------------|
| hosts | host1,host2,host3 |

Template:

```
{{$hosts := explode "/hosts" ","}}
{{$hosts}}
```

Output:

```
[host1 host2 host3]
```



### join
Alias for the [strings.Join](https://golang.org/pkg/strings/#Join) function.

Data:

| K     | V                 |
|-------|-------------------|
| hosts | host1,host2,host3 |

Template:

```
{{$hosts := explode "/hosts" ","}}
hosts = ["{{join $hosts "\",\""}}"]
```

Output:

```
hosts = ["host1","host2","host3"]
```

### last
Returns true if first argument is last index of second argument.
Second argument is array.

Data:

| K     | V                 |
|-------|-------------------|
| hosts | host1,host2,host3 |

Template:

```
{{$hosts := explode "/hosts" ","}}
hosts = [{{range $i, $e := $hosts}}"{{$e}}"{{if not (last $i $hosts)}},{{end}}{{end}}]
```

Output:

```
hosts = ["host1","host2","host3"]
```
