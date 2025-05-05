package article

import (
	"time"
)

// Article 文章管理表
type Article struct {
	ID            int64     `bson:"id,omitempty"`             // 文章管理ID
	IntID         int64     `bson:"int_id,omitempty"`         // 额外 int64 类型 ID 字段
	Cid           int64     `bson:"cid,omitempty"`            // 分类id
	Title         string    `bson:"title,omitempty"`          // 文章标题
	Author        string    `bson:"author,omitempty"`         // 文章作者
	ImageInput    string    `bson:"image_input,omitempty"`    // 文章图片
	Synopsis      string    `bson:"synopsis,omitempty"`       // 文章简介
	ShareTitle    string    `bson:"share_title,omitempty"`    // 文章分享标题
	ShareSynopsis string    `bson:"share_synopsis,omitempty"` // 文章分享简介
	Visit         int64     `bson:"visit,omitempty"`          // 浏览次数
	LikeCount     int64     `bson:"like_count,omitempty"`     //点赞次数
	Url           string    `bson:"url,omitempty"`            // 原文链接
	Status        string    `bson:"status,omitempty"`         // 状态
	AddTime       time.Time `bson:"add_time,omitempty"`       // 添加时间
	Hide          string    `bson:"hide,omitempty"`           // 是否隐藏
	AdminId       uint64    `bson:"admin_id,omitempty"`       // 管理员id
	MerId         int64     `bson:"mer_id,omitempty"`         // 商户id
	ProductId     int64     `bson:"product_id,omitempty"`     // 商品关联id
	IsDel         int64     `bson:"is_del,omitempty"`         // 1未删除2已删除
	UpdatedAt     time.Time `bson:"updated_at,omitempty"`     // 修改时间
}
