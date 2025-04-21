package model

import (
	"common/global"
	"time"
)

// 文章内容表
type EbArticleContent struct {
	Nid         uint32    `gorm:"column:nid;type:int UNSIGNED;comment:文章id;primaryKey;not null;" json:"nid"`                                    // 文章id
	Content     string    `gorm:"column:content;type:text;comment:文章内容;not null;" json:"content"`                                               // 文章内容
	CreatedTime time.Time `gorm:"column:created_time;type:datetime(3);comment:添加时间;not null;default:CURRENT_TIMESTAMP(3);" json:"created_time"` // 添加时间
}

func (c *EbArticleContent) Table() string {
	return "eb_article_content"
}
func (c *EbArticleContent) CreateEbArticleContent() bool {
	err := global.DB.Table("eb_article_content").Create(&c).Error
	if err != nil {

		return false
	}

	return true

}
