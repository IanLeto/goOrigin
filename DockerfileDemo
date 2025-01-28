FROM golang:1.17-alpine
# 指定工作目录
# 工作目录 也就是起始目录 毕竟我们不能吧 / 作为根目录吧
WORKDIR /Users/ian/go/src/goOrigin
COPY . /Users/ian/go/src/goOrigin
RUN go env -w GOPROXY=https://goproxy.cn,direct
#COPY --from=busybox:1.28 /bin/busybox /bin/busybox
#RUN apk add build-base
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# run go build build 因为我们已经指定工作目录了，相当于在/root 下执行go huild
RUN go build  -o ori main.go
EXPOSE 8008
#CMD ["dlv", "--listen=:2345", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "./ori"]
CMD ["./ori"]

#FROM scratch
#COPY --from=busybox:1.28 /bin/busybox /bin/busybox
