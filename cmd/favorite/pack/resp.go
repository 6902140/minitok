package pack

import (
	"errors"
	"fmt"

	"minitok/cmd/favorite/global"
	"minitok/kitex_gen/douyin/favorite"
	"minitok/kitex_gen/douyin/feed"
	"minitok/kitex_gen/douyin/user"
	"minitok/model"
	"minitok/pkg/errno"
)

func BuildActionResp(err error) *favorite.ActionResponse {
	if err == nil {
		return actionResp(errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return actionResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return actionResp(s)
}

func actionResp(err errno.ErrNo) *favorite.ActionResponse {
	return &favorite.ActionResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg}
}

func BuildListResp(videoList []*feed.Video, err error) *favorite.ListResponse {
	if err == nil {
		return listResp(videoList, errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return listResp(nil, e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return listResp(nil, s)
}

func listResp(videoList []*feed.Video, err errno.ErrNo) *favorite.ListResponse {
	return &favorite.ListResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg, VideoList: videoList}
}

func BuildRespVideo(videoInfo *model.Video, userInfo *model.User, isFollow bool, isFavorite bool) *feed.Video {
	if videoInfo == nil || userInfo == nil {
		return nil
	}
	return &feed.Video{
		Id:            videoInfo.Id,
		Author:        buildRespUser(userInfo, isFollow),
		Title:         videoInfo.Title,
		PlayUrl:       fmt.Sprintf("%s/%s", global.Configs.FileAccess.NginxUrl, videoInfo.VideoPath),
		CoverUrl:      fmt.Sprintf("%s/%s", global.Configs.FileAccess.NginxUrl, videoInfo.CoverPath),
		FavoriteCount: videoInfo.FavoriteCount,
		CommentCount:  videoInfo.CommentCount,
		IsFavorite:    isFavorite,
	}
}

func buildRespUser(userInfo *model.User, isFollow bool) *user.User {
	if userInfo == nil {
		return nil
	}
	return &user.User{
		Id:              userInfo.Id,
		Name:            userInfo.Nickname,
		IsFollow:        isFollow,
		Avatar:          &userInfo.Avatar,
		BackgroundImage: &userInfo.BackgroundImage,
		Signature:       &userInfo.Signature,
		FollowCount:     &userInfo.FollowCount,
		FollowerCount:   &userInfo.FollowerCount,
		TotalFavorited:  &userInfo.TotalFavorited,
		FavoriteCount:   &userInfo.FavoriteCount,
		WorkCount:       &userInfo.WorkCount,
	}
}
