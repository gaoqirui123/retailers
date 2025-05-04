package client

import (
	"common/proto/user_enter"
	"context"
	"google.golang.org/grpc"
)

func NewUserEnterClient(cc grpc.ClientConnInterface) user_enter.UserEnterClient {
	return user_enter.NewUserEnterClient(cc)
}

// UserEnterClients 封装的商户服务客户端处理函数
func UserEnterClients[TRequest, TResponse any](ctx context.Context, request TRequest, operation func(ctx context.Context, client user_enter.UserEnterClient, req TRequest) (TResponse, error)) (TResponse, error) {
	return ExecuteGRPCOperation(ctx, "127.0.0.1:8084", NewUserEnterClient, request, operation)
}

// Apply 商户申请
func Apply(ctx context.Context, in *user_enter.UserEnterApplyRequest) (*user_enter.UserEnterApplyResponse, error) {
	return UserEnterClients(ctx, in, func(ctx context.Context, client user_enter.UserEnterClient, req *user_enter.UserEnterApplyRequest) (*user_enter.UserEnterApplyResponse, error) {
		return client.Apply(ctx, req)
	})
}

// Login 商户登录
func Login(ctx context.Context, in *user_enter.UserEnterLoginRequest) (*user_enter.UserEnterLoginResponse, error) {
	return UserEnterClients(ctx, in, func(ctx context.Context, client user_enter.UserEnterClient, req *user_enter.UserEnterLoginRequest) (*user_enter.UserEnterLoginResponse, error) {
		return client.Login(ctx, req)
	})
}

// Register 商户注册
func Register(ctx context.Context, in *user_enter.UserEnterRegisterRequest) (*user_enter.UserEnterRegisterResponse, error) {
	return UserEnterClients(ctx, in, func(ctx context.Context, client user_enter.UserEnterClient, req *user_enter.UserEnterRegisterRequest) (*user_enter.UserEnterRegisterResponse, error) {
		return client.Register(ctx, req)
	})
}

// AddProduct 添加商品
func AddProduct(ctx context.Context, in *user_enter.AddProductRequest) (*user_enter.AddProductResponse, error) {
	return UserEnterClients(ctx, in, func(ctx context.Context, client user_enter.UserEnterClient, req *user_enter.AddProductRequest) (*user_enter.AddProductResponse, error) {
		return client.AddProduct(ctx, req)
	})
}

// AddCombinationProduct 发布拼团商品
func AddCombinationProduct(ctx context.Context, in *user_enter.AddCombinationProductRequest) (*user_enter.AddCombinationProductResponse, error) {
	return UserEnterClients(ctx, in, func(ctx context.Context, client user_enter.UserEnterClient, req *user_enter.AddCombinationProductRequest) (*user_enter.AddCombinationProductResponse, error) {
		return client.AddCombinationProduct(ctx, req)
	})
}

// DelProduct 下架商品
func DelProduct(ctx context.Context, in *user_enter.DelProductRequest) (*user_enter.DelProductResponse, error) {
	return UserEnterClients(ctx, in, func(ctx context.Context, client user_enter.UserEnterClient, req *user_enter.DelProductRequest) (*user_enter.DelProductResponse, error) {
		return client.DelProduct(ctx, req)
	})
}

// ProcessInvoice 审核发票
func ProcessInvoice(ctx context.Context, in *user_enter.ProcessInvoiceRequest) (*user_enter.ProcessInvoiceResponse, error) {
	return UserEnterClients(ctx, in, func(ctx context.Context, client user_enter.UserEnterClient, req *user_enter.ProcessInvoiceRequest) (*user_enter.ProcessInvoiceResponse, error) {
		return client.ProcessInvoice(ctx, req)
	})
}

// InvoiceList 发票列表展示
func InvoiceList(ctx context.Context, in *user_enter.InvoiceListRequest) (*user_enter.InvoiceListResponse, error) {
	return UserEnterClients(ctx, in, func(ctx context.Context, client user_enter.UserEnterClient, req *user_enter.InvoiceListRequest) (*user_enter.InvoiceListResponse, error) {
		return client.InvoiceList(ctx, req)
	})
}

// AddSeckillProduct 添加秒杀商品
func AddSeckillProduct(ctx context.Context, in *user_enter.AddSeckillProductRequest) (*user_enter.AddSeckillProductResponse, error) {
	return UserEnterClients(ctx, in, func(ctx context.Context, client user_enter.UserEnterClient, req *user_enter.AddSeckillProductRequest) (*user_enter.AddSeckillProductResponse, error) {
		return client.AddSeckillProduct(ctx, req)
	})
}

// ReverseStock 秒杀后反还剩余的商品
func ReverseStock(ctx context.Context, in *user_enter.ReverseStockRequest) (*user_enter.ReverseStockResponse, error) {
	return UserEnterClients(ctx, in, func(ctx context.Context, client user_enter.UserEnterClient, req *user_enter.ReverseStockRequest) (*user_enter.ReverseStockResponse, error) {
		return client.ReverseStock(ctx, req)
	})
}

// BatchReleaseOfProducts 商品批量发布
func BatchReleaseOfProducts(ctx context.Context, in *user_enter.BatchReleaseOfProductsRequest) (*user_enter.BatchReleaseOfProductsResponse, error) {
	return UserEnterClients(ctx, in, func(ctx context.Context, client user_enter.UserEnterClient, req *user_enter.BatchReleaseOfProductsRequest) (*user_enter.BatchReleaseOfProductsResponse, error) {
		return client.BatchReleaseOfProducts(ctx, req)
	})
}

// MerchantVerification 商家核销
func MerchantVerification(ctx context.Context, in *user_enter.MerchantVerificationRequest) (*user_enter.MerchantVerificationResponse, error) {
	return UserEnterClients(ctx, in, func(ctx context.Context, client user_enter.UserEnterClient, req *user_enter.MerchantVerificationRequest) (*user_enter.MerchantVerificationResponse, error) {
		return client.MerchantVerification(ctx, req)
	})
}

// CalculateOrderSummary 商家统计
func CalculateOrderSummary(ctx context.Context, in *user_enter.CalculateOrderSummaryRequest) (*user_enter.CalculateOrderSummaryResponse, error) {
	return UserEnterClients(ctx, in, func(ctx context.Context, client user_enter.UserEnterClient, req *user_enter.CalculateOrderSummaryRequest) (*user_enter.CalculateOrderSummaryResponse, error) {
		return client.CalculateOrderSummary(ctx, req)
	})
}
