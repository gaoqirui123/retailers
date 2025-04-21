package handler

import (
	"common/model"
	"common/proto/product"
)

// CombinationList TODO:拼团商品列表展示
func CombinationList(in *product.CombinationListRequest) (*product.CombinationListResponse, error) {
	c := model.Combination{}
	list, err := c.GetCombinationList()
	if err != nil {
		return nil, err
	}
	var lists []*product.CombinationList
	for _, combination := range list {
		l := product.CombinationList{
			Image:  combination.Images,
			Title:  combination.Title,
			People: int64(combination.People),
			Price:  float32(combination.Price),
			Stock:  int64(combination.Stock),
		}
		lists = append(lists, &l)
	}
	return &product.CombinationListResponse{List: lists}, nil
}
