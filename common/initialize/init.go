package initialize

import (
	"common/global"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-redis/redis/v8"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func init() {
	InitViper()
	InitNaCos()
	InitMysql()
	InitRedis()
	InitElasticsearch()
	MongoDbInit()

}

func InitViper() {
	viper.SetConfigFile("../../common/config/server.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&global.Config)
	if err != nil {
		panic(err)
	}
}

func InitNaCos() {

	Conf := global.Config.NaCos

	// 配置与 NaCos 服务的连接
	clientConfig := constant.ClientConfig{
		NamespaceId:         Conf.NameSpace, //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}
	// 配置 NaCos 服务的连接信息
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      Conf.Host, // NaCos 服务器 IP 地址
			ContextPath: "/nacos",  // NaCos 服务的上下文路径
			Port:        Conf.Port, // NaCos 服务器的端口
			Scheme:      "http",    // 使用的协议
		},
	}

	// 创建配置客户端
	configClient, _ := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs, // NaCos 服务器配置
		"clientConfig":  clientConfig,  // 客户端配置
	})

	// 从 NaCos 获取配置内容
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: Conf.DataId, // 配置的 DataId
		Group:  Conf.Group,  // 配置的 Group
	})

	err = json.Unmarshal([]byte(content), &global.NaCos)
	if err != nil {
		return
	}

	//err = configClient.ListenConfig(vo.ConfigParam{
	//	DataId: Conf.DataId, // 配置的 DataId
	//	Group:  Conf.Group,  // 配置的 Group
	//	OnChange: func(namespace, group, dataId, data string) {
	//		fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
	//		err = json.Unmarshal([]byte(data), &global.NaCos)
	//		if err != nil {
	//			return
	//		}
	//	},
	//})
	//fmt.Println(content)
}

func InitMysql() {
	var err error
	Conf := global.NaCos.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", Conf.User, Conf.Pass, Conf.Host, Conf.Port, Conf.Database)
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	sqlDB, err := global.DB.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err != nil {
		panic(err)
	}
	err = global.DB.AutoMigrate()
	if err != nil {
		panic(err)
	}
}
func InitRedis() {
	Conf := global.NaCos.Redis
	global.Rdb = redis.NewClient(&redis.Options{
		Addr:     Conf.Addr,
		Password: Conf.Pass, // no password set
		DB:       Conf.Db,   // use default DB
	})

	pong, err := global.Rdb.Ping(global.Ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(pong, err)
	// Output: PONG <nil>
}

func InitElasticsearch() {
	var err error
	Conf := global.NaCos.Elasticsearch
	cfg := elasticsearch.Config{
		Addresses: []string{
			Conf.Addr,
		},
		// ...
	}
	global.ES, err = elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}
}

func MongoDbInit() {
	var err error
	Conf := global.NaCos.Mongodb
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	global.MDB, err = mongo.Connect(options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%d", Conf.User, Conf.Pass, Conf.Host, Conf.Port)))
	if err != nil {
		return
	}
	err = global.MDB.Ping(ctx, nil)
	if err != nil {
		err = global.MDB.Disconnect(ctx)
		if err != nil {
			panic(err)
		}
		return
	}
	fmt.Println("mongodb --- success")

}
