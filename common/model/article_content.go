package model

import (
	"time"
)

// ArticleContent 文章内容表
type ArticleContent struct {
	Nid         int64     `bson:"nid,omitempty"`          // 文章id
	Content     string    `bson:"content,omitempty"`      // 文章内容
	CreatedTime time.Time `bson:"created_time,omitempty"` // 添加时间
	IsDel       int64     `bson:"is_del,omitempty"`       // 1未删2已删
}

//func (c *ArticleContent) CreateEbArticleContent() bool {
//	err := global.DB.Debug().Table("article_content").Create(&c).Error
//	if err != nil {
//		return false
//	}
//	return true
//}
