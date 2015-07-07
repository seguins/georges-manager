= Georges Manager

**This repo is work in progress. Actually nothing works.**

This project used:
* ETCD
* Vulcand
* Syntax file of docker-compose


== Dev help

=== Start ETCD

```
docker run -p 2380:2380 -p 7001:7001 -p 2379:2379 -p 4001:4001 --rm -it quay.io/coreos/etcd --advertise-client-urls 'http://0.0.0.0:2379,http://0.0.0.0:4001'  --listen-client-urls 'http://0.0.0.0:2379,http://0.0.0.0:4001' --initial-advertise-peer-urls 'http://0.0.0.0:2380,http://0.0.0.0:7001' --listen-peer-urls 'http://0.0.0.0:2380,http://0.0.0.0:7001' --initial-cluster 'default=http://0.0.0.0:2380,default=http://0.0.0.0:7001'
```

=== Start Vulcand

```
docker run --rm -it -p 8182:8182 -p 8181:8181 mailgun/vulcand:v0.8.0-beta.2 /go/bin/vulcand -apiInterface=0.0.0.0 --etcd=http://172.17.0.20:4001
```

=== Setup backend

```
etcdctl set /vulcand/backends/b1/servers/srv1 '{"URL": "http://172.17.0.22"}'
etcdctl set /vulcand/backends/b2/servers/srv1 '{"URL": "http://172.17.0.30"}'
```

=== Setup frontend

```
etcdctl set /vulcand/frontends/f1/frontend '{"Type": "http", "BackendId": "b1", "Route": "Host(`localhost`) && Path(`/`)"}'
etcdctl set /vulcand/frontends/f2/frontend '{"Type": "http", "BackendId": "b2", "Route": "Host(`seguins`) && Path(`/`)"}'
```
