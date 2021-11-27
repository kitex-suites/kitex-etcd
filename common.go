package kitexetcd

import (
	"errors"
	"time"
)

const (
	etcdPrefix    = "kitex-etcd"
	connTimeout   = time.Second * 2
	defaultWeight = 10
)

var (
	errorEtcdUrlEmpty = errors.New("")
)

type NewEtcdConfig struct {
	EtcdUrl  string
	Username string
	Password string
}

func getEtcdPrefix(serviceName string) string {
	return etcdPrefix + "$" + serviceName
}

func getEtcdKey(serviceName, addr string) string {
	return getEtcdPrefix(serviceName) + "$" + addr
}
