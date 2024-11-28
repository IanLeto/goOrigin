GOOS=linux GOARCH=amd64 go build -o ori main.go && \
    docker build -t ianleto/goorigin:last -f Dockerfile2 .&&\
    docker push ianleto/goorigin:last