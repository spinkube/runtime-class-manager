# Local testing during development

Run

```
$ go run . install -H test -a ./test/assets
```

This will copy the assets from `test/assets` into `test/opt/kwasm/bin` and
create containerd config files in `test/etc/containerd/conf.d`. The config will
have the absolute path of the binary from the host root (`-H test` - `test` is
the host root).