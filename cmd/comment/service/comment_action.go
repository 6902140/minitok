package service

import (
	"context"
	"time"

	"minitok/cmd/comment/dal/db"
	"minitok/cmd/comment/global"
	"minitok/kitex_gen/douyin/comment"
	"minitok/kitex_gen/douyin/user"
	"minitok/model"
	"minitok/pkg/constant"
	"minitok/pkg/errno"
	"minitok/pkg/jwt"
)

type CommentActionService struct {
	ctx context.Context
}

func NewCommentActionService(ctx context.Context) *CommentActionService {
	return &CommentActionService{ctx: ctx}
}

func (s *CommentActionService) CommentAction(req *comment.ActionRequest) (*comment.Comment, error) {
	claims, err := jwt.NewJWT(global.Configs.JWT.SigningKey).ParseToken(req.Token)
	if err != nil {
		return nil, err
	}
	if claims.Id == 0 || claims.Issuer != global.Configs.JWT.Issuer || claims.Subject != global.Configs.JWT.Subject {
		return nil, errno.AuthorizationFailedErr
	}
	userInfo, err := db.QueryFirstUserInfoByID(s.ctx, claims.Id, "nickname")
	if err != nil {
		return nil, err
	}
	if req.ActionType == constant.CommentActionPublish {
		if req.CommentText == nil {
			return nil, errno.ParamErr
		}
		newComment := &model.Comment{
			AuthorId:   claims.Id,
			VideoId:    req.VideoId,
			Content:    *req.CommentText,
			CreateTime: time.Now().Format("2006/01/02 15:04:05"),
		}
		err := db.CreateCommentInfos(s.ctx, []*model.Comment{newComment})
		if err != nil {
			return nil, err
		}
		return &comment.Comment{
			Id:         newComment.Id,
			User:       &user.User{Id: claims.Id, Name: userInfo.Nickname},
			Content:    newComment.Content,
			CreateDate: newComment.CreateTime,
		}, nil
	} else if req.ActionType == constant.CommentActionDelete {
		if req.CommentId == nil {
			return nil, errno.ParamErr
		}
		return nil, db.DeleteCommentInfo(s.ctx, *req.CommentId, req.VideoId)
	} else {
		return nil, errno.ParamErr
	}
}
