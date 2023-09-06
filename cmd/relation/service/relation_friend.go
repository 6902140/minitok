package service

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"minitok/cmd/relation/dal"
	"minitok/cmd/relation/dal/mongodb"
	"minitok/cmd/relation/global"
	"minitok/kitex_gen/douyin/relation"
	"minitok/pkg/constant"
	"minitok/pkg/errno"
	"minitok/pkg/jwt"
)

type RelationFriendListService struct {
	ctx context.Context
}

func NewRelationFriendListService(ctx context.Context) *RelationFriendListService {
	return &RelationFriendListService{ctx: ctx}
}

func (s *RelationFriendListService) FriendList(req *relation.FriendListRequest) ([]*relation.FriendUser, error) {
	claims, err := jwt.NewJWT(global.Configs.JWT.SigningKey).ParseToken(req.Token)
	if err != nil {
		return nil, err
	}
	if claims.Id == 0 || claims.Issuer != global.Configs.JWT.Issuer || claims.Subject != global.Configs.JWT.Subject {
		return nil, errno.AuthorizationFailedErr
	}
	relationInfos, err := mongodb.GetFriendList(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	userList := make([]*relation.FriendUser, len(relationInfos))
	for i, uid := range relationInfos {
		userInfo, err := dal.QueryUserInfoById(s.ctx, uid)
		if err != nil {
			return nil, err
		}
		message := "暂无聊天消息"
		msgType := constant.MessageTypeReceived
		latestMessage, err := mongodb.GetLatestMessage(s.ctx, req.UserId, uid)
		if err != nil {
			if !errors.Is(err, mongo.ErrNoDocuments) {
				return nil, err
			}
		}
		if latestMessage != nil {
			message = latestMessage.Content
			if latestMessage.Receiver == req.UserId {
				msgType = constant.MessageTypeSend
			}
		}

		userList[i] = &relation.FriendUser{
			Id:      userInfo.Id,
			Name:    userInfo.Nickname,
			Avatar:  &userInfo.Avatar,
			Message: &message,
			MsgType: int64(msgType),
		}
	}
	return userList, nil
}
