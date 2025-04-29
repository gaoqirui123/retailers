package model

import "common/global"

type Seckill struct {
	Id           int64   `gorm:"column:id;type:int UNSIGNED;comment:商品秒杀商品表id;primaryKey;not null;" json:"id"`                           // 商品秒杀商品表id
	ProductId    int64   `gorm:"column:product_id;type:int UNSIGNED;comment:商品id;not null;" json:"product_id"`                           // 商品id
	Image        string  `gorm:"column:image;type:varchar(255);comment:推荐图;default:NULL;" json:"image"`                                  // 推荐图
	Images       string  `gorm:"column:images;type:varchar(2000);comment:轮播图;default:NULL;" json:"images"`                               // 轮播图
	Name         string  `gorm:"column:name;type:varchar(255);comment:商品名称;not null;" json:"name"`                                       // 商品名称
	Info         string  `gorm:"column:info;type:varchar(255);comment:简介;not null;" json:"info"`                                         // 简介
	Price        float64 `gorm:"column:price;type:decimal(10, 2) UNSIGNED;comment:价格;not null;" json:"price"`                            // 价格
	Cost         float64 `gorm:"column:cost;type:decimal(8, 2) UNSIGNED;comment:成本;default:0.00;" json:"cost"`                           // 成本
	OtPrice      float64 `gorm:"column:ot_price;type:decimal(10, 2) UNSIGNED;comment:原价;not null;" json:"ot_price"`                      // 原价
	GiveIntegral float64 `gorm:"column:give_integral;type:decimal(10, 2) UNSIGNED;comment:返多少积分;default:NULL;" json:"give_integral"`     // 返多少积分
	Sort         int64   `gorm:"column:sort;type:int UNSIGNED;comment:排序;default:NULL;" json:"sort"`                                     // 排序
	Stock        int64   `gorm:"column:stock;type:int UNSIGNED;comment:库存;not null;" json:"stock"`                                       // 库存
	Sales        int64   `gorm:"column:sales;type:int UNSIGNED;comment:销量;not null;default:0;" json:"sales"`                             // 销量
	UnitName     string  `gorm:"column:unit_name;type:varchar(16);comment:单位名;default:NULL;" json:"unit_name"`                           // 单位名
	Postage      float64 `gorm:"column:postage;type:decimal(8, 2) UNSIGNED;comment:邮费;default:0.00;" json:"postage"`                     // 邮费
	Description  string  `gorm:"column:description;type:text;comment:内容;" json:"description"`                                            // 内容
	StartTime    string  `gorm:"column:start_time;type:varchar(128);comment:开始时间;not null;" json:"start_time"`                           // 开始时间
	StopTime     string  `gorm:"column:stop_time;type:varchar(128);comment:结束时间;not null;" json:"stop_time"`                             // 结束时间
	AddTime      string  `gorm:"column:add_time;type:varchar(128);comment:添加时间;default:NULL;" json:"add_time"`                           // 添加时间
	Status       int64   `gorm:"column:status;type:tinyint UNSIGNED;comment:商品状态（0：上架，1：下架）;not null;" json:"status"`                    // 商品状态（0：上架，1：下架）
	IsPostage    int64   `gorm:"column:is_postage;type:tinyint UNSIGNED;comment:是否包邮（0：包邮，1：不包邮）;not null;default:0;" json:"is_postage"` // 是否包邮（0：包邮，1：不包邮）
	IsHot        int64   `gorm:"column:is_hot;type:tinyint UNSIGNED;comment:热门推荐;default:0;" json:"is_hot"`                              // 热门推荐
	IsDel        int64   `gorm:"column:is_del;type:tinyint UNSIGNED;comment:删除 0未删除1已删除;default:0;" json:"is_del"`                       // 删除 0未删除1已删除
	Num          int64   `gorm:"column:num;type:int UNSIGNED;comment:最多秒杀几个;default:NULL;" json:"num"`                                   // 最多秒杀几个
	IsShow       int64   `gorm:"column:is_show;type:tinyint UNSIGNED;comment:显示;default:1;" json:"is_show"`                              // 显示
	TimeId       int64   `gorm:"column:time_id;type:int;comment:时间段ID;default:NULL;" json:"time_id"`                                     // 时间段ID
	TempId       int64   `gorm:"column:temp_id;type:int;comment:运费模板ID;default:NULL;" json:"temp_id"`                                    // 运费模板ID
	Weight       float64 `gorm:"column:weight;type:decimal(8, 2);comment:商品重量;default:0.00;" json:"weight"`                              // 商品重量
	Volume       float64 `gorm:"column:volume;type:decimal(8, 2);comment:商品体积;default:0.00;" json:"volume"`                              // 商品体积
	Quota        int64   `gorm:"column:quota;type:int;comment:限购总数;not null;default:0;" json:"quota"`                                    // 限购总数
	QuotaShow    int64   `gorm:"column:quota_show;type:int;comment:限购总数显示;default:0;" json:"quota_show"`                                 // 限购总数显示
}

func (s *Seckill) AddSeckillProduct() error {
	return global.DB.Debug().Table("seckill").Create(&s).Error
}
