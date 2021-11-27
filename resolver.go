package kitexetcd

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type etcdResolver struct {
	etcdClient *clientv3.Client
}

func NewEtcdResolver(config *NewResolverConfig) (discovery.Resolver, error) {
	c, err := clientv3.New(clientv3.Config{
		Endpoints: []string{config.EtcdUrl},
		Username:  config.Username,
		Password:  config.Password,
	})

	if err != nil {
		return nil, err
	}

	return &etcdResolver{
		etcdClient: c,
	}, nil
}

func (resolver *etcdResolver) Target(ctx context.Context, target rpcinfo.EndpointInfo) (description string) {
	return target.ServiceName()
}

func (resolver *etcdResolver) Resolve(ctx context.Context, desc string) (discovery.Result, error) {
	// keys are in the format of "kitex-etcd$serviceName$address", and there can be many instances for
	// one serviceName, so get all the kvs with prefix "kitex-etcd$serviceName"
	resp, err := resolver.etcdClient.Get(ctx, getEtcdPrefix(desc), clientv3.WithPrefix())
	if err != nil {
		return discovery.Result{}, err
	}

	var info registryInfo
	var instances []discovery.Instance
	for _, kv := range resp.Kvs {
		err := json.Unmarshal(kv.Value, &info)
		if err != nil {
			continue
		}

		instances = append(instances, discovery.NewInstance(info.Network, info.Addr, info.Weight, info.Tags))
	}

	if len(instances) == 0 {
		return discovery.Result{}, errors.New("no instance resolved error")
	}

	return discovery.Result{
		Cacheable: true,
		CacheKey:  desc,
		Instances: instances,
	}, nil
}

func (resolver *etcdResolver) Diff(cacheKey string, prev, next discovery.Result) (discovery.Change, bool) {
	return discovery.DefaultDiff(cacheKey, prev, next)
}

func (resolver *etcdResolver) Name() string {
	return "kitex-etcd"
}
