package model

import (
	"common/global"
	"gorm.io/gorm"
)

// Combination 拼团商品表
type Combination struct {
	Id            int     `gorm:"column:id;type:int UNSIGNED;primaryKey;not null;" json:"id"`
	ProductId     int     `gorm:"column:product_id;type:int UNSIGNED;comment:商品id;not null;" json:"product_id"`       // 商品id
	MerId         int     `gorm:"column:mer_id;type:int UNSIGNED;comment:商户id;default:0;" json:"mer_id"`              // 商户id
	Image         string  `gorm:"column:image;type:varchar(255);comment:推荐图;not null;" json:"image"`                  // 推荐图
	Images        string  `gorm:"column:images;type:varchar(2000);comment:轮播图;not null;" json:"images"`               // 轮播图
	Title         string  `gorm:"column:title;type:varchar(255);comment:活动标题;not null;" json:"title"`                 // 活动标题
	Attr          string  `gorm:"column:attr;type:varchar(255);comment:活动属性;default:NULL;" json:"attr"`               // 活动属性
	People        int     `gorm:"column:people;type:int UNSIGNED;comment:参团人数;not null;" json:"people"`               // 参团人数
	Info          string  `gorm:"column:info;type:varchar(255);comment:简介;not null;" json:"info"`                     // 简介
	Price         float64 `gorm:"column:price;type:decimal(10, 2) UNSIGNED;comment:价格;not null;" json:"price"`        // 价格
	Sort          int     `gorm:"column:sort;type:int UNSIGNED;comment:排序;not null;" json:"sort"`                     // 排序
	Sales         int     `gorm:"column:sales;type:int UNSIGNED;comment:销量;not null;default:0;" json:"sales"`         // 销量
	Stock         int     `gorm:"column:stock;type:int UNSIGNED;comment:库存;not null;" json:"stock"`                   // 库存
	AddTime       string  `gorm:"column:add_time;type:varchar(128);comment:添加时间;not null;" json:"add_time"`           // 添加时间
	IsHost        uint8   `gorm:"column:is_host;type:tinyint UNSIGNED;comment:推荐;not null;default:0;" json:"is_host"` // 推荐
	IsShow        uint8   `gorm:"column:is_show;type:tinyint UNSIGNED;comment:商品状态;not null;" json:"is_show"`         // 商品状态
	IsDel         uint8   `gorm:"column:is_del;type:tinyint UNSIGNED;not null;default:0;" json:"is_del"`
	Combination   uint8   `gorm:"column:combination;type:tinyint UNSIGNED;not null;default:1;" json:"combination"`
	MerUse        uint8   `gorm:"column:mer_use;type:tinyint UNSIGNED;comment:商户是否可用1可用0不可用;default:NULL;" json:"mer_use"`   // 商户是否可用1可用0不可用
	IsPostage     uint8   `gorm:"column:is_postage;type:tinyint UNSIGNED;comment:是否包邮1是0否;not null;" json:"is_postage"`      // 是否包邮1是0否
	Postage       float64 `gorm:"column:postage;type:decimal(10, 2) UNSIGNED;comment:邮费;not null;" json:"postage"`           // 邮费
	StartTime     int     `gorm:"column:start_time;type:int UNSIGNED;comment:拼团开始时间;not null;" json:"start_time"`            // 拼团开始时间
	StopTime      int     `gorm:"column:stop_time;type:int UNSIGNED;comment:拼团结束时间;not null;" json:"stop_time"`              // 拼团结束时间
	EffectiveTime int32   `gorm:"column:effective_time;type:int;comment:拼团订单有效时间;not null;default:0;" json:"effective_time"` // 拼团订单有效时间
	Cost          int     `gorm:"column:cost;type:int UNSIGNED;comment:拼图商品成本;not null;default:0;" json:"cost"`              // 拼图商品成本
	Browse        int32   `gorm:"column:browse;type:int;comment:浏览量;default:0;" json:"browse"`                               // 浏览量
	UnitName      string  `gorm:"column:unit_name;type:varchar(32);comment:单位名;not null;" json:"unit_name"`                  // 单位名
	TempId        int32   `gorm:"column:temp_id;type:int;comment:运费模板ID;default:NULL;" json:"temp_id"`                       // 运费模板ID
	Weight        float64 `gorm:"column:weight;type:decimal(8, 2);comment:重量;default:0.00;" json:"weight"`                   // 重量
	Volume        float64 `gorm:"column:volume;type:decimal(8, 2);comment:体积;default:0.00;" json:"volume"`                   // 体积
	Num           int32   `gorm:"column:num;type:int;comment:单次购买数量;default:NULL;" json:"num"`                               // 单次购买数量
	Quota         int32   `gorm:"column:quota;type:int;comment:限购总数;not null;default:0;" json:"quota"`                       // 限购总数
	QuotaShow     int32   `gorm:"column:quota_show;type:int;comment:限量总数显示;not null;default:0;" json:"quota_show"`           // 限量总数显示
}

func (c *Combination) TableName() string {
	return "combination"
}

func (c *Combination) GetCombinationList() (result []*Combination, err error) {
	err = global.DB.Find(&result).Error
	if err != nil {
		return nil, err
	}
	return
}
func (c *Combination) Add() error {
	return global.DB.Table("combination").Create(&c).Error
}

func (c *Combination) GetCombinationById(pid int64) (result *Combination, err error) {
	err = global.DB.Table("combination").Where("id = ? and is_del = 0", pid).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return
}
func (c *Combination) UpdateCombinationStock(id, num int64) error {
	return global.DB.Debug().Table("combination").Model(&Combination{}).Where("id = ?", id).Limit(1).Update("stock", gorm.Expr("stock - ?", num)).Error
}

func (c *Combination) ReverseCombinationStock(id int64, num int64) error {
	return global.DB.Debug().Table("combination").Model(&Combination{}).Where("id = ?", id).Limit(1).Update("stock", gorm.Expr("stock + ?", num)).Error
}
