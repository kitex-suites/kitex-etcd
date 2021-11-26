package kitexetcd

import (
	"context"

	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

type etcdResolver struct{}

func NewEtcdResolver() {
	//
}

func (resolver *etcdResolver) Target(ctx context.Context, target rpcinfo.EndpointInfo) (description string) {
	return ""
}

func (resolver *etcdResolver) Resolve(ctx context.Context, desc string) (discovery.Result, error) {
	return discovery.Result{}, nil
}

func (resolver *etcdResolver) Diff(cacheKey string, prev, next discovery.Result) (discovery.Change, bool) {
	return discovery.Change{}, false
}

func (resolver *etcdResolver) Name() string {
	return ""
}
