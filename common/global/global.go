package global

import (
	"common/config"
	"context"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	Config config.AppViper
	NaCos  config.T
	DB     *gorm.DB
	Rdb    *redis.Client
	Ctx    = context.Background()
	ES     *elasticsearch.Client
)
