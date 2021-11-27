package kitexetcd_test

import (
	"testing"

	kitexetcd "github.com/kitex-suites/kitex-etcd"
)

func TestEtcdResolverName(t *testing.T) {
	resolver, err := kitexetcd.NewEtcdResolver(&kitexetcd.NewResolverConfig{
		EtcdUrl: "http://127.0.0.1:2397",
	})
	if err != nil {
		panic(err)
	}

	t.Logf("resolver name: %s", resolver.Name())
}
