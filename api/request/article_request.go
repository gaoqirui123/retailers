package request

// 文章管理添加
type ArticleRelease struct {
	Content    string `json:"content" form:"content" binding:"required"`
	Cid        uint32 `json:"cid" form:"cid" binding:"required"`
	Title      string `json:"title" form:"title" binding:"required"`
	Author     string `json:"author" form:"author" binding:"required"`
	ImageInput string `json:"imageInput" form:"imageInput" binding:"required"`
	Synopsis   string `json:"synopsis" form:"synopsis" binding:"required"`
	Hide       string `json:"hide" form:"hide" binding:"required"`
}

// 文章分类添加
type CategoryAdd struct {
	Pid        uint32 `json:"pid" form:"pid" binding:"required"`
	Title      string `json:"title" form:"title" binding:"required"`
	Intr       string `json:"intr" form:"intr" binding:"required"`
	ImageInput string `json:"image_input" form:"image_input" binding:"required"`
	Sort       uint32 `json:"sort" form:"sort" binding:"required"`
	Status     uint32 `json:"status" form:"status" binding:"required"`
}
