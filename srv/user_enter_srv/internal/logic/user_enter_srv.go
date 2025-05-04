package logic

import (
	"common/proto/user_enter"
	"context"
	"user_enter_srv/internal/handler"
)

type UserEnterServer struct {
	user_enter.UnimplementedUserEnterServer
}

// Apply TODO:商户申请注册
func (es UserEnterServer) Apply(ctx context.Context, in *user_enter.UserEnterApplyRequest) (*user_enter.UserEnterApplyResponse, error) {
	register, err := handler.Apply(in)
	if err != nil {
		return nil, err
	}
	return register, nil
}

// AddProduct TODO:商户发布商品
func (es UserEnterServer) AddProduct(ctx context.Context, in *user_enter.AddProductRequest) (*user_enter.AddProductResponse, error) {
	register, err := handler.AddProduct(in)
	if err != nil {
		return nil, err
	}
	return register, nil
}

// AddCombinationProduct TODO:添加拼团商品
func (es UserEnterServer) AddCombinationProduct(ctx context.Context, in *user_enter.AddCombinationProductRequest) (*user_enter.AddCombinationProductResponse, error) {
	product, err := handler.AddCombinationProduct(in)
	if err != nil {
		return nil, err
	}
	return product, nil
}

// ProcessInvoice TODO:商家审核用户的发票invoice申请
func (es UserEnterServer) ProcessInvoice(ctx context.Context, in *user_enter.ProcessInvoiceRequest) (*user_enter.ProcessInvoiceResponse, error) {
	product, err := handler.ProcessInvoice(in)
	if err != nil {
		return nil, err
	}
	return product, nil
}

// InvoiceList TODO:发票列表展示
func (es UserEnterServer) InvoiceList(ctx context.Context, in *user_enter.InvoiceListRequest) (*user_enter.InvoiceListResponse, error) {
	product, err := handler.InvoiceList(in)
	if err != nil {
		return nil, err
	}
	return product, nil
}

// DelProduct TODO:商户下架商品
func (es UserEnterServer) DelProduct(ctx context.Context, in *user_enter.DelProductRequest) (*user_enter.DelProductResponse, error) {
	register, err := handler.DelProduct(in)
	if err != nil {
		return nil, err
	}
	return register, nil
}

// Register TODO:商户申请注册
func (es UserEnterServer) Register(ctx context.Context, in *user_enter.UserEnterRegisterRequest) (*user_enter.UserEnterRegisterResponse, error) {
	register, err := handler.Register(in)
	if err != nil {
		return nil, err
	}
	return register, nil
}

// Login TODO:商户登录
func (es UserEnterServer) Login(ctx context.Context, in *user_enter.UserEnterLoginRequest) (*user_enter.UserEnterLoginResponse, error) {
	register, err := handler.Login(in)
	if err != nil {
		return nil, err
	}
	return register, nil
}

// AddSeckillProduct  TODO: 添加秒杀商品
func (es UserEnterServer) AddSeckillProduct(ctx context.Context, in *user_enter.AddSeckillProductRequest) (*user_enter.AddSeckillProductResponse, error) {
	add, err := handler.AddSeckillProduct(in)
	if err != nil {
		return nil, err
	}
	return add, nil
}

// ReverseStock  TODO: 秒杀后反还剩余的商品
func (es UserEnterServer) ReverseStock(ctx context.Context, in *user_enter.ReverseStockRequest) (*user_enter.ReverseStockResponse, error) {
	reverse, err := handler.ReverseStock(in)
	if err != nil {
		return nil, err
	}
	return reverse, nil
}
