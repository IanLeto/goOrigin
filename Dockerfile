FROM golang:1.16-alpine
# 指定工作目录
# 工作目录 也就是起始目录 毕竟我们不能吧 / 作为根目录吧
WORKDIR /root
COPY . /root
# run go build build 因为我们已经指定工作目录了，相当于在/root 下执行go huild
RUN go build  main.go
EXPOSE 1919
ENTRYPOINT ["./main"]

#FROM scratch
#COPY --from=busybox:1.28 /bin/busybox /bin/busybox
