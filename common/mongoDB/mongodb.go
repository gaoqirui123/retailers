package mongoDB

import (
	"common/global"
	"common/model"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"time"
)

// 添加
func CreateArticleContent(dateBase, collectionName string, doc interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := global.MDB.Database(dateBase).Collection(collectionName)
	_, err := collection.InsertOne(ctx, doc)
	return err
}

// 查询文章管理列表
func FindArticleCategory(dateBase, collectionName string) ([]model.Article, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	collection := global.MDB.Database(dateBase).Collection(collectionName)
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("%w", err)

	}
	// 文章分类表
	defer cur.Close(ctx)

	var res []model.Article

	for cur.Next(ctx) {

		var result model.Article

		err = cur.Decode(&result)

		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		res = append(res, result)
	}

	err = cur.Err()

	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return res, nil
}

// 查询文章分类表的类型id
func FindArticleCategoryPid(dateBase, collectionName string, pid int) (model.ArticleCategory, error) {
	var date model.ArticleCategory

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()

	collection := global.MDB.Database(dateBase).Collection(collectionName)

	fist := bson.D{{"pid", pid}}

	err := collection.FindOne(ctx, fist).Decode(&date)

	if err != nil {
		return model.ArticleCategory{}, nil
	}
	if errors.Is(err, mongo.ErrNoDocuments) {
		return model.ArticleCategory{}, nil
	} else if err != nil {
		return model.ArticleCategory{}, nil
	}

	return date, err
}

// 查询文章管理分类列表
func FindArticleCid(dateBase, collectionName string, cid int) ([]model.Article, error) {
	var articles []model.Article

	// 设置上下文超时
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// 获取集合
	collection := global.MDB.Database(dateBase).Collection(collectionName)

	// 构造查询条件
	var filter bson.D
	if cid == 0 {
		// 如果 cid 为 0，查询全部数据
		filter = bson.D{}
	} else {
		// 如果 cid 不为 0，按 cid 查询
		filter = bson.D{{"cid", cid}}
	}

	// 执行查询
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to query articles: %w", err)
	}
	defer cursor.Close(ctx)

	// 遍历游标并解码文档
	for cursor.Next(ctx) {
		var article model.Article
		if err := cursor.Decode(&article); err != nil {
			return nil, fmt.Errorf("failed to decode article: %w", err)
		}
		articles = append(articles, article)
	}

	// 检查游标操作是否有错误
	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return articles, nil
}

// 文章标题搜索
func FindArticleTitle(dateBase, collectionName string, title string) ([]model.Article, error) {
	var articles []model.Article

	// 设置上下文超时
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// 获取集合
	collection := global.MDB.Database(dateBase).Collection(collectionName)
	var filter bson.D
	if title == "" {
		// 如果 title 为 空，查询全部数据
		filter = bson.D{}
	} else {
		// 如果 title 不为 0，按 title 查询
		filter = bson.D{{"title", title}}
	}

	// 执行查询
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to query articles: %w", err)
	}
	defer cursor.Close(ctx)

	// 遍历游标并解码文档
	for cursor.Next(ctx) {
		var article model.Article
		if err = cursor.Decode(&article); err != nil {
			return nil, fmt.Errorf("failed to decode article: %w", err)
		}
		articles = append(articles, article)
	}

	// 检查游标操作是否有错误
	if err = cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return articles, nil
}

// 编辑文章
func EditArticle(dateBase, collectionName string, id int, date model.Article) error {
	// 设置上下文超时
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()

	// 获取集合
	collection := global.MDB.Database(dateBase).Collection(collectionName)

	filter := bson.D{{"id", id}}
	update := bson.D{{"$set", date}}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil
	}
	return nil

}

// 查询文章管理的id
func FindArticle(dateBase, collectionName string, id int) (model.Article, error) {
	var date model.Article

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()

	collection := global.MDB.Database(dateBase).Collection(collectionName)

	fist := bson.D{{"id", id}}

	err := collection.FindOne(ctx, fist).Decode(&date)

	if err != nil {
		return model.Article{}, nil
	}
	if errors.Is(err, mongo.ErrNoDocuments) {
		return model.Article{}, nil
	} else if err != nil {
		return model.Article{}, nil
	}

	return date, err
}

// 删除文章管理
func DeleteArticle(dateBase, collectionName string, id int) error {
	// 设置上下文超时
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// 获取集合
	coll := global.MDB.Database(dateBase).Collection(collectionName)

	// 构造查询条件
	filter := bson.D{{"id", id}}

	// 执行删除操作
	result, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete articles: %w", err)
	}

	// 检查是否有文档被删除
	if result.DeletedCount == 0 {
		return fmt.Errorf("no documents found with cid %d", id)
	}

	return nil
}

// 删除文章类容
func DeleteArticleContent(dateBase, collectionName string, id int) error {
	// 设置上下文超时
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// 获取集合
	coll := global.MDB.Database(dateBase).Collection(collectionName)

	// 构造查询条件
	filter := bson.D{{"nid", id}}

	// 执行删除操作
	result, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete articles: %w", err)
	}

	// 检查是否有文档被删除
	if result.DeletedCount == 0 {
		return fmt.Errorf("no documents found with cid %d", id)
	}

	return nil
}
