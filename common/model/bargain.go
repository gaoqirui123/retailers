package model

import (
	"common/global"
)

// 砍价商品信息表
type Bargain struct {
	Id              uint32  `gorm:"column:id;type:int;comment:砍价商品ID;primaryKey;not null;" json:"id"`                                        // 砍价商品ID
	ProductId       uint32  `gorm:"column:product_id;type:int;comment:关联商品ID;not null;" json:"product_id"`                                   // 关联商品ID
	Title           string  `gorm:"column:title;type:varchar(255);comment:砍价活动名称;not null;" json:"title"`                                    // 砍价活动名称
	Image           string  `gorm:"column:image;type:varchar(150);comment:砍价活动图片;not null;" json:"image"`                                    // 砍价活动图片
	UnitName        string  `gorm:"column:unit_name;type:varchar(16);comment:单位名称;default:NULL;" json:"unit_name"`                           // 单位名称
	Stock           int     `gorm:"column:stock;type:int;comment:库存;default:NULL;" json:"stock"`                                             // 库存
	Sales           int     `gorm:"column:sales;type:int;comment:销量;default:NULL;" json:"sales"`                                             // 销量
	Images          string  `gorm:"column:images;type:varchar(2000);comment:砍价商品轮播图;not null;" json:"images"`                                // 砍价商品轮播图
	StartTime       int     `gorm:"column:start_time;type:int;comment:砍价开启时间;not null;" json:"start_time"`                                   // 砍价开启时间
	StopTime        int     `gorm:"column:stop_time;type:int;comment:砍价结束时间;not null;" json:"stop_time"`                                     // 砍价结束时间
	StoreName       string  `gorm:"column:store_name;type:varchar(255);comment:砍价商品名称;default:NULL;" json:"store_name"`                      // 砍价商品名称
	Price           float64 `gorm:"column:price;type:decimal(8, 2);comment:砍价金额;default:NULL;" json:"price"`                                 // 砍价金额
	MinPrice        float64 `gorm:"column:min_price;type:decimal(8, 2);comment:砍价商品最低价;default:NULL;" json:"min_price"`                      // 砍价商品最低价
	Num             int     `gorm:"column:num;type:int;comment:每次购买的砍价商品数量;default:NULL;" json:"num"`                                        // 每次购买的砍价商品数量
	BargainMaxPrice float64 `gorm:"column:bargain_max_price;type:decimal(8, 2);comment:用户每次砍价的最大金额;default:NULL;" json:"bargain_max_price"`  // 用户每次砍价的最大金额
	BargainMinPrice float64 `gorm:"column:bargain_min_price;type:decimal(8, 2);comment:用户每次砍价的最小金额;default:NULL;" json:"bargain_min_price"`  // 用户每次砍价的最小金额
	BargainNum      int     `gorm:"column:bargain_num;type:int;comment:用户每次砍价的次数;not null;default:1;" json:"bargain_num"`                    // 用户每次砍价的次数
	Status          uint8   `gorm:"column:status;type:tinyint;comment:砍价状态 0(到砍价时间不自动开启)  1(到砍价时间自动开启时间);not null;default:1;" json:"status"` // 砍价状态，0表示到砍价时间不自动开启，1表示到砍价时间自动开启
	GiveIntegral    float64 `gorm:"column:give_integral;type:decimal(10, 2);comment:反多少积分;not null;default:0.00;" json:"give_integral"`      // 反多少积分
	Info            string  `gorm:"column:info;type:varchar(255);comment:砍价活动简介;default:NULL;" json:"info"`                                  // 砍价活动简介
	Cost            float64 `gorm:"column:cost;type:decimal(8, 2);comment:成本价;default:NULL;" json:"cost"`                                    // 成本价
	Sort            int     `gorm:"column:sort;type:int;comment:排序;not null;default:0;" json:"sort"`                                         // 排序
	IsHot           uint8   `gorm:"column:is_hot;type:tinyint;comment:是否推荐0不推荐1推荐;not null;default:0;" json:"is_hot"`                        // 是否推荐，0表示不推荐，1表示推荐
	IsDel           uint8   `gorm:"column:is_del;type:tinyint;comment:是否删除 0未删除 1删除;not null;default:0;" json:"is_del"`                      // 是否删除，0表示未删除，1表示删除
	AddTime         int     `gorm:"column:add_time;type:int;comment:添加时间;default:NULL;" json:"add_time"`                                     // 添加时间
	IsPostage       uint8   `gorm:"column:is_postage;type:tinyint;comment:是否包邮 0不包邮 1包邮;not null;default:1;" json:"is_postage"`              // 是否包邮，0表示不包邮，1表示包邮
	Postage         float64 `gorm:"column:postage;type:decimal(10, 2);comment:邮费;default:NULL;" json:"postage"`                              // 邮费
	Rule            string  `gorm:"column:rule;type:text;comment:砍价规则;" json:"rule"`                                                         // 砍价规则
	Look            int     `gorm:"column:look;type:int;comment:砍价商品浏览量;default:0;" json:"look"`                                             // 砍价商品浏览量
	Share           int     `gorm:"column:share;type:int;comment:砍价商品分享量;default:0;" json:"share"`                                           // 砍价商品分享量
	TempId          int32   `gorm:"column:temp_id;type:int;comment:运费模板ID;default:NULL;" json:"temp_id"`                                     // 运费模板ID
	Weight          float64 `gorm:"column:weight;type:decimal(8, 2);comment:重量;default:0.00;" json:"weight"`                                 // 重量
	Volume          float64 `gorm:"column:volume;type:decimal(8, 2);comment:体积;default:0.00;" json:"volume"`                                 // 体积
	Quota           int32   `gorm:"column:quota;type:int;comment:限购总数;not null;default:0;" json:"quota"`                                     // 限购总数
	QuotaShow       int32   `gorm:"column:quota_show;type:int;comment:限量总数显示;not null;default:0;" json:"quota_show"`                         // 限量总数显示
}

func (m *Bargain) TableName() string {
	return "bargain"
}

// 创建砍价活动记录
func (m *Bargain) BargainCreate() error {
	return global.DB.Create(&m).Error
}

// 砍价商品表详情
func (m *Bargain) BargainShow(Id uint32) error {
	return global.DB.Where("product_id = ?", Id).Find(&m).Error
}

// 砍价商品表ID详情
func (m *Bargain) BargainShowID(Id uint32) error {
	return global.DB.Where("id = ?", Id).Find(&m).Error
}

// 砍价商品表列表
func (m *Bargain) BargainList() (b []*Bargain, err error) {
	err = global.DB.Find(&b).Error
	if err != nil {
		return nil, err
	}
	return b, err
}

// 修改砍价商品表是否删除
func (m *Bargain) BargainUpdate() error {
	return global.DB.Model(&m).Where("product_id = ?", m.ProductId).Update("is_del", m.IsDel).Error
}

func (m *Bargain) GetBargainIdBy(id int64) error {
	return global.DB.Where("product_id = ? and is_del = 0", id).Limit(1).Find(&m).Error
}
