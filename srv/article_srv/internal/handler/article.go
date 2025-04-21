package handler

import (
	"common/global"
	"common/model"
	"common/mongoDB"
	"common/proto/article"
	"log"
)

// 文章管理添加
func ArticleAdd(in *article.ArticleAddReq) (*article.ArticleAddResp, error) {
	ab := model.EbArticle{

		Cid:        int32(in.Cid),
		Title:      in.Title,
		Author:     in.Author,
		ImageInput: in.ImageInput,
		Synopsis:   in.Synopsis,
		Hide:       in.Hide,
	}

	c := model.EbArticleContent{
		Content: in.Content,
	}

	//查询分类是否存在
	pid, err := mongoDB.FindEbArticleCategoryPid(global.NaCos.Mongodb.Database, "eb_article_category", int(in.Cid))
	if err != nil {
		return &article.ArticleAddResp{Success: "分类查询失败"}, nil
	}

	if pid.Id == 0 {
		return &article.ArticleAddResp{Success: "此分类不存在"}, nil
	}

	//文章类容添加

	if !c.CreateEbArticleContent() {
		return &article.ArticleAddResp{Success: "文章类容添加失败"}, nil
	}

	//文章管理添加
	if !ab.CreateEbArticle() {
		return &article.ArticleAddResp{Success: "文章管理添加失败"}, nil
	}

	//同步mongodb

	err = mongoDB.CreateEbArticleContent(global.NaCos.Mongodb.Database, "eb_article", ab)

	if err != nil {
		log.Println(err)
		return &article.ArticleAddResp{Success: "文章类容添加失败"}, nil
	}

	err = mongoDB.CreateEbArticleContent(global.NaCos.Mongodb.Database, "eb_article_content", c)
	if err != nil {
		log.Println(err)
		return &article.ArticleAddResp{Success: "文章管理添加失败"}, nil

	}

	return &article.ArticleAddResp{Success: "文章添加成功"}, nil
}

// 文章分类添加
func CategoryAdd(in *article.CategoryAddReq) (*article.CategoryAddResp, error) {
	a := model.EbArticleCategory{
		Pid:    int32(in.Pid),
		Title:  in.Title,
		Intr:   in.Intr,
		Image:  in.ImageInput,
		Status: uint8(in.Status),
		Sort:   in.Sort,
	}
	if !a.CreateEbArticleCategory() {
		return &article.CategoryAddResp{Success: "分类添加失败"}, nil
	}

	err := mongoDB.CreateEbArticleContent(global.NaCos.Mongodb.Database, "eb_article_category", a)

	if err != nil {
		log.Println(err)

		return &article.CategoryAddResp{Success: "分类添加失败"}, nil

	}
	return &article.CategoryAddResp{Success: "分类添加成功"}, nil
}

// 查询文章管理列表
func ArticleList(in *article.ArticleListReq) (*article.ArticleListResp, error) {
	category, err := mongoDB.FindEbArticleCategory("db", "eb_article")
	var sli []*article.List
	for _, e := range category {
		sli = append(sli, &article.List{
			Cid:        uint32(e.Cid),
			Title:      e.Title,
			Author:     e.Author,
			ImageInput: e.ImageInput,
			Synopsis:   e.Synopsis,
			Hide:       e.Hide,
		})
	}
	if err != nil {
		return nil, nil
	}
	return &article.ArticleListResp{List: sli}, nil
}
