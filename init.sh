#minikube start
#git pull origin master
#GOOS=linux GOARCH=amd64 go build -o ori main.go
#docker build -t ianleto/goorigin -f Dockerfile2 .  && docker push ianleto/goorigin:latest

alias  k=kubectl
k create configmap config --from-file=/Users/ian/workdir/cc/goOrigin/config.yaml
k delete -f deployment.yaml
k apply -f deployment.yaml
curl http://localhost:8080/origin/healthz