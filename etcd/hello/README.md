# etcd hello world

## Start etcd by docker

```zsh
docker run --name etcd \
    -p 2379:2379 \
    --volume=etcd-data:/etcd-data \
    --name etcd gcr.io/etcd-development/etcd:v3.4.13 \
    /usr/local/bin/etcd \
      --name=etcd-1 \
      --data-dir=/etcd-data \
      --advertise-client-urls http://0.0.0.0:2379 \
      --listen-client-urls http://0.0.0.0:2379
```
