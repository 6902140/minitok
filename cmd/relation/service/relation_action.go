package service

import (
	"context"
	"minitok/cmd/relation/dal/mongodb"

	"minitok/cmd/relation/global"
	"minitok/kitex_gen/douyin/relation"
	"minitok/pkg/constant"
	"minitok/pkg/errno"
	"minitok/pkg/jwt"
)

type RelationActionService struct {
	ctx context.Context
}

func NewRelationActionService(ctx context.Context) *RelationActionService {
	return &RelationActionService{ctx: ctx}
}

func (s *RelationActionService) RelationAction(req *relation.ActionRequest) error {
	claims, err := jwt.NewJWT(global.Configs.JWT.SigningKey).ParseToken(req.Token)
	if err != nil {
		return err
	}
	if claims.Id == 0 || claims.Issuer != global.Configs.JWT.Issuer || claims.Subject != global.Configs.JWT.Subject {
		return errno.AuthorizationFailedErr
	}
	if claims.Id == req.ToUserId {
		return errno.ParamErr
	}
	if req.ActionType == constant.RelationActionDo {
		return mongodb.AddRelationInfo(s.ctx, claims.Id, req.ToUserId)
	} else if req.ActionType == constant.RelationActionCancel {
		return mongodb.DeleteRelationInfo(s.ctx, claims.Id, req.ToUserId)
	} else {
		return errno.ParamErr
	}
}
