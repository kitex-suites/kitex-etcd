package kitexetcd

import "github.com/cloudwego/kitex/pkg/registry"

type etcdRegistry struct{}

func NewEtcdRegistry() {

}

func (registry *etcdRegistry) Register(info *registry.Info) error {
	return nil
}

func (registry *etcdRegistry) DeRegister(info *registry.Info) error {
	return nil
}
