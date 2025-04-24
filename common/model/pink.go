package model

import (
	"common/global"
	"context"
	"encoding/json"
	"time"
)

// Pink 拼团表
type Pink struct {
	Id         int     `gorm:"column:id;type:int UNSIGNED;primaryKey;not null;" json:"id"`
	Uid        int     `gorm:"column:uid;type:int UNSIGNED;comment:用户id;not null;" json:"uid"`                             // 用户id
	OrderId    string  `gorm:"column:order_id;type:varchar(64);comment:订单id 生成;not null;" json:"order_id"`                 // 订单id 生成
	OrderIdKey string  `gorm:"column:order_id_key;type:varchar(64);comment:订单id  数据库;not null;" json:"order_id_key"`       // 订单id  数据库
	TotalNum   int     `gorm:"column:total_num;type:int UNSIGNED;comment:购买商品个数;not null;" json:"total_num"`               // 购买商品个数
	TotalPrice float64 `gorm:"column:total_price;type:decimal(10, 2) UNSIGNED;comment:购买总金额;not null;" json:"total_price"` // 购买总金额
	Cid        int     `gorm:"column:cid;type:int UNSIGNED;comment:拼团商品id;not null;" json:"cid"`                           // 拼团商品id
	Pid        int     `gorm:"column:pid;type:int UNSIGNED;comment:商品id;not null;" json:"pid"`                             // 商品id
	People     int64   `gorm:"column:people;type:int UNSIGNED;comment:拼团总人数;not null;" json:"people"`                      // 拼团总人数
	CurrentNum int64   `gorm:"column:current_num;type:int;comment:参与拼团的人数;not null;" json:"current_num"`                   // 参与拼团的人数
	Price      float64 `gorm:"column:price;type:decimal(10, 2) UNSIGNED;comment:拼团商品单价;not null;" json:"price"`            // 拼团商品单价
	AddTime    string  `gorm:"column:add_time;type:varchar(24);comment:开始时间;not null;" json:"add_time"`                    // 开始时间
	StopTime   string  `gorm:"column:stop_time;type:varchar(24);not null;" json:"stop_time"`
	KId        int     `gorm:"column:k_id;type:int UNSIGNED;comment:团长id 0为团长;not null;default:0;" json:"k_id"`                    // 团长id 0为团长
	IsTpl      int     `gorm:"column:is_tpl;type:tinyint UNSIGNED;comment:是否发送模板消息0未发送1已发送;not null;default:0;" json:"is_tpl"`     // 是否发送模板消息0未发送1已发送
	IsRefund   int     `gorm:"column:is_refund;type:tinyint UNSIGNED;comment:是否退款 0未退款 1已退款;not null;default:0;" json:"is_refund"` // 是否退款 0未退款 1已退款
	Status     int     `gorm:"column:status;type:tinyint UNSIGNED;comment:状态1进行中2已完成3未完成;not null;default:1;" json:"status"`       // 状态1进行中2已完成3未完成
}

func (p *Pink) Create() error {
	return global.DB.Table("pink").Create(&p).Error
}

// updateGroupStatus 更新拼团状态
func (p *Pink) UpdateGroupStatus(key string, status int) error {
	var pink Pink
	groupInfoJSON, err := global.Rdb.Get(context.Background(), key).Result()
	if err != nil {
		return err
	}
	if err = json.Unmarshal([]byte(groupInfoJSON), &pink); err != nil {
		return err
	}
	pink.Status = status
	pinkJSON, err := json.Marshal(pink)
	if err != nil {
		return err
	}
	return global.Rdb.Set(context.Background(), key, pinkJSON, time.Hour).Err()
}
