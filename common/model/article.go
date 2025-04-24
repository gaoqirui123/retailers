package model

import (
	"common/global"
	"time"
)

// Article 文章管理表
type Article struct {
	Id            uint32    `gorm:"column:id;type:int UNSIGNED;comment:文章管理ID;primaryKey;not null;" json:"id"`                            // 文章管理ID
	Cid           int32     `gorm:"column:cid;type:int;comment:分类id;default:1;" json:"cid"`                                               // 分类id
	Title         string    `gorm:"column:title;type:varchar(255);comment:文章标题;not null;" json:"title"`                                   // 文章标题
	Author        string    `gorm:"column:author;type:varchar(255);comment:文章作者;not null;" json:"author"`                                 // 文章作者
	ImageInput    string    `gorm:"column:image_input;type:varchar(255);comment:文章图片;not null;" json:"image_input"`                       // 文章图片
	Synopsis      string    `gorm:"column:synopsis;type:varchar(255);comment:文章简介;default:NULL;" json:"synopsis"`                         // 文章简介
	ShareTitle    string    `gorm:"column:share_title;type:varchar(255);comment:文章分享标题;default:NULL;" json:"share_title"`                 // 文章分享标题
	ShareSynopsis string    `gorm:"column:share_synopsis;type:varchar(255);comment:文章分享简介;default:NULL;" json:"share_synopsis"`           // 文章分享简介
	Visit         int64     `gorm:"column:visit;type:bigint;comment:浏览次数;default:NULL;" json:"visit"`                                     // 浏览次数
	Url           string    `gorm:"column:url;type:varchar(255);comment:原文链接;default:NULL;" json:"url"`                                   // 原文链接
	Status        string    `gorm:"column:status;type:enum('私密', '公开');comment:状态;not null;default:公开;" json:"status"`                    // 状态
	AddTime       time.Time `gorm:"column:add_time;type:datetime(3);comment:添加时间;not null;default:CURRENT_TIMESTAMP(3);" json:"add_time"` // 添加时间
	Hide          string    `gorm:"column:hide;type:enum('隐藏', '未隐藏');comment:是否隐藏;not null;default:未隐藏;" json:"hide"`                    // 是否隐藏
	AdminId       uint64    `gorm:"column:admin_id;type:bigint UNSIGNED;comment:管理员id;not null;default:0;" json:"admin_id"`               // 管理员id
	MerId         uint32    `gorm:"column:mer_id;type:int UNSIGNED;comment:商户id;default:0;" json:"mer_id"`                                // 商户id
	ProductId     int32     `gorm:"column:product_id;type:int;comment:商品关联id;default:0;" json:"product_id"`                               // 商品关联id
	IsDel         int64     `gorm:"column:is_del;type:bigint;comment:1未删除2已删除;default:1;" json:"is_del"`                                  // 1未删除2已删除
	UpdatedAt     time.Time `gorm:"column:updated_at;type:datetime(3);not null;default:CURRENT_TIMESTAMP(3);" json:"updated_at"`
}

func (ab *Article) CreateEbArticle() bool {
	err := global.DB.Debug().Table("article").Create(&ab).Error
	if err != nil {
		return false
	}
	return true
}
