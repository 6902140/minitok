package rpc

import (
	"context"

	"minitok/cmd/api/global"
	"minitok/kitex_gen/douyin/feed"
	"minitok/pkg/errno"
)

func Feed(ctx context.Context, req *feed.FeedRequest) (*feed.FeedResponse, error) {
	if global.FeedServiceClient == nil {
		return nil, errno.ServiceErr.WithMessage("用户微服务客户端未初始化或初始化失败")
	}
	return (*global.FeedServiceClient).Feed(ctx, req)
}
