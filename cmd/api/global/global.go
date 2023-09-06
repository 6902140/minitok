package global

import (
	"github.com/spf13/viper"
	"minitok/cmd/api/config"
	"minitok/kitex_gen/douyin/comment/commentservice"
	"minitok/kitex_gen/douyin/favorite/favoriteservice"
	"minitok/kitex_gen/douyin/feed/feedservice"
	"minitok/kitex_gen/douyin/message/messageservice"
	"minitok/kitex_gen/douyin/publish/publishservice"
	"minitok/kitex_gen/douyin/relation/relationservice"
	"minitok/kitex_gen/douyin/user/userservice"
)

var (
	Configs config.ServiceConfigs
	Viper   *viper.Viper

	UserServiceClient     *userservice.Client
	PublishServiceClient  *publishservice.Client
	FeedServiceClient     *feedservice.Client
	FavoriteServiceClient *favoriteservice.Client
	CommentServiceClient  *commentservice.Client
	MessageServiceClient  *messageservice.Client
	RelationServiceClient *relationservice.Client
)
