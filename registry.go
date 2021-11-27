package kitexetcd

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/cloudwego/kitex/pkg/registry"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type etcdRegistry struct {
	etcdClient *clientv3.Client
}

// for mashal and unmashal
type registryInfo struct {
	ServiceName string            `json:"service_name"`
	Network     string            `json:"network"`
	Addr        string            `json:"addr"`
	Weight      int               `json:"weight"`
	Tags        map[string]string `json:"tags"`
}

func NewEtcdRegistry(config *NewEtcdConfig) (registry.Registry, error) {
	if config.EtcdUrl == "" {
		return nil, errorEtcdUrlEmpty
	}

	c, err := clientv3.New(clientv3.Config{
		Endpoints: []string{config.EtcdUrl},
		Username:  config.Username,
		Password:  config.Password,
	})
	if err != nil {
		return nil, err
	}

	return &etcdRegistry{
		etcdClient: c,
	}, nil
}

func (registry *etcdRegistry) Register(info *registry.Info) error {
	if info.ServiceName == "" || info.Addr.String() == "" {
		return errors.New("ServiceName or Addr empty error")
	}

	return registry.register(info)
}

func (registry *etcdRegistry) Deregister(info *registry.Info) error {
	return registry.deregister(info)
}

func (registry *etcdRegistry) register(info *registry.Info) error {
	regInfo := &registryInfo{
		ServiceName: info.ServiceName,
		Addr:        info.Addr.String(),
		Network:     info.Addr.Network(),
		Weight:      info.Weight,
		Tags:        info.Tags,
	}

	infoStr, err := json.Marshal(regInfo)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), connTimeout)
	defer cancel()

	_, err = registry.etcdClient.Put(ctx, getEtcdKey(regInfo.ServiceName, regInfo.Addr), string(infoStr))
	return err
}

func (registry *etcdRegistry) deregister(info *registry.Info) error {
	ctx, cancel := context.WithTimeout(context.Background(), connTimeout)
	defer cancel()

	_, err := registry.etcdClient.Delete(ctx, getEtcdKey(info.ServiceName, info.Addr.String()))
	return err
}
