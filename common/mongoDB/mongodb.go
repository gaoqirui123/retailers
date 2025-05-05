package mongoDB

import (
	"common/global"
	"common/model/article"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// getCollection 封装获取集合和设置上下文超时的逻辑
func getCollection(ctx context.Context, dateBase, collectionName string) (*mongo.Collection, context.Context, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	collection := global.MDB.Database(dateBase).Collection(collectionName)
	return collection, ctx, cancel, nil
}

// CreateArticleContent 向指定集合插入一条文章内容文档
func CreateArticleContent(dateBase, collectionName string, doc interface{}) error {
	ctx := context.Background()
	collection, ctx, cancel, err := getCollection(ctx, dateBase, collectionName)
	if err != nil {
		return err
	}
	defer cancel()

	_, err = collection.InsertOne(ctx, doc)
	return err
}

// FindArticleCategory 查询文章管理列表
func FindArticleCategory(dateBase, collectionName string) ([]article.Article, error) {
	ctx := context.Background()
	collection, ctx, cancel, err := getCollection(ctx, dateBase, collectionName)
	if err != nil {
		return nil, err
	}
	defer cancel()

	cur, err := collection.Find(ctx, bson.D{{"is_del", 1}})
	if err != nil {
		return nil, fmt.Errorf("failed to find articles: %w", err)
	}
	defer cur.Close(ctx)

	var res []article.Article
	for cur.Next(ctx) {
		var result article.Article
		err = cur.Decode(&result)
		if err != nil {
			return nil, fmt.Errorf("failed to decode article: %w", err)
		}
		res = append(res, result)
	}

	if err = cur.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return res, nil
}

// FindArticleCategoryPid 根据 pid 查询文章分类
func FindArticleCategoryPid(dateBase, collectionName string, pid int64) (article.ArticleCategory, error) {
	var category article.ArticleCategory
	ctx := context.Background()
	collection, ctx, cancel, err := getCollection(ctx, dateBase, collectionName)
	if err != nil {
		return category, err
	}
	defer cancel()

	filter := bson.D{
		{"pid", pid},
		{"is_del", 1},
	}

	err = collection.FindOne(ctx, filter).Decode(&category)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return category, fmt.Errorf("no document found with pid %d", pid)
		}
		return category, fmt.Errorf("failed to find article category: %w", err)
	}

	return category, nil
}

// FindArticleCid 根据 cid 查询文章管理分类列表
func FindArticleCid(dateBase, collectionName string, cid int64) ([]article.Article, error) {
	ctx := context.Background()
	collection, ctx, cancel, err := getCollection(ctx, dateBase, collectionName)
	if err != nil {
		return nil, err
	}
	defer cancel()

	var filter bson.D
	if cid == 0 {
		filter = bson.D{{"is_del", 1}}
	} else {
		filter = bson.D{
			{"cid", cid},
			{"is_del", 1},
		}
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to query articles: %w", err)
	}
	defer cursor.Close(ctx)

	var articles []article.Article
	for cursor.Next(ctx) {
		var article article.Article
		if err := cursor.Decode(&article); err != nil {
			return nil, fmt.Errorf("failed to decode article: %w", err)
		}
		articles = append(articles, article)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return articles, nil
}

// FindArticleTitle 根据标题搜索文章
func FindArticleTitle(dateBase, collectionName string, title string) ([]article.Article, error) {
	ctx := context.Background()
	collection, ctx, cancel, err := getCollection(ctx, dateBase, collectionName)
	if err != nil {
		return nil, err
	}
	defer cancel()

	var filter bson.D
	if title == "" {
		filter = bson.D{{"is_del", 1}}
	} else {
		filter = bson.D{{"title", title}, {"is_del", 1}}
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to query articles: %w", err)
	}
	defer cursor.Close(ctx)

	var articles []article.Article
	for cursor.Next(ctx) {
		var article article.Article
		if err = cursor.Decode(&article); err != nil {
			return nil, fmt.Errorf("failed to decode article: %w", err)
		}
		articles = append(articles, article)
	}

	if err = cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return articles, nil
}

// EditArticle 编辑文章
func EditArticle(dateBase, collectionName string, id int64, data article.Article) error {
	ctx := context.Background()
	collection, ctx, cancel, err := getCollection(ctx, dateBase, collectionName)
	if err != nil {
		return err
	}
	defer cancel()

	filter := bson.D{
		{"int_id", id},
		{"is_del", 1},
	}
	update := bson.D{{"$set", data}}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update article: %w", err)
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("article with id %d not found", id)
	}
	return nil
}

// FindArticle 根据 id 查询文章
func FindArticle(dateBase, collectionName string, id int64) (article.Article, error) {
	var category article.Article
	ctx := context.Background()
	collection, ctx, cancel, err := getCollection(ctx, dateBase, collectionName)
	if err != nil {
		return category, err
	}
	defer cancel()

	// 使用传入的 id 参数构建查询条件
	filter := bson.M{
		"int_id": id,
		"is_del": 1,
	}

	err = collection.FindOne(ctx, filter).Decode(&category)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {

			return category, fmt.Errorf("no document found with int_id %d", id)
		}
		return category, fmt.Errorf("failed to find article category: %w", err)
	}

	return category, nil
}

// DeleteArticle 删除文章管理
func DeleteArticle(dateBase, collectionName string, id int64) error {
	ctx := context.Background()
	collection, ctx, cancel, err := getCollection(ctx, dateBase, collectionName)
	if err != nil {
		return err
	}
	defer cancel()

	filter := bson.D{{"int_id", id}}
	update := bson.D{{"$set", bson.D{{"is_del", 2}}}}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update article for deletion: %w", err)
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("article with id %d not found", id)
	}
	return nil
}

// DeleteArticleContent 删除文章内容
func DeleteArticleContent(dateBase, collectionName string, id int64) error {
	ctx := context.Background()
	collection, ctx, cancel, err := getCollection(ctx, dateBase, collectionName)
	if err != nil {
		return err
	}
	defer cancel()

	filter := bson.D{{"nid", id}}
	update := bson.D{{"$set", bson.D{{"is_del", 2}}}}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update article content for deletion: %w", err)
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("article content with nid %d not found", id)
	}
	return nil
}

// 根据文章CommentID查寻文章内容点赞表
func FindArticleUserCommentLikeCount(dateBase, collectionName string, id int64) (article.ArticleUserCommentLikeCount, error) {
	var category article.ArticleUserCommentLikeCount
	ctx := context.Background()
	collection, ctx, cancel, err := getCollection(ctx, dateBase, collectionName)
	if err != nil {
		return category, err
	}
	defer cancel()
	filter := bson.D{
		{"comment_id", id},
		{"is_del", 1},
	}
	err = collection.FindOne(ctx, filter).Decode(&category)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return category, fmt.Errorf("no document found with int_id %d", id)
		}
		return category, fmt.Errorf("failed to find article category: %w", err)
	}

	return category, nil

}

// Article 更新点赞次数
func ArticleUserCommentLikeCountIncr(dateBase, collectionName string, id int64, count int64) error {
	ctx := context.Background()
	collection, ctx, cancel, err := getCollection(ctx, dateBase, collectionName)
	if err != nil {
		return err
	}
	defer cancel()

	filter := bson.D{
		{"int_id", id},
		{"is_del", 1},
	}
	update := bson.D{{"$set", bson.M{"like_count": count}}}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update article: %w", err)
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("article with id %d not found", id)
	}
	return nil
}

// CancelLikes 取消点赞
func CancelLikes(dateBase, collectionName string, id, uId int64) error {
	ctx := context.Background()
	collection, ctx, cancel, err := getCollection(ctx, dateBase, collectionName)
	if err != nil {
		return err
	}
	defer cancel()

	filter := bson.D{{"comment_id", id}, {"user_id", uId}}
	update := bson.D{{"$set", bson.D{{"is_del", 2}}}}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update article for deletion: %w", err)
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("article with id %d not found", id)
	}
	return nil
}

// DeleteArticleContent 删除文章内容
func DeleteComment(dateBase, collectionName string, aId, uId, cID int64) error {
	ctx := context.Background()
	collection, ctx, cancel, err := getCollection(ctx, dateBase, collectionName)
	if err != nil {
		return err
	}
	defer cancel()

	filter := bson.D{{"article_id", aId}, {"user_id", uId}, {"comment_id", cID}}
	update := bson.D{{"$set", bson.D{{"is_del", 2}}}}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update article content for deletion: %w", err)
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("article content with nid %d not found", aId)
	}
	return nil
}

// FindTopLikeArticles 查询点赞数最多的文章
func FindTopLikeArticles(dateBase, collectionName string, top int64) ([]article.Article, error) {
	ctx := context.Background()
	collection, ctx, cancel, err := getCollection(ctx, dateBase, collectionName)
	if err != nil {
		return nil, err
	}
	defer cancel()

	filter := bson.D{{"is_del", 1}}
	sort := bson.D{{"like_count", top}}

	cur, err := collection.Find(ctx, filter, options.Find().SetSort(sort))
	if err != nil {
		return nil, fmt.Errorf("failed to find top like articles: %w", err)
	}
	defer cur.Close(ctx)

	var articles []article.Article
	for cur.Next(ctx) {
		var article article.Article
		if err = cur.Decode(&article); err != nil {
			return nil, fmt.Errorf("failed to decode article: %w", err)
		}
		articles = append(articles, article)
	}

	if err = cur.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return articles, nil
}
