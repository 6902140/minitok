package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"minitok/cmd/user/dal"
	"minitok/cmd/user/dal/mongo"
	"minitok/cmd/user/global"
	"minitok/cmd/user/pack"
	"minitok/kitex_gen/douyin/user"
	"minitok/pkg/errno"
	"minitok/pkg/jwt"
)

type UserInfoService struct {
	ctx context.Context
}

func NewUserInfoService(ctx context.Context) *UserInfoService {
	return &UserInfoService{ctx: ctx}
}

func (s *UserInfoService) UserInfo(req *user.InfoRequest) (*user.User, error) {
	// 根据uid查询用户信息
	userInfo, err := dal.QueryUserInfoById(s.ctx, req.UserId)
	if err != nil {
		// 错误信息已经处理，可直接返回
		return nil, err
	}
	// 解析token, 判断关注状态
	isFollow := false
	if len(req.Token) != 0 {
		// 解析token
		claims, err := jwt.NewJWT(global.Configs.JWT.SigningKey).ParseToken(req.Token)
		if err != nil {
			klog.Errorf("jwt error: %v\n", err)
			return nil, errno.ServiceErr.WithMessage(err.Error())
		}
		// 校验信息
		if claims.Id == 0 || claims.Issuer != global.Configs.JWT.Issuer || claims.Subject != global.Configs.JWT.Subject {
			return nil, errno.AuthorizationFailedErr
		}
		// 判断关注状态
		if claims.Id != req.UserId {
			isFollow, err = mongo.GetFollowInfo(s.ctx, claims.Id, req.UserId)
			//isFollow, err = cache.GetFollowState(s.ctx, claims.Id, req.UserId)
			if err != nil {
				klog.Errorf("mongo query error: %v\n", err)
				return nil, errno.ServiceErr.WithMessage(err.Error())
			}
		}
	}
	return pack.BuildRespUser(userInfo, isFollow), nil
}
