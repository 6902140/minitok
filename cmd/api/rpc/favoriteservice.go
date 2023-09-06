package rpc

import (
	"context"

	"minitok/cmd/api/global"
	"minitok/kitex_gen/douyin/favorite"
	"minitok/pkg/errno"
)

func FavoriteAction(ctx context.Context, req *favorite.ActionRequest) (*favorite.ActionResponse, error) {
	if global.FavoriteServiceClient == nil {
		return nil, errno.ServiceErr.WithMessage("用户微服务客户端未初始化或初始化失败")
	}
	return (*global.FavoriteServiceClient).FavoriteAction(ctx, req)
}

func FavoriteList(ctx context.Context, req *favorite.ListRequest) (*favorite.ListResponse, error) {
	if global.FavoriteServiceClient == nil {
		return nil, errno.ServiceErr.WithMessage("用户微服务客户端未初始化或初始化失败")
	}
	return (*global.FavoriteServiceClient).FavoriteList(ctx, req)
}
