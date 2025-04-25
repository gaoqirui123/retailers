package handler

import (
	"common/global"
	"common/model"
	"common/mongoDB"
	"common/proto/article"
	"context"
	"fmt"
	"time"
)

const KeyArticleDID = "key_article_DID"
const KeyArticleCategoryDID = "key_ArticleCategory_DID"

// 文章管理添加
func ArticleAdd(in *article.ArticleAddRequest) (*article.ArticleAddResponse, error) {

	u := model.User{}
	id, err := u.FindId(int(in.Uid))
	if err != nil {
		return &article.ArticleAddResponse{Success: fmt.Sprintf("用户登录故障%v", err)}, nil
	}
	if id.Uid == 0 {
		return &article.ArticleAddResponse{Success: "用户未登录"}, nil
	}

	err = global.Rdb.Incr(context.Background(), KeyArticleDID).Err()

	if err != nil {
		return &article.ArticleAddResponse{Success: "自增失败"}, nil
	}

	ID, err := global.Rdb.Get(context.Background(), KeyArticleDID).Int()

	if err != nil {
		return &article.ArticleAddResponse{Success: "获取自增id失败"}, nil
	}

	ab := model.Article{
		IntID:         int64(ID),
		Cid:           int64(in.Cid),
		Title:         in.Title,
		Author:        in.Author,
		ImageInput:    in.ImageInput,
		Synopsis:      in.Synopsis,
		ShareTitle:    " ",
		ShareSynopsis: " ",
		Visit:         1,
		Url:           " ",
		Status:        "关闭",
		AddTime:       time.Now(),
		Hide:          "未隐藏",
		AdminId:       1,
		MerId:         1,
		ProductId:     1,
		IsDel:         1,
		UpdatedAt:     time.Now(),
	}

	c := model.ArticleContent{
		Nid:     ab.IntID,
		Content: in.Content,
		IsDel:   1,
	}

	// 创建 MongoDB 会话
	session, err := global.MDB.StartSession()
	if err != nil {
		return &article.ArticleAddResponse{Success: "启动会话失败"}, nil
	}
	defer session.EndSession(context.Background())

	//查询分类是否存在
	pid, err := mongoDB.FindArticleCategoryPid(global.NaCos.Mongodb.Database, "article_category", int64(in.Cid))

	if err != nil {
		return &article.ArticleAddResponse{Success: fmt.Sprintf("%v", err)}, nil
	}

	if pid.ID.IsZero() {
		return &article.ArticleAddResponse{Success: "此分类不存在"}, nil
	}

	//同步mongodb

	err = mongoDB.CreateArticleContent(global.NaCos.Mongodb.Database, "article", ab)

	if err != nil {

		return &article.ArticleAddResponse{Success: "文章类容添加失败"}, nil
	}

	err = mongoDB.CreateArticleContent(global.NaCos.Mongodb.Database, "article_content", c)

	if err != nil {

		return &article.ArticleAddResponse{Success: "文章管理添加失败"}, nil

	}

	return &article.ArticleAddResponse{Success: "文章添加成功"}, nil

}

// 文章分类添加
func CategoryAdd(in *article.CategoryAddRequest) (*article.CategoryAddResponse, error) {

	//查询用户表
	u := model.User{}
	id, err := u.FindId(int(in.Uid))

	if err != nil {
		return &article.CategoryAddResponse{Success: "用户登录故障，请从新登录"}, nil
	}
	if id.Uid == 0 {
		return &article.CategoryAddResponse{Success: "用户未登录"}, nil
	}

	err = global.Rdb.Incr(context.Background(), KeyArticleCategoryDID).Err()

	if err != nil {
		return &article.CategoryAddResponse{Success: "自增失败"}, nil
	}

	ID, err := global.Rdb.Get(context.Background(), KeyArticleCategoryDID).Int()

	if err != nil {
		return &article.CategoryAddResponse{Success: "获取自增id失败"}, nil
	}

	a := model.ArticleCategory{
		IntID:  int64(ID),
		Pid:    int64(in.Pid),
		Title:  in.Title,
		Intr:   in.Intr,
		Image:  in.ImageInput,
		Status: int64(uint8(in.Status)),
		Sort:   int64(in.Sort),
		IsDel:  1,
	}

	err = mongoDB.CreateArticleContent(global.NaCos.Mongodb.Database, "article_category", a)

	if err != nil {

		return &article.CategoryAddResponse{Success: "分类添加失败"}, nil

	}
	return &article.CategoryAddResponse{Success: "分类添加成功"}, nil
}

// 查询文章管理列表
func ArticleList(in *article.ArticleListRequest) (*article.ArticleListResponse, error) {

	category, err := mongoDB.FindArticleCategory(global.NaCos.Mongodb.Database, "article")
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
	pid, err := mongoDB.FindArticleCategoryPid(global.NaCos.Mongodb.Database, "article_category", int64(in.Cid))
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	if pid.ID.IsZero() {
		return nil, fmt.Errorf("此分类不存在")
	}

	cid, err := mongoDB.FindArticleCid(global.NaCos.Mongodb.Database, "article", int64(in.Cid))
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
			Id:         uint32(e.ID),
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
			Id:         uint32(e.ID),
		})
	}
	if err != nil {
		return nil, err
	}

	return &article.ArticleSearchResponse{List: list}, nil
}

// 编辑文章
func EditArticle(in *article.EditArticleRequest) (*article.EditArticleResponse, error) {

	find, err := mongoDB.FindArticle(global.NaCos.Mongodb.Database, "article", int64(in.Id))
	if err != nil {
		return &article.EditArticleResponse{Success: fmt.Sprintf("%v", err)}, nil
	}

	if find.IntID == 0 {
		return &article.EditArticleResponse{Success: "文章不存在"}, nil
	}

	//查询分类是否存在
	pid, err := mongoDB.FindArticleCategoryPid(global.NaCos.Mongodb.Database, "article_category", int64(int(in.Cid)))
	if err != nil {
		return &article.EditArticleResponse{Success: fmt.Sprintf("%v", err)}, nil
	}

	if pid.ID.IsZero() {
		return &article.EditArticleResponse{Success: "此分类不存在"}, nil
	}

	date := model.Article{
		Title:      in.Title,
		Author:     in.Author,
		ImageInput: in.ImageInput,
		Synopsis:   in.Synopsis,
		Hide:       in.Hide,
		Cid:        int64(in.Cid),
		UpdatedAt:  time.Now(),
	}

	err = mongoDB.EditArticle(global.NaCos.Mongodb.Database, "article", int64(in.Id), date)

	if err != nil {
		return &article.EditArticleResponse{Success: fmt.Sprintf("%v", err)}, nil
	}

	return &article.EditArticleResponse{Success: "文章编辑成功"}, nil

}

// 删除文章管理
func DeleteArticle(in *article.DeleteRequest) (*article.DeleteResponse, error) {

	//判断文章是否存在
	find, _ := mongoDB.FindArticle(global.NaCos.Mongodb.Database, "article", int64(in.Ids))

	if find.ID == 0 {
		return &article.DeleteResponse{Success: "文章不存在"}, nil
	}

	//删除文章管理表
	err := mongoDB.DeleteArticle(global.NaCos.Mongodb.Database, "article", int64(in.Ids))
	if err != nil {
		return &article.DeleteResponse{Success: "文章管理删除失败"}, nil
	}

	//同步删除文章类容表
	err = mongoDB.DeleteArticleContent(global.NaCos.Mongodb.Database, "article_content", int64(in.Ids))

	if err != nil {

		return &article.DeleteResponse{Success: "文章管理删除失败"}, nil

	}

	return &article.DeleteResponse{Success: "文章管理删除成功"}, nil
}

// 删除文章分类
func DeleteArticleCategory(in *article.DeleteRequest) (*article.DeleteResponse, error) {

	err := mongoDB.DeleteArticle(global.NaCos.Mongodb.Database, "article_category", int64(in.Ids))

	if err != nil {
		return &article.DeleteResponse{Success: "文章管理删除失败"}, nil
	}

	return &article.DeleteResponse{Success: "文章管理删除成功"}, nil

}
