package article

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// UserComment 用户评论表
type ArticleUserComment struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`           // MongoDB 自动生成的唯一标识符
	CommentID   int64              `bson:"comment_id,omitempty"`    // 评论的唯一 ID
	UserID      int64              `bson:"user_id,omitempty"`       // 用户 ID，关联到用户表
	ArticleID   int64              `bson:"article_id,omitempty"`    // 文章 ID，关联到 Article 表
	Content     string             `bson:"content,omitempty"`       // 评论内容
	Pid         int64              `bson:"pid,omitempty"`           // 父级评论 ID，0 表示顶级评论
	ReplyUserID int64              `bson:"reply_user_id,omitempty"` // 被回复的用户 ID
	Status      int64              `bson:"status,omitempty"`        // 评论状态：1 未删除，2 已删除
	IsDel       int64              `bson:"is_del,omitempty"`        // 是否删除：1 未删除，2 已删除
	AddTime     time.Time          `bson:"add_time,omitempty"`      // 添加时间
	UpdateTime  time.Time          `bson:"update_time,omitempty"`   // 更新时间
}
