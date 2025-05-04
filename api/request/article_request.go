package request

// 文章管理添加
type ArticleAdd struct {
	Content string `json:"content" form:"content" binding:"required"`
	//	Uid        uint32 `json:"uid" form:"uid" binding:"required"`
	Cid        uint32 `json:"cid" form:"cid" binding:"required"`
	Title      string `json:"title" form:"title" binding:"required"`
	Author     string `json:"author" form:"author" binding:"required"`
	ImageInput string `json:"imageInput" form:"imageInput" binding:"required"`
	Synopsis   string `json:"synopsis" form:"synopsis" binding:"required"`
	Hide       string `json:"hide" form:"hide" binding:"required"`
}

// 文章分类添加
type CategoryAdd struct {
	//Cid        uint32 `json:"cid" form:"cid" binding:"required"`
	Pid        uint32 `json:"pid" form:"pid" binding:"required"`
	Title      string `json:"title" form:"title" binding:"required"`
	Intr       string `json:"intr" form:"intr" binding:"required"`
	ImageInput string `json:"image_input" form:"image_input" binding:"required"`
	Sort       uint32 `json:"sort" form:"sort" binding:"required"`
	Status     uint32 `json:"status" form:"status" binding:"required"`
}

// 查询文章管理分类列表
type CategoryList struct {
	Cid uint32 `json:"cid" form:"cid"` //分类id为0查询所有，可以为空
}

// 文章标题搜索
type ArticleSearch struct {
	Title string `json:"title" form:"title"` //文章标题搜索，可以为空
}

// 编辑文章
type EditArticle struct {
	Id         uint32 `json:"id" form:"id" binding:"required"`
	Cid        uint32 `json:"cid" form:"cid" binding:"required"`
	Title      string `json:"title" form:"title" binding:"required"`
	Author     string `json:"author" form:"author" binding:"required"`
	ImageInput string `json:"imageInput" form:"imageInput" binding:"required"`
	Synopsis   string `json:"synopsis" form:"synopsis" binding:"required"`
	Hide       string `json:"hide" form:"hide" binding:"required"`
}

// 删除文章
type DeleteArticle struct {
	Id uint32 `json:"id" form:"id" binding:"required"`
}

// 发布评论
type PostAComment struct {
	Pid         uint32 `json:"pid" form:"pid"`                                  //父级评论 ID
	ArticleID   uint32 `json:"article_id" form:"article_id" binding:"required"` //文章 ID
	Content     string `json:"content" form:"content" binding:"required"`       //评论内容
	ReplyUserID uint32 `json:"reply_user_id" form:"reply_user_id"`              //被回复的用户 ID

}

// 文章点赞
type ArticleThumbsUp struct {
	ArticleID uint32 `json:"article_id" form:"article_id" binding:"required"` //文章 ID
	Button    uint32 `json:"button" form:"button" binding:"required"`         //1点赞2取消

}

// 删除评论
type DeleteComment struct {
	ArticleID uint32 `json:"article_id" form:"article_id" binding:"required"` //文章 ID
	CommentID uint32 `json:"comment_id" form:"comment_id" binding:"required"` //文章 ID
}

// 高赞文章排序
type TopLikeArticleRanking struct {
	Top uint32 `json:"top" form:"top" binding:"required"` //文章 ID

}
