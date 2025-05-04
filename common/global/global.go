package global

import (
	"common/config"
	"context"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"

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
	IMGName           = "user_%d_order_%d"
)

var Order struct {
	Id      int64  `json:"id"`
	OrderSn string `json:"order_sn"`
	Uid     int64  `json:"uid"`
	Paid    int64  `json:"paid"`
	Status  int64  `json:"status"`
}

type OrderProduct struct {
	Price       float64
	ProductName string
	Postage     float64
	Image       string
	IsShow      int64
}
