package article

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// ArticleContent 文章内容点赞表
type ArticleUserCommentLikeCount struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`          // MongoDB 自动生成的唯一标识符
	CommentID   int64              `bson:"comment_id,omitempty"`   // 文章id
	UserID      int64              `bson:"user_id,omitempty"`      // 用户 ID，关联到用户表
	CreatedTime time.Time          `bson:"created_time,omitempty"` // 添加时间
	IsDel       int64              `bson:"is_del,omitempty"`       // 1未删2已删
	UpdateTime  time.Time          `bson:"update_time,omitempty"`  // 更新时间
}
