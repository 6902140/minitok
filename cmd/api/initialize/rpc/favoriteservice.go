package rpc

import (
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"minitok/cmd/api/global"
	favorite "minitok/kitex_gen/douyin/favorite/favoriteservice"
)

func InitFavoriteRPC() (*favorite.Client, error) {
	r, err := etcd.NewEtcdResolver([]string{global.Configs.ETCD.Addr()})
	if err != nil {
		return nil, err
	}
	newClient, err := favorite.NewClient(
		global.Configs.RPCClient.FavoriteServiceName,
		client.WithMuxConnection(1),
		client.WithRPCTimeout(3*time.Second),
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),
	)
	if err != nil {
		return nil, err
	}
	return &newClient, nil
}
