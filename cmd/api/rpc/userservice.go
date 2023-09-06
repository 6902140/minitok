package rpc

import (
	"context"
	"minitok/cmd/api/global"
	"minitok/kitex_gen/douyin/user"
	"minitok/pkg/errno"
)

func UserRegister(ctx context.Context, req *user.RegisterRequest) (*user.RegisterResponse, error) {
	if global.UserServiceClient == nil {
		return nil, errno.ServiceErr.WithMessage("用户微服务客户端未初始化或初始化失败")
	}
	return (*global.UserServiceClient).UserRegister(ctx, req)
}

func UserLogin(ctx context.Context, req *user.LoginRequest) (*user.LoginResponse, error) {
	if global.UserServiceClient == nil {
		return nil, errno.ServiceErr.WithMessage("用户微服务客户端未初始化或初始化失败")
	}
	return (*global.UserServiceClient).UserLogin(ctx, req)
}

func UserInfo(ctx context.Context, req *user.InfoRequest) (*user.InfoResponse, error) {
	if global.UserServiceClient == nil {
		return nil, errno.ServiceErr.WithMessage("用户微服务客户端未初始化或初始化失败")
	}
	return (*global.UserServiceClient).UserInfo(ctx, req)
}
