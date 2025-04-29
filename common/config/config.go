package config

type AppViper struct {
	NaCos
	UserSrv
	ProductSrv
	OrderSrv
	UserEnterSrv
	ArticleSrv
	AdministratorsSrv
	CartSrv
	DistributionSrv
	RandomString
	AliPay
}
type NaCos struct {
	NameSpace string
	DataId    string
	Group     string
	Host      string
	Port      uint64
}
type UserSrv struct {
	Host string
	Port int
}
type ProductSrv struct {
	Host string
	Port int
}
type OrderSrv struct {
	Host string
	Port int
}
type UserEnterSrv struct {
	Host string
	Port int
}
type ArticleSrv struct {
	Host string
	Port int
}
type AdministratorsSrv struct {
	Host string
	Port int
}
type CartSrv struct {
	Host string
	Port int
}
type DistributionSrv struct {
	Host string
	Port int
}

// 邀请码长度
type RandomString struct {
	Log int
}

type AliPay struct {
	APPID     string
	NotifyURL string
	ReturnURL string
	Key       string
}

type T struct {
	Mysql struct {
		User     string `json:"user"`
		Pass     string `json:"pass"`
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Database string `json:"database"`
	} `json:"mysql"`
	Redis struct {
		Addr string `json:"addr"`
		Pass string `json:"pass"`
		Db   int    `json:"db"`
	} `json:"redis"`
	Elasticsearch struct {
		Addr string `json:"addr"`
	} `json:"elasticsearch"`
	Mongodb struct {
		User     string `json:"user"`
		Pass     string `json:"pass"`
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Database string `json:"database"`
	} `json:"mongodb"`
	Consul struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	} `json:"consul"`
	Aliyun struct {
		AccessKeyId     string `json:"accessKeyId"`
		AccessKeySecret string `json:"accessKeySecret"`
	} `json:"aliyun"`
	RealName struct {
		SecretId  string `json:"secretId"`
		SecretKey string `json:"secretKey"`
	} `json:"realName"`
	Gaode struct {
		Key string `json:"key"`
	} `json:"gaode"`
}
