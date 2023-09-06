package main

import (
	"context"

	"minitok/cmd/feed/pack"
	"minitok/cmd/feed/service"
	feed "minitok/kitex_gen/douyin/feed"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// Feed implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) Feed(ctx context.Context, req *feed.FeedRequest) (*feed.FeedResponse, error) {
	videoList, nextTime, err := service.NewFeedService(ctx).Feed(req)
	if err != nil {
		return pack.BuildFeedResp(nil, 0, err), nil
	}
	return pack.BuildFeedResp(videoList, nextTime, nil), nil
}
