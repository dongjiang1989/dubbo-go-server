package main

import (
	"context"
	"os"

	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"github.com/dubbogo/gost/log/logger"

	node "github.com/dongjiang1989/dubbo-go-server/pkg/proto"
)

type NodeProvider struct {
	node.UnimplementedNodeServiceServer
}

func (s *NodeProvider) GetNodeInfo(ctx context.Context, _ *node.NodeRequest) (*node.NodeResponse, error) {
	logger.Infof("Dubbo3 NodeProvider get node info")
	podIP := os.Getenv("POD_IP")
	podName := os.Getenv("POD_NAME")
	namespace := os.Getenv("POD_NAMESPACE")
	return &node.NodeResponse{
		Name:      podName,
		Ip:        podIP,
		Namespace: namespace}, nil
}

// export DUBBO_GO_CONFIG_PATH= ./conf/dubbogo.yml
func main() {
	config.SetProviderService(&NodeProvider{})
	if err := config.Load(); err != nil {
		panic(err)
	}
	select {}
}
