# goOrigin

自用golang 服务原型

#### package 说明

    agent 服务端 也就是svc - agent 架构中的agent 和boke历史
    client 客户端 用来请求其它服务的 比如请求 tool 和 agent
    config 配置文件  这东西很恶心
    logs 日志 还没配置
    internal 服务内部的一些东西 
    pkg 服务外部的一些东西, 也就是对外的一些东西, 其中他和internal 的区别在于,internal 包含了mvc 架构,会处理业务逻辑,而pkg 只是对外的一些东西,比如grpc 的接口,http 的接口,或者是一些工具类相当于dao层
    internal 里面会定义业务库表结构,而pkg 里面只会定义接口,不会定义库表结构

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
    ./ori --init --config /home/ian/workdir/cc/goOrigin/conn/config.yaml
    migrate
    ./ori migrate --path /Users/ian/workdir/cc/goOrigin/configV2.yaml -p Root3sadf@#$ -u root --host bj-cdb-g2ujielc.sql.tencentcdb.com:22553 --db go_ori --type mysql

    初始化es 模板
    ./ori --initEs --config /home/ian/workdir/cc/goOrigin/conn/config.yaml --data /home/ian/workdir/cc/goOrigin/conn/elasticsearch.json
    
    性能测试
    cd /home/ian/workdir/goOrigin/pkg/processor
    go test -run=^$ -bench=BenchmarkFullCache -benchmem -cpuprofile full_cache_cpu.prof -memprofile full_cache_mem.prof
    go tool pprof -alloc_space -top full_cache_mem.prof

#### ci/cd

    GOOS=linux GOARCH=amd64 go build -o ori main.go && \
    docker build -t ianleto/goorigin:$(git rev-parse --short HEAD) -f Dockerfile2 .&&\
    docker push ianleto/goorigin:$(git rev-parse --short HEAD) &&\
    docker images | grep ianleto/ianhello | awk '{print $3}' | xargs docker rmi -f
    docker ps -a | grep 10c831814d55 | awk '{print $1}' | xargs docker rm -f
    
####      
