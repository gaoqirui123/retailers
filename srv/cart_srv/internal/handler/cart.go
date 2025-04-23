package handler

import (
	"common/model"
	"common/proto/cart"
	"errors"
	"fmt"
)

func AddCart(in *cart.AddCartRequest) (*cart.AddCartResponse, error) {
	// 判断商品是否存在和下架
	p := &model.Product{}
	err := p.GetProductIdBy(in.ProductId)
	if err != nil {
		return nil, err
	}
	if p.Id == 0 {
		return nil, errors.New("未查询到商品信息")
	}
	// 商品是否下架
	if p.IsShow == 0 {
		return nil, errors.New("该商品已经下架")
	}
	// 判断商品库存是否充足
	if p.Stock < int(in.CartNum) {
		return nil, errors.New("商品库存不足")
	}

	// 判断购物车是否存在该商品，如果存在，则数量累加一，如果不在，则添加购物车
	c := &model.Cart{}
	err = c.GetStoreCart(in.Uid, in.ProductId)
	if err != nil {
		return nil, err
	}
	if c.Id == 0 {
		c = &model.Cart{
			Uid:               uint32(in.Uid),
			Type:              in.Type,
			ProductId:         uint32(in.ProductId),
			ProductAttrUnique: in.ProductAttrUnique,
			CartNum:           uint16(in.CartNum),
			IsPay:             int8(in.IsPay),
			IsNew:             int8(in.IsNew),
			CombinationId:     uint32(in.CombinationId),
			SeckillId:         uint32(in.SeckillId),
			BargainId:         uint32(in.BargainId),
		}
		err = c.AddCart()
		if err != nil {
			return nil, err
		}
		if c.Id == 0 {
			return nil, errors.New("添加购物车失败")
		}
	} else {
		err = c.UpdateCartNum(in.Uid, in.ProductId, in.CartNum)
		if err != nil {
			return nil, errors.New("购物车数量累加失败")
		}
	}
	return &cart.AddCartResponse{CartId: int64(c.Id)}, nil
}

func ClearCart(in *cart.ClearCartRequest) (*cart.ClearCartResponse, error) {
	carts := &model.Cart{}
	err := carts.ClearCart(in.Uid)
	if err != nil {
		return nil, errors.New("清空购物车失败")
	}
	return &cart.ClearCartResponse{Success: true}, nil
}

func DeleteCart(in *cart.DeleteCartRequest) (*cart.DeleteCartResponse, error) {
	carts := &model.Cart{}
	err := carts.DeleteCart(in.Uid, in.ProductId)
	if err != nil {
		return nil, errors.New("删除购物车商品失败")
	}
	return &cart.DeleteCartResponse{Success: true}, nil
}

func GetCartList(in *cart.GetCartListRequest) (*cart.GetCartListResponse, error) {
	carts := &model.Cart{}
	list, err := carts.GetCartList(in.Uid)
	if err != nil {
		return nil, errors.New("购物车列表展示失败")
	}
	if list == nil {
		return nil, errors.New("购物车列表展示失败")
	}
	var cartList []*cart.GetCartList
	for _, c := range list {
		cartList = append(cartList, &cart.GetCartList{
			Uid:               int64(c.Uid),
			Type:              c.Type,
			ProductId:         int64(c.ProductId),
			ProductAttrUnique: c.ProductAttrUnique,
			CartNum:           int64(c.CartNum),
			IsPay:             int64(c.IsPay),
			IsNew:             int64(c.IsNew),
			CombinationId:     int64(c.CombinationId),
			SeckillId:         int64(c.SeckillId),
			BargainId:         int64(c.BargainId),
			CartId:            int64(c.Id),
			AddTime:           int64(c.AddTime),
		})
	}
	fmt.Println(list, "222222")
	fmt.Println(cartList, "3333333333")
	return &cart.GetCartListResponse{List: cartList}, nil
}
