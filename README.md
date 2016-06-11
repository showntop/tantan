# tantan
==========

users and relationships restful api

##requirement

```
go 1.6
postgresql 9.5.1
git 2.7.2
```

## install

```
1.copy the project to your go work path
    mkdir -p $GOPATH/src/github.com/showntop
    cp {tantan} -rf $GOPATH/src/github.com/showntop/
    cd $GOPATH/src/github.com/showntop/tantan
    go get
2. use go get install
    go get github.com/showntop/tantan
    cd $GOPATH/src/github.com/showntop/tantan
    go get
```

## config
```
the config info in the tantan/config directory
in the config go modify the http port and database option

```

## run 

```
1. createdb tantan or the name you specified in the config file
1. go run main.go or go build && ./tantan
2. then test it
    curl -XGET "http://localhost:9000/users"
    .........
```