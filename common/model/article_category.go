package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// ArticleCategory 文章分类表

type ArticleCategory struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`      // 文章分类id
	IntID   int64              `bson:"int_id,omitempty"`   // 额外 int64 类型 ID 字段
	Pid     int64              `bson:"pid,omitempty"`      // 父级ID
	Title   string             `bson:"title,omitempty"`    // 文章分类标题
	Intr    string             `bson:"intr,omitempty"`     // 文章分类简介
	Image   string             `bson:"image,omitempty"`    // 文章分类图片
	Status  int64              `bson:"status,omitempty"`   // 状态1删除0未删除
	Sort    int64              `bson:"sort,omitempty"`     // 排序
	IsDel   int64              `bson:"is_del,omitempty"`   // 1未删除2删除
	AddTime time.Time          `bson:"add_time,omitempty"` // 添加时间
	Hidden  string             `bson:"hidden,omitempty"`   // 是否隐藏
}

//func (a *ArticleCategory) CreateArticleCategory() bool {
//	err := global.DB.Debug().Table("article_category").Create(&a).Error
//	if err != nil {
//		return false
//	}
//	return true
//}
