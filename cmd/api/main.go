// Code generated by hertz generator.

package main

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/network/standard"
	"minitok/cmd/api/global"
	"minitok/cmd/api/initialize"
	"minitok/cmd/api/initialize/rpc"
	"minitok/pkg/path"
	"os"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func LoadConfigsAndInit() {
	var err error
	if exist, err := path.FileExist("config.yaml"); err != nil || !exist {
		fmt.Println("未找到配置文件，无法启动服务")
		os.Exit(0)
	}
	if global.Viper, err = initialize.Viper("config.yaml"); err != nil {
		panic(err)
	}
	if global.UserServiceClient, err = rpc.InitUserRPC(); err != nil {
		panic(err)
	}
	if global.PublishServiceClient, err = rpc.InitPublishRPC(); err != nil {
		panic(err)
	}
	if global.FeedServiceClient, err = rpc.InitFeedRPC(); err != nil {
		panic(err)
	}
	if global.FavoriteServiceClient, err = rpc.InitFavoriteRPC(); err != nil {
		panic(err)
	}
	if global.CommentServiceClient, err = rpc.InitCommentRPC(); err != nil {
		panic(err)
	}
	if global.RelationServiceClient, err = rpc.InitRelationRPC(); err != nil {
		panic(err)
	}
	if global.MessageServiceClient, err = rpc.InitMessageRPC(); err != nil {
		panic(err)
	}
}

func main() {
	LoadConfigsAndInit()

	h := server.Default(
		server.WithStreamBody(true),
		server.WithAltTransport(standard.NewTransporter),
		server.WithHostPorts(global.Configs.Hertz.Addr()),
	)

	register(h)
	h.Spin()
}
