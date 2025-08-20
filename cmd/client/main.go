package main

import (
	"context"
	"fmt"
	"time"

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
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	done := make(chan bool)

	go func() {
		var input string
		fmt.Scanln(&input)
		done <- true
	}()

	for {
		select {
		case <-done:
			logger.Infof("stop dubbo client")
			return
		case <-ticker.C:
			req := &node.NodeRequest{}
			reply, err := grpcNodeImpl.GetNodeInfo(context.Background(), req)
			if err != nil {
				logger.Error(err)
			}
			logger.Infof("client response result: %v\n", reply)
		}
	}
}
