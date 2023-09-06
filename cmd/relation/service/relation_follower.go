package service

import (
	"context"
	"minitok/cmd/relation/dal"
	"minitok/cmd/relation/dal/mongodb"

	"minitok/cmd/relation/global"
	"minitok/kitex_gen/douyin/relation"
	"minitok/kitex_gen/douyin/user"
	"minitok/pkg/errno"
	"minitok/pkg/jwt"
)

type RelationFollowerListService struct {
	ctx context.Context
}

func NewRelationFollowerListService(ctx context.Context) *RelationFollowerListService {
	return &RelationFollowerListService{ctx: ctx}
}

func (s *RelationFollowerListService) FollowerList(req *relation.FollowerListRequest) ([]*user.User, error) {
	claims, err := jwt.NewJWT(global.Configs.JWT.SigningKey).ParseToken(req.Token)
	if err != nil {
		return nil, err
	}
	if claims.Id == 0 || claims.Issuer != global.Configs.JWT.Issuer || claims.Subject != global.Configs.JWT.Subject {
		return nil, errno.AuthorizationFailedErr
	}
	relationInfos, err := mongodb.GetFollowerList(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	userList := make([]*user.User, len(relationInfos))
	for i, uid := range relationInfos {
		userInfo, err := dal.QueryUserInfoById(s.ctx, uid)
		if err != nil {
			return nil, err
		}
		// 关注状态
		isFollow := false
		if claims.Id == req.UserId {
			isFollow = true
		} else {
			isFollow, err = mongodb.GetFollowInfo(s.ctx, claims.Id, req.UserId)
			if err != nil {
				return nil, err
			}
		}
		userList[i] = &user.User{
			Id:       uid,
			Name:     userInfo.Nickname,
			IsFollow: isFollow,
		}
	}
	return userList, nil
}
