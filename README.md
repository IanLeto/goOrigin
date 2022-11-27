# goOrigin
自用golang 服务原型

#### logs
    proto 三方库问题
    add log init  9/30
    add gorm mysql init 10/16
    增加 k8s curd api 2022/7/26 未测试
#### 常用命令
    docker build --tag ianleto/goorigin:
    build  GOOS=linux GOARCH=amd64 go build -o ianhello 
    cd  /Users/ian/go/src/goOrigin/agent/protos
    protoc -I ../protos --go_out=plugins=grpc:../pbs ../protos/task.proto
    cd /Users/ian/go/src/goOrigin/agent/protos
    protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative agent.proto
    生成调用关系图
    cd /Users/ian/go/src/goOrigin/agent
    godepgraph -s ../agent | dot -Tpng -o godepgraph.png
    open .
    