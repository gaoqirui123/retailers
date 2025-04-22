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
