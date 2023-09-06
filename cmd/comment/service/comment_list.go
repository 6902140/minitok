package service

import (
	"context"

	"minitok/cmd/comment/dal/db"
	"minitok/cmd/comment/global"
	"minitok/kitex_gen/douyin/comment"
	"minitok/kitex_gen/douyin/user"
	"minitok/pkg/constant"
)

type CommentListService struct {
	ctx context.Context
}

func NewCommentListService(ctx context.Context) *CommentListService {
	return &CommentListService{ctx: ctx}
}

func (s *CommentListService) CommentList(req *comment.ListRequest) ([]*comment.Comment, error) {
	commentInfos, err := db.QueryCommentInfos(s.ctx, req.VideoId, constant.MaxQueryCommentNum, "id, author_id, content, create_time")
	if err != nil {
		return nil, err
	}
	res := make([]*comment.Comment, len(commentInfos))
	for i, commentInfo := range commentInfos {
		userInfo, err := db.QueryFirstUserInfoByID(s.ctx, commentInfo.AuthorId, "nickname, avatar")
		if err != nil {
			return nil, err
		}
		if len(userInfo.Avatar) == 0 {
			userInfo.Avatar = global.Configs.StaticResource.DefaultAvatar
		}
		res[i] = &comment.Comment{
			Id:         commentInfo.Id,
			User:       &user.User{Id: commentInfo.AuthorId, Name: userInfo.Nickname, Avatar: &userInfo.Avatar},
			Content:    commentInfo.Content,
			CreateDate: commentInfo.CreateTime,
		}
	}
	return res, nil
}
