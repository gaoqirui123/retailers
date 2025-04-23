package handler

import (
	"common/global"
	"common/model"
	"common/mongoDB"
	"common/proto/article"
	"fmt"
	"log"
)

// 文章管理添加
func ArticleAdd(in *article.ArticleAddRequest) (*article.ArticleAddResponse, error) {
	ab := model.Article{
		Cid:        int32(in.Cid),
		Title:      in.Title,
		Author:     in.Author,
		ImageInput: in.ImageInput,
		Synopsis:   in.Synopsis,
		Hide:       in.Hide,
	}

	c := model.ArticleContent{
		Content: in.Content,
	}

	//查询分类是否存在
	pid, err := mongoDB.FindArticleCategoryPid(global.NaCos.Mongodb.Database, "article_category", int(in.Cid))
	if err != nil {
		return &article.ArticleAddResponse{Success: "分类查询失败"}, nil
	}

	if pid.Id == 0 {
		return &article.ArticleAddResponse{Success: "此分类不存在"}, nil
	}

	//文章类容添加

	if !c.CreateEbArticleContent() {
		return &article.ArticleAddResponse{Success: "文章类容添加失败"}, nil
	}

	//文章管理添加
	if !ab.CreateEbArticle() {
		return &article.ArticleAddResponse{Success: "文章管理添加失败"}, nil
	}

	//同步mongodb

	err = mongoDB.CreateArticleContent(global.NaCos.Mongodb.Database, "article", ab)

	if err != nil {
		log.Println(err)
		return &article.ArticleAddResponse{Success: "文章类容添加失败"}, nil
	}

	err = mongoDB.CreateArticleContent(global.NaCos.Mongodb.Database, "article_content", c)
	if err != nil {
		log.Println(err)
		return &article.ArticleAddResponse{Success: "文章管理添加失败"}, nil

	}

	return &article.ArticleAddResponse{Success: "文章添加成功"}, nil
}

// 文章分类添加
func CategoryAdd(in *article.CategoryAddRequest) (*article.CategoryAddResponse, error) {
	a := model.ArticleCategory{
		Pid:    int32(in.Pid),
		Title:  in.Title,
		Intr:   in.Intr,
		Image:  in.ImageInput,
		Status: uint8(in.Status),
		Sort:   in.Sort,
	}
	if !a.CreateArticleCategory() {
		return &article.CategoryAddResponse{Success: "分类添加失败"}, nil
	}

	err := mongoDB.CreateArticleContent(global.NaCos.Mongodb.Database, "article_category", a)

	if err != nil {
		log.Println(err)

		return &article.CategoryAddResponse{Success: "分类添加失败"}, nil

	}
	return &article.CategoryAddResponse{Success: "分类添加成功"}, nil
}

// 查询文章管理列表
func ArticleList(in *article.ArticleListRequest) (*article.ArticleListResponse, error) {
	category, err := mongoDB.FindArticleCategory("db", "article")
	var sli []*article.ArticleList
	for _, e := range category {
		sli = append(sli, &article.ArticleList{
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
	return &article.ArticleListResponse{List: sli}, nil
}

// 查询文章管理分类列表
func CategoryList(in *article.CategoryListRequest) (*article.CategoryListResponse, error) {

	//查询分类是否存在
	pid, err := mongoDB.FindArticleCategoryPid(global.NaCos.Mongodb.Database, "article_category", int(in.Cid))
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	if pid.Id == 0 {
		return nil, fmt.Errorf("此分类不存在")
	}

	cid, err := mongoDB.FindArticleCid(global.NaCos.Mongodb.Database, "article", int(in.Cid))
	if err != nil {
		return nil, err
	}

	var list []*article.ArticleList
	for _, e := range cid {
		list = append(list, &article.ArticleList{
			Cid:        uint32(e.Cid),
			Title:      e.Title,
			Author:     e.Author,
			ImageInput: e.ImageInput,
			Synopsis:   e.Synopsis,
			Hide:       e.Hide,
			Id:         e.Id,
		})
	}
	if err != nil {
		return nil, err
	}

	return &article.CategoryListResponse{List: list}, nil
}

// 文章标题搜索
func ArticleSearch(in *article.ArticleSearchRequest) (*article.ArticleSearchResponse, error) {
	cid, err := mongoDB.FindArticleTitle(global.NaCos.Mongodb.Database, "article", in.Title)
	if err != nil {
		return nil, err
	}

	var list []*article.ArticleList
	for _, e := range cid {
		list = append(list, &article.ArticleList{
			Cid:        uint32(e.Cid),
			Title:      e.Title,
			Author:     e.Author,
			ImageInput: e.ImageInput,
			Synopsis:   e.Synopsis,
			Hide:       e.Hide,
			Id:         e.Id,
		})
	}
	if err != nil {
		return nil, err
	}

	return &article.ArticleSearchResponse{List: list}, nil
}

// 编辑文章
func EditArticle(in *article.EditArticleRequest) (*article.EditArticleResponse, error) {

	find, _ := mongoDB.FindArticle(global.NaCos.Mongodb.Database, "article", int(in.Id))

	if find.Id == 0 {
		return &article.EditArticleResponse{Success: "文章不存在"}, nil
	}

	//查询分类是否存在
	pid, err := mongoDB.FindArticleCategoryPid(global.NaCos.Mongodb.Database, "article_category", int(in.Cid))
	if err != nil {
		return &article.EditArticleResponse{Success: "分类查询失败"}, nil
	}

	if pid.Id == 0 {
		return &article.EditArticleResponse{Success: "此分类不存在"}, nil
	}

	date := model.Article{
		Id:         in.Id,
		Title:      in.Title,
		Author:     in.Author,
		ImageInput: in.ImageInput,
		Synopsis:   in.Synopsis,
		Hide:       in.Hide,
		Cid:        int32(in.Cid),
	}

	err = mongoDB.EditArticle(global.NaCos.Mongodb.Database, "article", int(in.Id), date)

	if err != nil {
		return &article.EditArticleResponse{Success: "文章编辑失败"}, nil
	}

	return &article.EditArticleResponse{Success: "文章编辑成功"}, nil

}

func DeleteArticle(in *article.DeleteRequest) (*article.DeleteResponse, error) {

	//判断文章是否存在
	find, _ := mongoDB.FindArticle(global.NaCos.Mongodb.Database, "article", int(in.Ids))

	if find.Id == 0 {
		return &article.DeleteResponse{Success: "文章不存在"}, nil
	}

	//删除文章管理表
	err := mongoDB.DeleteArticle(global.NaCos.Mongodb.Database, "article", int(in.Ids))
	if err != nil {
		return &article.DeleteResponse{Success: "文章管理删除失败"}, nil
	}

	//同步删除文章类容表
	err = mongoDB.DeleteArticleContent(global.NaCos.Mongodb.Database, "article_content", int(in.Ids))

	if err != nil {
		return &article.DeleteResponse{Success: "文章管理删除失败"}, nil
	}

	return &article.DeleteResponse{Success: "文章管理删除成功"}, nil
}

func DeleteArticleCategory(in *article.DeleteRequest) (*article.DeleteResponse, error) {

	err := mongoDB.DeleteArticle(global.NaCos.Mongodb.Database, "article_category", int(in.Ids))

	if err != nil {
		return &article.DeleteResponse{Success: "文章管理删除失败"}, nil
	}

	return &article.DeleteResponse{Success: "文章管理删除成功"}, nil

}
