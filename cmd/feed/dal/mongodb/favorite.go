package mongodb

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"minitok/cmd/feed/global"
)

func GetFavoriteInfo(ctx context.Context, uid, toUid int64) (bool, error) {
	favoriteCollection := global.MongoClient.Database(global.Configs.MongoDB.Database).Collection("favorite")
	filter := bson.D{{"favorite_list", bson.A{toUid}}}
	projection := bson.D{{"uid", uid}}
	err := favoriteCollection.FindOne(ctx, filter, options.FindOne().SetProjection(projection)).Decode(&bson.M{})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
