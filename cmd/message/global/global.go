package global

import (
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"minitok/cmd/message/config"
)

var (
	Configs     config.ServiceConfigs
	MongoClient *mongo.Client
	Viper       *viper.Viper
)
