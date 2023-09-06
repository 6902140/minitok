package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"minitok/cmd/publish/config"
	"time"
)

var (
	Configs     config.ServiceConfigs
	GormDB      *gorm.DB
	RedisClient *redis.Client
	MongoClient *mongo.Client
	Viper       *viper.Viper
)

var (
	ExpireDurationNullKey       time.Duration
	ExpireDurationUserBaseInfo  time.Duration
	ExpireDurationVideoBaseInfo time.Duration
)
