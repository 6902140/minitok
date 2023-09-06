package rpc

import (
	"context"
	"minitok/cmd/api/global"
	"minitok/kitex_gen/douyin/publish"
	"minitok/pkg/errno"
)

func PublishAction(ctx context.Context, req *publish.ActionRequest) (*publish.ActionResponse, error) {
	if global.PublishServiceClient == nil {
		return nil, errno.ServiceErr.WithMessage("用户微服务客户端未初始化或初始化失败")
	}
	return (*global.PublishServiceClient).PublishAction(ctx, req)
}

func PublishList(ctx context.Context, req *publish.ListRequest) (*publish.ListResponse, error) {
	if global.PublishServiceClient == nil {
		return nil, errno.ServiceErr.WithMessage("用户微服务客户端未初始化或初始化失败")
	}
	return (*global.PublishServiceClient).PublishList(ctx, req)
}
