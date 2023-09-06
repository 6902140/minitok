package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"minitok/cmd/relation/global"
	"minitok/model"
)

func GetLatestMessage(ctx context.Context, uid1, uid2 int64) (*model.MongoMessage, error) {
	messageCollection := global.MongoClient.Database(global.Configs.MongoDB.Database).Collection("message")
	filter := bson.M{
		"$or": []bson.M{
			{"sender": uid1, "receiver": uid2},
			{"sender": uid2, "receiver": uid1},
		},
	}
	var message model.MongoMessage
	err := messageCollection.FindOne(ctx, filter, options.FindOne().SetSort(bson.M{"_id": -1})).Decode(&message)
	if err != nil {
		return nil, err
	}
	return &message, nil
}
