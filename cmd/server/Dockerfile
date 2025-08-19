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
RUN go build -o server ./cmd/server/main.go

FROM alpine:3.22
WORKDIR /dubbo-go-server
COPY --from=builder /dubbo-go-server/server ./server
COPY --from=builder /dubbo-go-server/cmd/server/conf/dubbogo.yml .

ENV DUBBO_GO_CONFIG_PATH="./dubbogo.yml"
EXPOSE 50052
ENTRYPOINT ["./server"]