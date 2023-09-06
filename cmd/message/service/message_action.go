package service

import (
	"context"
	"minitok/cmd/message/dal/mongodb"
	"time"

	"minitok/cmd/message/global"
	"minitok/kitex_gen/douyin/message"
	"minitok/model"
	"minitok/pkg/constant"
	"minitok/pkg/errno"
	"minitok/pkg/jwt"
)

type MessageActionService struct {
	ctx context.Context
}

func NewMessageActionService(ctx context.Context) *MessageActionService {
	return &MessageActionService{ctx: ctx}
}

func (s *MessageActionService) MessageAction(req *message.ActionRequest) error {
	claims, err := jwt.NewJWT(global.Configs.JWT.SigningKey).ParseToken(req.Token)
	if err != nil {
		return err
	}
	if claims.Id == 0 || claims.Issuer != global.Configs.JWT.Issuer || claims.Subject != global.Configs.JWT.Subject {
		return errno.AuthorizationFailedErr
	}
	if req.ActionType != constant.MessageChatTypeSend {
		return errno.ParamErr
	}
	newMessage := &model.MongoMessage{
		Receiver:  req.ToUserId,
		Sender:    claims.Id,
		Content:   req.Content,
		Timestamp: time.Now().UnixNano() / 1e6,
	}
	return mongodb.AddMessageInfo(s.ctx, newMessage)
}
