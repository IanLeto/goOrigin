# goOrigin
自用golang 服务原型

#### logs
    add log init  9/30
    add gorm mysql init 10/16
#### 常用命令
    cd  /Users/ian/go/src/goOrigin/agent/protos
    protoc -I ../protos --go_out=plugins=grpc:../pbs ../protos/task.proto
    cd /Users/ian/go/src/goOrigin/agent/protos
    protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative agent.proto
    生成调用关系图
    cd /Users/ian/go/src/goOrigin/agent
    godepgraph -s ../agent | dot -Tpng -o godepgraph.png
    open .
    