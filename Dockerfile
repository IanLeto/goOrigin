FROM golang:1.22
WORKDIR /app

COPY ./ori $workdir
EXPOSE 8008
CMD ["./ori -c /Users/ian/workdir/cc/goOrigin/config.yaml"]
