package initialize

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"minitok/cmd/feed/global"
)

var ctx = context.Background()

func Mongo() (*mongo.Client, error) {
	return mongo.Connect(ctx, options.Client().ApplyURI(global.Configs.MongoDB.Url()))
}
