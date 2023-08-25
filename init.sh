#### 首次使用
git pull origin master
GOOS=linux GOARCH=amd64 go build -o ori main.go
docker build -t ianleto/goorigin -f Dockerfile2 .
docker push ianleto/goorigin:latest
alias  k=kubectl
k apply -f config.yaml
k delete -f deploy.yaml
k apply -f deploy.yaml
curl http://localhost:8080/origin/healthz