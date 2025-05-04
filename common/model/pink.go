package model

import (
	"common/global"
	"context"
	"encoding/json"
	"fmt"
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

// UpdateGroupStatus 更新拼团状态到 MySQL 和 Redis
func (p *Pink) UpdateGroupStatus(ctx context.Context, key string, status int) error {
	// 更新 MySQL 中的拼团状态
	if err := global.DB.Table("pink").Model(&p).Where("order_id = ?", key).Update("status", status).Error; err != nil {
		return fmt.Errorf("更新 MySQL 拼团状态失败: %w", err)
	}

	// 从 Redis 中获取当前拼团信息
	groupInfoJSON, err := global.Rdb.Get(ctx, key).Result()
	if err != nil {
		return fmt.Errorf("获取 Redis 拼团信息失败: %w", err)
	}

	// 反序列化
	if err = json.Unmarshal([]byte(groupInfoJSON), p); err != nil {
		return fmt.Errorf("反序列化 Redis 拼团信息失败: %w", err)
	}

	// 更新拼团状态
	p.Status = status

	// 序列化更新后的拼团信息
	pinkJSON, err := json.Marshal(p)
	if err != nil {
		return fmt.Errorf("序列化更新后的 Redis 拼团信息失败: %w", err)
	}

	// 更新 Redis 中的拼团信息
	if err = global.Rdb.Set(ctx, key, pinkJSON, time.Hour).Err(); err != nil {
		return fmt.Errorf("更新 Redis 拼团信息失败: %w", err)
	}

	return nil
}
func (p *Pink) UpdateGroupNum(pinkId string, num int64) error {
	err := global.DB.Table("pink").Where("order_id = ?", pinkId).Update("current_num", p.CurrentNum+num).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *Pink) UpdateStatus(pinkId string, status int) error {
	return global.DB.Where("order_id = ?", pinkId).Update("status", status).Error
}
