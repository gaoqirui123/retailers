package model

import (
	"common/global"
	"time"
)

// ArticleContent 文章内容表
type ArticleContent struct {
	Nid         uint32    `gorm:"column:nid;type:int UNSIGNED;comment:文章id;primaryKey;not null;" json:"nid"`                                    // 文章id
	Content     string    `gorm:"column:content;type:text;comment:文章内容;not null;" json:"content"`                                               // 文章内容
	CreatedTime time.Time `gorm:"column:created_time;type:datetime(3);comment:添加时间;not null;default:CURRENT_TIMESTAMP(3);" json:"created_time"` // 添加时间
	IsDel       int64     `gorm:"column:is_del;type:bigint;comment:1未删2已删;default:1;" json:"is_del"`                                            // 1未删2已删
}

func (c *ArticleContent) CreateEbArticleContent() bool {
	err := global.DB.Debug().Table("article_content").Create(&c).Error
	if err != nil {
		return false
	}
	return true
}
