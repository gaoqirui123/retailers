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
func CreateEbArticleContent(dateBase, collectionName string, doc interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := global.MDB.Database(dateBase).Collection(collectionName)
	_, err := collection.InsertOne(ctx, doc)
	return err
}

// 查询文章管理列表
func FindEbArticleCategory(dateBase, collectionName string) ([]model.EbArticle, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	collection := global.MDB.Database(dateBase).Collection(collectionName)
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("%w", err)

	}
	// 文章分类表
	defer cur.Close(ctx)

	var res []model.EbArticle

	for cur.Next(ctx) {

		var result model.EbArticle

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
func FindEbArticleCategoryPid(dateBase, collectionName string, pid int) (model.EbArticleCategory, error) {
	var date model.EbArticleCategory

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()

	collection := global.MDB.Database(dateBase).Collection(collectionName)

	fist := bson.D{{"pid", pid}}

	err := collection.FindOne(ctx, fist).Decode(&date)

	if err != nil {
		return model.EbArticleCategory{}, nil
	}
	if errors.Is(err, mongo.ErrNoDocuments) {
		return model.EbArticleCategory{}, nil
	} else if err != nil {
		return model.EbArticleCategory{}, nil
	}

	return date, err
}
