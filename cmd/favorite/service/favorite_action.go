package service

import (
	"context"
	"minitok/cmd/favorite/dal/cache"
	"minitok/cmd/favorite/dal/mongodb"

	"minitok/cmd/favorite/global"

	"minitok/kitex_gen/douyin/favorite"
	"minitok/pkg/constant"
	"minitok/pkg/errno"
	"minitok/pkg/jwt"
)

type FavoriteActionService struct {
	ctx context.Context
}

func NewFavoriteActionService(ctx context.Context) *FavoriteActionService {
	return &FavoriteActionService{ctx: ctx}
}

func (s *FavoriteActionService) FavoriteAction(req *favorite.ActionRequest) error {
	// 1.解析token
	if len(req.Token) == 0 {
		return errno.ParamErr
	}
	claims, err := jwt.NewJWT(global.Configs.JWT.SigningKey).ParseToken(req.Token)
	if err != nil {
		return err
	}
	if claims.Id == 0 || claims.Issuer != global.Configs.JWT.Issuer || claims.Subject != global.Configs.JWT.Subject {
		return errno.AuthorizationFailedErr
	}
	// 2.根据ActionType处理请求
	if req.ActionType == constant.FavoriteActionLike {
		if err := mongodb.AddFavoriteInfo(s.ctx, claims.Id, req.VideoId); err != nil {
			return err
		}
		if err := cache.IncrFavoriteCount(s.ctx, req.VideoId); err != nil {
			return err
		}
	} else if req.ActionType == constant.FavoriteActionCancel {
		if err := mongodb.DeleteFavoriteInfo(s.ctx, claims.Id, req.VideoId); err != nil {
			return err
		}
		if err := cache.DecrFavoriteCount(s.ctx, req.VideoId); err != nil {
			return err
		}
	} else {
		return errno.ParamErr
	}
	return nil
}
