package main

import (
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"minitok/cmd/comment/global"
	"minitok/cmd/comment/initialize"
	comment "minitok/kitex_gen/douyin/comment/commentservice"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func LoadConfigsAndInit() {
	var err error
	if global.Viper, err = initialize.Viper("config.yaml"); err != nil {
		panic(err)
	}
	if global.GormDB, err = initialize.GormMySQL(); err != nil {
		panic(err)
	}
}

func main() {
	LoadConfigsAndInit()

	r, err := etcd.NewEtcdRegistry([]string{global.Configs.ETCD.Addr()})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", global.Configs.RPCServer.Addr())
	if err != nil {
		panic(err)
	}

	svr := comment.NewServer(new(CommentServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: global.Configs.RPCServer.ServiceName,
		}),
		server.WithServiceAddr(addr),
		server.WithLimit(&limit.Option{
			MaxConnections: 1000,
			MaxQPS:         100,
		}),
		server.WithMuxTransport(),
		server.WithSuite(trace.NewDefaultServerSuite()),
		server.WithRegistry(r),
	)

	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
