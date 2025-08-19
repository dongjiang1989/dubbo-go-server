package zookeeper

import (
	"dubbo.apache.org/dubbo-go/v3"
	"dubbo.apache.org/dubbo-go/v3/registry"
)

func NewRegistryZKOption(url string) dubbo.InstanceOption {
	return dubbo.WithRegistry(
		registry.WithZookeeper(),
		registry.WithAddress(url),
	)
}
