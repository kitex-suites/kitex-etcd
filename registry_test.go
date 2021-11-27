package kitexetcd_test

import (
	"testing"

	kitexetcd "github.com/kitex-suites/kitex-etcd"
)

func TestEtcdRegister(t *testing.T) {
	registry, err := kitexetcd.NewEtcdRegistry(&kitexetcd.NewRegistryConfig{
		EtcdUrl: "http://127.0.0.1:2397",
	})
	if err != nil {
		t.Errorf("error: %v", err)
		t.Fail()
	}

	t.Logf("registry: %v", registry)
}
