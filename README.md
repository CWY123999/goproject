## Usage:

```shell
$ go get -u github.com/CWY123999/goproject
$ cd $GOPATH/src/github.com/CWY123999/goproject/cmd/goproject
$ go install

$ goproject new testproject
```

## Golang 编码规范

https://github.com/CWY123999/goproject/blob/master/standard.md


## Golang 项目 CI 支持说明

https://github.com/CWY123999/goproject/-/tree/master/scripts

## 关于项目中的Assets

通过 `go-bindata` 工具使用以下命令生成:

```shell
go-bindata -o assets/assets.go -pkg=assets -ignore="\\.DS_Store|README.md" scripts/...
```
