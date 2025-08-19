package main

import (
	"context"

	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"github.com/dubbogo/gost/log/logger"

	node "github.com/dongjiang1989/dubbo-go-server/pkg/proto"
)

var grpcNodeImpl = new(node.NodeServiceClientImpl)

func init() {
	config.SetConsumerService(grpcNodeImpl)
}

// export DUBBO_GO_CONFIG_PATH= ./conf/dubbogo.yml
func main() {
	if err := config.Load(); err != nil {
		panic(err)
	}

	logger.Info("start to dubbo client")
	req := &node.NodeRequest{}
	reply, err := grpcNodeImpl.GetNodeInfo(context.Background(), req)
	if err != nil {
		logger.Error(err)
	}
	logger.Infof("client response result: %v\n", reply)
}
