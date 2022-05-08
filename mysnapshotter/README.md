# mysnapshotter

ref. https://github.com/containerd/containerd/blob/68d9d462c934ac727ebafab1c6ebeae10716fbcc/docs/PLUGINS.md

```toml
# /etc/containerd/config.toml
[proxy_plugins]
  [proxy_plugins.customsnapshot]
    type = "snapshot"
    address = "/var/run/mysnapshotter.sock"
```

```shell
$ go build main.go
```

```shell
$ sudo ./main /var/run/mysnapshotter.sock /tmp/snapshots

$ sudo CONTAINERD_SNAPSHOTTER=customsnapshot ctr images pull docker.io/library/hello-world:latest
docker.io/library/hello-world:latest:                                             resolved       |++++++++++++++++++++++++++++++++++++++|
index-sha256:10d7d58d5ebd2a652f4d93fdd86da8f265f5318c6a73cc5b6a9798ff6d2b2e67:    done           |++++++++++++++++++++++++++++++++++++++|
manifest-sha256:f54a58bc1aac5ea1a25d796ae155dc228b3f0e11d046ae276b39c4bf2f13d8c4: done           |++++++++++++++++++++++++++++++++++++++|
layer-sha256:2db29710123e3e53a794f2694094b9b4338aa9ee5c40b930cb8063a1be392c54:    done           |++++++++++++++++++++++++++++++++++++++|
config-sha256:feb5d9fea6a5e9606aa995e879d862b825965ba48de054caab5ef356dc6b3412:   done           |++++++++++++++++++++++++++++++++++++++|
elapsed: 5.5 s                                                                    total:  4.4 Ki (828.0 B/s)
unpacking linux/amd64 sha256:10d7d58d5ebd2a652f4d93fdd86da8f265f5318c6a73cc5b6a9798ff6d2b2e67...
done: 34.574841ms

$ sudo ls /tmp/snapshots/snapshots/2/hello
/tmp/snapshots/snapshots/2/hello
```
