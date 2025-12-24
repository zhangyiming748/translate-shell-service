FROM golang:1.25.4-alpine3.21 AS builder
WORKDIR /app
COPY . .
# RUN go env -w GO111MODULE=on
# RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN sed -i 's#https\?://dl-cdn.alpinelinux.org/alpine#http://mirrors4.tuna.tsinghua.edu.cn/alpine#g' /etc/apk/repositories
RUN go build -o gin .

FROM alpine:3.21
COPY --from=builder /app/gin /usr/local/bin/gin
EXPOSE 6380
RUN apk update
RUN apk add translate-shell curl
RUN sed -i 's#https\?://dl-cdn.alpinelinux.org/alpine#http://mirrors4.tuna.tsinghua.edu.cn/alpine#g' /etc/apk/repositories
ENTRYPOINT ["/usr/local/bin/gin"]
# docker build -t gin .
# docker run -d -p 8192:80 --name gin  gin