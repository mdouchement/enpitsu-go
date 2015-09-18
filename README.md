# Enpitsu

## Compiling

- Install Goop dependencies manager
```sh
$ go get github.com/nitrous-io/goop
```

- Install packing resources manager (QML assets)
```sh
$ go get gopkg.in/qml.v1/cmd/genqrc
```

- Install dependencies
```sh
$ goop install
```
> go-qml compilation details: https://github.com/go-qml/qml

- Packaging resources
```sh
$ goop go generate
```
> Packaging details: http://blog.labix.org/2014/09/26/packing-resources-into-go-qml-binaries

- Enpitsu
```sh
# Execution
$ goop go run *.go

# Execution in development
$ QRC_REPACK=1 goop go run *.go

# Compilation
$ goop go build *.go
```
