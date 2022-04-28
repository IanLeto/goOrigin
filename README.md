# goOrigin
自用golang 服务原型

#### logs
    add log init  9/30
    add gorm mysql init 10/16
#### 常用命令
    cd  /Users/ian/go/src/goOrigin/agent/protos
    protoc -I ../protos --go_out=plugins=grpc:../pbs ../protos/task.proto