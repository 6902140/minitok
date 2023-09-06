package main

import (
	"context"
	"minitok/cmd/publish/pack"
	"minitok/cmd/publish/service"
	publish "minitok/kitex_gen/douyin/publish"
	"minitok/pkg/errno"
)

// PublishServiceImpl implements the last service interface defined in the IDL.
type PublishServiceImpl struct{}

// PublishAction implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishAction(ctx context.Context, req *publish.ActionRequest) (*publish.ActionResponse, error) {
	if len(req.Title) == 0 || len(req.Token) == 0 || len(req.Data) == 0 {
		return pack.BuildActionResp(errno.ParamErr), nil
	}
	if err := service.NewPublishActionService(ctx).PublishAction(req); err != nil {
		return pack.BuildActionResp(err), nil
	}
	return pack.BuildActionResp(nil), nil
}

// PublishList implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishList(ctx context.Context, req *publish.ListRequest) (*publish.ListResponse, error) {
	if req.UserId == 0 {
		return pack.BuildListResp(nil, errno.ParamErr), nil
	}
	videoList, err := service.NewPublishListService(ctx).PublishList(req)
	if err != nil {
		return pack.BuildListResp(nil, err), nil
	}
	return pack.BuildListResp(videoList, nil), nil
}
