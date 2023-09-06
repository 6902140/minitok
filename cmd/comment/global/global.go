package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"minitok/cmd/comment/config"
)

var (
	Configs     config.ServiceConfigs
	GormDB      *gorm.DB
	RedisClient *redis.Client
	Viper       *viper.Viper
	MongoDB     *mongo.Database
)
