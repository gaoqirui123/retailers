package model

import (
	"common/global"
	"gorm.io/gorm"
)

// Product 商品表
type Product struct {
	Id           int64   `gorm:"column:id;type:mediumint;comment:商品id;primaryKey;not null;" json:"id"`                                     // 商品id
	MerId        int64   `gorm:"column:mer_id;type:int UNSIGNED;comment:商户Id(0为总后台管理员创建,不为0的时候是商户后台创建);not null;default:0;" json:"mer_id"` // 商户Id(0为总后台管理员创建,不为0的时候是商户后台创建)
	Image        string  `gorm:"column:image;type:varchar(256);comment:商品图片;not null;" json:"image"`                                       // 商品图片
	SliderImage  string  `gorm:"column:slider_image;type:varchar(2000);comment:轮播图;default:NULL;" json:"slider_image"`                     // 轮播图
	StoreName    string  `gorm:"column:store_name;type:varchar(128);comment:商品名称;not null;" json:"store_name"`                             // 商品名称
	StoreInfo    string  `gorm:"column:store_info;type:varchar(256);comment:商品简介;default:NULL;" json:"store_info"`                         // 商品简介
	Keyword      string  `gorm:"column:keyword;type:varchar(256);comment:关键字;default:NULL;" json:"keyword"`                                // 关键字
	BarCode      string  `gorm:"column:bar_code;type:varchar(15);comment:商品条码（一维码）;" json:"bar_code"`                                      // 商品条码（一维码）
	CateId       string  `gorm:"column:cate_id;type:varchar(64);comment:分类id;default:NULL;" json:"cate_id"`                                // 分类id
	Price        float64 `gorm:"column:price;type:decimal(8, 2) UNSIGNED;comment:商品价格;not null;default:0.00;" json:"price"`                // 商品价格
	VipPrice     float64 `gorm:"column:vip_price;type:decimal(8, 2) UNSIGNED;comment:会员价格;default:0.00;" json:"vip_price"`                 // 会员价格
	OtPrice      float64 `gorm:"column:ot_price;type:decimal(8, 2) UNSIGNED;comment:市场价;default:0.00;" json:"ot_price"`                    // 市场价
	Postage      float64 `gorm:"column:postage;type:decimal(8, 2) UNSIGNED;comment:邮费;default:0.00;" json:"postage"`                       // 邮费
	UnitName     string  `gorm:"column:unit_name;type:varchar(32);comment:单位名;default:NULL;" json:"unit_name"`                             // 单位名
	Sort         int64   `gorm:"column:sort;type:smallint;comment:排序;default:0;" json:"sort"`                                              // 排序
	Sales        int64   `gorm:"column:sales;type:mediumint UNSIGNED;comment:销量;default:0;" json:"sales"`                                  // 销量
	Stock        int64   `gorm:"column:stock;type:mediumint UNSIGNED;comment:库存;not null;default:0;" json:"stock"`                         // 库存
	IsShow       int64   `gorm:"column:is_show;type:tinyint(1);comment:状态（0：未上架，1：上架）;not null;default:1;" json:"is_show"`                 // 状态（0：未上架，1：上架）
	IsHot        int64   `gorm:"column:is_hot;type:tinyint(1);comment:是否热卖;default:0;" json:"is_hot"`                                      // 是否热卖
	IsBenefit    int64   `gorm:"column:is_benefit;type:tinyint(1);comment:是否优惠;default:0;" json:"is_benefit"`                              // 是否优惠
	IsBest       int64   `gorm:"column:is_best;type:tinyint(1);comment:是否精品;default:0;" json:"is_best"`                                    // 是否精品
	IsNew        int64   `gorm:"column:is_new;type:tinyint(1);comment:是否新品;default:0;" json:"is_new"`                                      // 是否新品
	AddTime      int64   `gorm:"column:add_time;type:int UNSIGNED;comment:添加时间;not null;" json:"add_time"`                                 // 添加时间
	IsPostage    int64   `gorm:"column:is_postage;type:tinyint UNSIGNED;comment:是否包邮;default:0;" json:"is_postage"`                        // 是否包邮
	IsDel        int64   `gorm:"column:is_del;type:tinyint UNSIGNED;comment:是否删除;default:0;" json:"is_del"`                                // 是否删除
	MerUse       int64   `gorm:"column:mer_use;type:tinyint UNSIGNED;comment:商户是否代理 0不可代理1可代理;default:0;" json:"mer_use"`                  // 商户是否代理 0不可代理1可代理
	GiveIntegral float64 `gorm:"column:give_integral;type:decimal(8, 2) UNSIGNED;comment:获得积分;not null;" json:"give_integral"`             // 获得积分
	Cost         float64 `gorm:"column:cost;type:decimal(8, 2) UNSIGNED;comment:成本价;default:NULL;" json:"cost"`                            // 成本价
	IsSeckill    int64   `gorm:"column:is_seckill;type:tinyint UNSIGNED;comment:秒杀状态 0 未开启 1已开启;default:0;" json:"is_seckill"`             // 秒杀状态 0 未开启 1已开启
	IsBargain    int64   `gorm:"column:is_bargain;type:tinyint UNSIGNED;comment:砍价状态 0未开启 1开启;default:NULL;" json:"is_bargain"`            // 砍价状态 0未开启 1开启
	IsGood       int64   `gorm:"column:is_good;type:tinyint(1);comment:是否优品推荐;default:0;" json:"is_good"`                                  // 是否优品推荐
	IsSub        int64   `gorm:"column:is_sub;type:tinyint(1);comment:是否单独分佣;default:0;" json:"is_sub"`                                    // 是否单独分佣
	Ficti        int64   `gorm:"column:ficti;type:mediumint;comment:虚拟销量;default:100;" json:"ficti"`                                       // 虚拟销量
	Browse       int64   `gorm:"column:browse;type:int;comment:浏览量;default:0;" json:"browse"`                                              // 浏览量
	CodePath     string  `gorm:"column:code_path;type:varchar(64);comment:商品二维码地址(用户小程序海报);" json:"code_path"`                             // 商品二维码地址(用户小程序海报)
	Activity     string  `gorm:"column:activity;type:varchar(255);comment:活动显示排序1=秒杀，2=砍价，3=拼团;" json:"activity"`                          // 活动显示排序1=秒杀，2=砍价，3=拼团
}

func (p *Product) GetProductIdBy(productId int64) error {
	return global.DB.Debug().Table("product").Where("id = ?", productId).Limit(1).Find(&p).Error
}

func (p *Product) UpdateProductStock(id, num int64) error {
	return global.DB.Debug().Table("product").Model(&Product{}).Where("id = ?", id).Limit(1).Update("stock", gorm.Expr("stock - ?", num)).Error
}

func (p *Product) Add() error {
	return global.DB.Debug().Table("product").Create(&p).Error
}

func (p *Product) GetProductById(productId int64) (result *Product, err error) {
	err = global.DB.Debug().Table("product").Where("id = ?", productId).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return
}

func (p *Product) ReverseProductStock(productId, stock int64) error {
	return global.DB.Debug().Table("product").Model(&Product{}).Where("id = ?", productId).Update("good_stock", gorm.Expr("good_stock + ?", stock)).Error
}
