package rpc

import (
	"context"

	"minitok/cmd/api/global"
	"minitok/kitex_gen/douyin/message"
	"minitok/pkg/errno"
)

func MessageAction(ctx context.Context, req *message.ActionRequest) (*message.ActionResponse, error) {
	if global.MessageServiceClient == nil {
		return nil, errno.ServiceErr.WithMessage("用户微服务客户端未初始化或初始化失败")
	}
	return (*global.MessageServiceClient).MessageAction(ctx, req)
}

func MessageChat(ctx context.Context, req *message.ChatRequest) (*message.ChatResponse, error) {
	if global.MessageServiceClient == nil {
		return nil, errno.ServiceErr.WithMessage("用户微服务客户端未初始化或初始化失败")
	}
	return (*global.MessageServiceClient).MessageChat(ctx, req)
}
