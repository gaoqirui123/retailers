package handler

import (
	"common/global"
	"common/model"
	article2 "common/model/article"
	"common/mongoDB"
	"common/proto/article"
	"context"
	"fmt"
	"time"
)

const KeyArticleDID = "key_article_DID"                 //文章id
const KeyArticleCategoryDID = "key_ArticleCategory_DID" //文章分类id
const KeyPostACommentDID = "key_post_a_comment_DID"     //评论id

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

	ab := article2.Article{
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
		LikeCount:     1,
	}

	c := article2.ArticleContent{
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

	a := article2.ArticleCategory{
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

	date := article2.Article{
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

	err := mongoDB.DeleteArticle(global.NaCos.Mongodb.Database, "article_category", in.Ids)

	if err != nil {
		return &article.DeleteResponse{Success: "文章管理删除失败"}, nil
	}

	return &article.DeleteResponse{Success: "文章管理删除成功"}, nil

}

//发布评论

func PostAComment(in *article.PostACommentRequest) (*article.PostACommentResponse, error) {

	//查找文章id
	find, err := mongoDB.FindArticle(global.NaCos.Mongodb.Database, "article", in.ArticleID)
	if err != nil {
		return &article.PostACommentResponse{Success: "查找失败"}, nil
	}

	if find.IntID == 0 {
		return &article.PostACommentResponse{Success: "文章不存在"}, nil
	}

	//限制评论字数100字以内-失败 小于0-失败

	if len(in.Content) > 100 && len(in.Content) < 0 {
		return &article.PostACommentResponse{Success: "评论字数过长或过短"}, nil
	}

	err = global.Rdb.Incr(context.Background(), KeyArticleDID).Err()

	if err != nil {
		return &article.PostACommentResponse{Success: "id存入失败"}, nil
	}

	ID, err := global.Rdb.Get(context.Background(), KeyArticleDID).Int()

	if err != nil {
		return &article.PostACommentResponse{Success: "id获取失败"}, nil
	}

	uComment := article2.ArticleUserComment{
		CommentID:   int64(ID),
		UserID:      in.Uid,
		ArticleID:   in.ArticleID,
		Content:     in.Content,
		Pid:         in.Pid,
		ReplyUserID: in.ReplyUserID,
		Status:      1,
		IsDel:       1,
		AddTime:     time.Time{},
		UpdateTime:  time.Time{},
	}
	//同步mongodb

	err = mongoDB.CreateArticleContent(global.NaCos.Mongodb.Database, "article_user_comment", uComment)

	if err != nil {
		return &article.PostACommentResponse{Success: "评论失败"}, nil
	}
	return &article.PostACommentResponse{Success: "评论成功"}, nil

}

// 文章点赞
func ArticleThumbsUp(in *article.ArticleThumbsUpRequest) (*article.ArticleThumbsUpResponse, error) {

	// 创建 MongoDB 会话
	session, err := global.MDB.StartSession()
	if err != nil {
		return &article.ArticleThumbsUpResponse{Success: "启动会话失败"}, nil
	}
	defer session.EndSession(context.Background())

	// 开始事务
	err = session.StartTransaction()
	if err != nil {
		return &article.ArticleThumbsUpResponse{Success: "启动事务失败"}, nil
	}

	// 定义提交或回滚事务的函数
	commitOrAbort := func(err error) error {
		if err != nil {
			if abortErr := session.AbortTransaction(context.Background()); abortErr != nil {
				return fmt.Errorf("事务回滚失败: %v; 原始错误: %v", abortErr, err)
			}
			return err
		}
		return session.CommitTransaction(context.Background())
	}

	lc := article2.ArticleUserCommentLikeCount{
		UserID:      in.Uid,
		CommentID:   in.ArticleID,
		CreatedTime: time.Now(),
		IsDel:       1,
		UpdateTime:  time.Now(),
	}

	find, err := mongoDB.FindArticle(global.NaCos.Mongodb.Database, "article", in.ArticleID)
	if err != nil {
		return &article.ArticleThumbsUpResponse{Success: fmt.Sprintf("%v", err)}, nil
	}

	if find.IntID == 0 {
		return &article.ArticleThumbsUpResponse{Success: "文章不存在"}, nil
	}

	if in.Button == 1 {
		//1用户点赞，判断用户有没有对商品进行点赞，未点赞可以点赞
		//查找用户点赞记录
		like, _ := mongoDB.FindArticleUserCommentLikeCount(global.NaCos.Mongodb.Database, "article_user_comment_like_count", in.ArticleID)
		if like.UserID == 0 {
			//点赞记录不存在则记录
			//同步mongodb
			err = mongoDB.CreateArticleContent(global.NaCos.Mongodb.Database, "article_user_comment_like_count", lc)

			if err != nil {
				if err = commitOrAbort(err); err != nil {
					return &article.ArticleThumbsUpResponse{Success: fmt.Sprintf("文章点赞失败: %v", err)}, nil
				}
			}
		}
		//总赞数+1
		count := find.LikeCount + 1
		err = mongoDB.ArticleUserCommentLikeCountIncr(global.NaCos.Mongodb.Database, "article", in.ArticleID, count)

		if err != nil {
			if err = commitOrAbort(err); err != nil {
				return &article.ArticleThumbsUpResponse{Success: fmt.Sprintf("%v", err)}, nil
			}
		}

	} else if in.Button == 2 {
		//查找用户点赞记录

		like, _ := mongoDB.FindArticleUserCommentLikeCount(global.NaCos.Mongodb.Database, "article_user_comment_like_count", in.ArticleID)
		if like.UserID == 0 {
			return &article.ArticleThumbsUpResponse{Success: "文章未点赞取消不了"}, nil
		}

		//总赞数-1
		count := find.LikeCount - 1

		err = mongoDB.ArticleUserCommentLikeCountIncr(global.NaCos.Mongodb.Database, "article", in.ArticleID, count)

		if err != nil {
			if err = commitOrAbort(err); err != nil {
				return &article.ArticleThumbsUpResponse{Success: fmt.Sprintf("%v", err)}, nil
			}
		}
		err = mongoDB.CancelLikes(global.NaCos.Mongodb.Database, "article_user_comment_like_count", in.ArticleID, in.Uid)
		if err != nil {
			if err = commitOrAbort(err); err != nil {
				return &article.ArticleThumbsUpResponse{Success: fmt.Sprintf("%v", err)}, nil
			}
		}

	}

	if err = commitOrAbort(nil); err != nil {
		return &article.ArticleThumbsUpResponse{Success: fmt.Sprintf("事务提交失败: %v", err)}, nil
	}

	return &article.ArticleThumbsUpResponse{Success: "成功"}, nil

}

// 删除评论
func DeleteComment(in *article.DeleteCommentRequest) (*article.DeleteCommentResponse, error) {
	find, err := mongoDB.FindArticle(global.NaCos.Mongodb.Database, "article", in.ArticleID)
	if err != nil {
		return &article.DeleteCommentResponse{Success: fmt.Sprintf("%v", err)}, nil
	}

	if find.IntID == 0 {
		return &article.DeleteCommentResponse{Success: "文章不存在"}, nil
	}
	err = mongoDB.DeleteComment(global.NaCos.Mongodb.Database, "article_user_comment", in.ArticleID, in.Uid, in.CommentID)

	if err != nil {
		return &article.DeleteCommentResponse{Success: "评论删除失败"}, nil
	}

	return &article.DeleteCommentResponse{Success: "评论删除成功"}, nil
}

func TopLikeArticleRanking(in *article.TopLikeArticleRankingRequest) (*article.TopLikeArticleRankingResponse, error) {
	if in.Top == 2 {
		in.Top = -1
	}
	articles, err := mongoDB.FindTopLikeArticles(global.NaCos.Mongodb.Database, "article", in.Top)
	if err != nil {
		return nil, err
	}

	var list []*article.Ranking
	for _, e := range articles {
		list = append(list, &article.Ranking{
			Cid:        uint32(e.Cid),
			Title:      e.Title,
			Author:     e.Author,
			ImageInput: e.ImageInput,
			Synopsis:   e.Synopsis,
			Hide:       e.Hide,
			Id:         uint32(e.ID),
			LikeCount:  uint32(e.LikeCount),
		})
	}

	return &article.TopLikeArticleRankingResponse{List: list}, nil
}
