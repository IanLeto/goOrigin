GOOS=linux GOARCH=amd64 go build -o ori main.go && \
    docker build -t ianleto/goorigin:$(git rev-parse --short HEAD) -f Dockerfile2 .&&\
    docker push ianleto/goorigin:$(git rev-parse --short HEAD)