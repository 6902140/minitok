package mongodb

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"minitok/cmd/feed/global"
)

func GetFollowInfo(ctx context.Context, uid, toUid int64) (bool, error) {
	relationCollection := global.MongoClient.Database(global.Configs.MongoDB.Database).Collection("relation")
	filter := bson.D{{"follow_list", bson.A{toUid}}}
	projection := bson.D{{"uid", uid}}
	err := relationCollection.FindOne(ctx, filter, options.FindOne().SetProjection(projection)).Decode(&bson.M{})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
