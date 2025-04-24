package global

import (
	"common/config"
	"context"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"gorm.io/gorm"
)

var (
	Config config.AppViper
	NaCos  config.T
	DB     *gorm.DB
	Rdb    *redis.Client
	Ctx    = context.Background()
	ES     *elasticsearch.Client
	MDB    *mongo.Client
)

const (
	TimeFormat        = "2006-01-02 15:04:05"
	GroupBuyKeyPrefix = "group_buy:"
)
