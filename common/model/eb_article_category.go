package model

import (
	"common/global"
	"time"
)

// 文章分类表
type EbArticleCategory struct {
	Id      uint32    `gorm:"column:id;type:int UNSIGNED;comment:文章分类id;primaryKey;not null;" json:"id"`                            // 文章分类id
	Pid     int32     `gorm:"column:pid;type:int;comment:父级ID;not null;default:0;" json:"pid"`                                      // 父级ID
	Title   string    `gorm:"column:title;type:varchar(255);comment:文章分类标题;not null;" json:"title"`                                 // 文章分类标题
	Intr    string    `gorm:"column:intr;type:varchar(255);comment:文章分类简介;default:NULL;" json:"intr"`                               // 文章分类简介
	Image   string    `gorm:"column:image;type:varchar(255);comment:文章分类图片;not null;" json:"image"`                                 // 文章分类图片
	Status  uint8     `gorm:"column:status;type:tinyint UNSIGNED;comment:状态;not null;" json:"status"`                               // 状态
	Sort    uint32    `gorm:"column:sort;type:int UNSIGNED;comment:排序;not null;default:0;" json:"sort"`                             // 排序
	IsDel   uint8     `gorm:"column:is_del;type:tinyint UNSIGNED;comment:1删除0未删除;not null;default:0;" json:"is_del"`                // 1删除0未删除
	AddTime time.Time `gorm:"column:add_time;type:datetime(3);comment:添加时间;not null;default:CURRENT_TIMESTAMP(3);" json:"add_time"` // 添加时间
	Hidden  string    `gorm:"column:hidden;type:enum('隐藏', '未隐藏');comment:是否隐藏;not null;default:未隐藏;" json:"hidden"`                // 是否隐藏
}

func (a *EbArticleCategory) Table() string {
	return "eb_article_category"
}

func (a *EbArticleCategory) CreateEbArticleCategory() bool {
	err := global.DB.Table("eb_article_category").Create(&a).Error
	if err != nil {
		return false
	}

	return true
}
