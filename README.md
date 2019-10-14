# dkv
DKV is a distributed key value store server written in [Go](https://golang.org). It exposes all its functionality over
[gRPC](http://www.grpc.io) & [Protocol Buffers](https://developers.google.com/protocol-buffers/).

## Features
- Data Sharding
- Tunable consistency
- Data replication over WANs

## Supported APIs
- createVBucket(replicationFactor)
- put(K,V,vBucket)
- get(K,consistency)

## Design
<img src="https://github.com/flipkart-incubator/dkv/raw/master/docs/design.png">

## Dependencies
- Go version 1.11+
- [RocksDB](https://github.com/facebook/rocksdb) v5.16+ as the storage engine
- [GoRocksDB](https://github.com/tecbot/gorocksdb) provides the CGo bindings with RocksDB
- [Raft consensus algorithm](https://raft.github.io/) by [etcd/raft](https://github.com/etcd-io/etcd/tree/master/raft)

## Building DKV on Mac OSX

### Building RocksDB
- Ensure [HomeBrew](https://brew.sh/) is installed
- brew install gcc49
- brew install rocksdb
- Install [ZStd](https://github.com/facebook/zstd)
- Execute this command:

```bash
CGO_CFLAGS="-I/usr/local/Cellar/rocksdb/6.1.2/include" CGO_LDFLAGS="-L/usr/local/Cellar/rocksdb/6.1.2 -lrocksdb -lstdc++ -lm -lz -lbz2 -lsnappy -llz4 -lzstd" go get github.com/tecbot/gorocksdb
```

### Building DKV

```bash
$ mkdir -p ${GOPATH}/src/github.com/flipkart-incubator
$ cd ${GOPATH}/src/github.com/flipkart-incubator
$ git clone https://github.com/flipkart-incubator/dkv
$ cd dkv
$ make build
```

If you want to build for other platform, set `GOOS`, `GOARCH` environment variables. For example, build for macOS like following:

```bash
$ make GOOS=linux build
```

## Testing DKV

If you want to test your changes, run command like following:

```bash
$ make test
```

## Packaging DKV

###  Linux

```bash
$ make GOOS=linux dist
```

### macOS

```bash
$ make GOOS=darwin dist
```

