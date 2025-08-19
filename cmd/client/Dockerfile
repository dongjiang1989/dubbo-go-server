FROM golang:1.24-alpine3.22 AS builder
ENV TIMEZONE Asia/Shanghai
WORKDIR /dubbo-go-server
#添加文件
COPY go.mod ./
COPY go.sum ./
COPY conf ./conf
COPY cmd ./cmd
COPY pkg ./pkg

##编译
RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct && go mod download
RUN go build -o client ./cmd/client/main.go

FROM alpine:3.22
WORKDIR /dubbo-go-server
COPY --from=builder /dubbo-go-server/client ./client
COPY --from=builder /dubbo-go-server/cmd/client/conf/dubbogo.yml .

ENV DUBBO_GO_CONFIG_PATH="./dubbogo.yml"
EXPOSE 50052
ENTRYPOINT ["./client"]