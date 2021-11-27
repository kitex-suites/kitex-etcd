# kitex etcd

## Introduction

``kitexetcd`` is an implemention of service registry and service discovery for [kitex]("https://github.com/cloudwego/kitex") based on etcd.

## Installation

```shell
go get -u github.com/kitex-suites/kitex-etcd
```

## Documenttation

``kitexetcd`` should be used both on server-side and client-side.

### Server-Side Service Registry

```go
func main() {
    rgst, _ := kitexetcd.NewEtcdRegistry(&kitexetcd.NewRegistryConfig{
        // http url of etcd server
        EtcdUrl: "http://localhost:2379",
        Weight: 10,
    })

    svr := item.NewServer(new(ItemServiceImpl), server.WithRegistry(rgst))

    err := svr.Run()

    if err != nil {
        log.Println(err.Error())
    }
}
```

### Client-Side service Discovery

```go
func main() {
    rsv, _ := kitexetcd.NewEtcdResolver(&kitexetcd.NewResolverConfig{
        // http url of etcd server
        EtcdUrl: "http://localhost:2379",
    })

    c, err := itemservice.NewClient("kitex.demo.item", client.WithResolver(rsv))
    if err != nil {
        panic(err)
    }

    req := &item.GetItemRequest{ Id: 1, }
    resp, _ := c.GetItem(context.Background(), req)

    fmt.Println(resp)
}
```
