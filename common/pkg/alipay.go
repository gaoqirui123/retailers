package pkg

import (
	"common/global"
	"fmt"
	"github.com/smartwalle/alipay/v3"
)

type Payment interface {
	Pay(subject, orderSn, amount string) string
}

type Alipay struct {
	AppId      string
	PrivateKey string
	NotifyURL  string
	ReturnURL  string
}

func NewPay() *Alipay {
	return &Alipay{
		AppId:      global.Config.AliPay.APPID,
		PrivateKey: global.Config.AliPay.Key,
		NotifyURL:  global.Config.AliPay.NotifyURL,
		ReturnURL:  global.Config.AliPay.ReturnURL,
	}
}

func (a *Alipay) Pay(subject, orderSn, amount string) string {
	var privateKey = a.PrivateKey // 必须，上一步中使用 RSA签名验签工具 生成的私钥
	var client, err = alipay.New(a.AppId, privateKey, false)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	var p = alipay.TradeWapPay{}
	p.NotifyURL = a.NotifyURL
	p.ReturnURL = a.ReturnURL
	p.Subject = subject
	p.OutTradeNo = orderSn
	p.TotalAmount = amount
	p.ProductCode = "QUICK_WAP_WAY"

	url, err := client.TradeWapPay(p)
	if err != nil {
		fmt.Println(err)
	}

	// 这个 payURL 即是用于打开支付宝支付页面的 URL，可将输出的内容复制，到浏览器中访问该 URL 即可打开支付页面。
	var payURL = url.String()
	return payURL
}
