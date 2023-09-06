package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"minitok/cmd/favorite/config"
	"time"
)

var (
	Configs     config.ServiceConfigs
	GormDB      *gorm.DB
	RedisClient *redis.Client
	Viper       *viper.Viper
	MongoClient *mongo.Client
)

var (
	ExpireDurationNullKey       time.Duration
	ExpireDurationUserBaseInfo  time.Duration
	ExpireDurationVideoBaseInfo time.Duration
)
