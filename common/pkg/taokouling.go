package pkg

import (
	"common/model"
	"encoding/base64"
	"fmt"
	"strings"
)

// GenerateTaokouling 生成淘口令
func GenerateTaokouling(group model.Pink) string {
	// 将拼团信息拼接成一个字符串
	info := fmt.Sprintf("%s|%s|%.2f|%d", group.Id, group.OrderId, group.Price, group.CurrentNum)
	// 对信息进行 Base64 编码
	encoded := base64.StdEncoding.EncodeToString([]byte(info))
	// 添加特殊标识作为淘口令
	taokouling := fmt.Sprintf("￥%s￥", encoded)
	return taokouling
}

// ParseTaokouling 解析淘口令
func ParseTaokouling(taokouling string) (model.Pink, error) {
	// 去除特殊标识
	taokouling = strings.TrimPrefix(taokouling, "￥")
	taokouling = strings.TrimSuffix(taokouling, "￥")
	// 对 Base64 编码的信息进行解码
	decoded, err := base64.StdEncoding.DecodeString(taokouling)
	if err != nil {
		return model.Pink{}, nil
	}
	// 分割信息字符串
	parts := strings.Split(string(decoded), "|")
	if len(parts) != 4 {
		return model.Pink{}, fmt.Errorf("invalid taokouling format")
	}
	var price float64
	var requiredNum int
	fmt.Sscanf(parts[2], "%f", &price)
	fmt.Sscanf(parts[3], "%d", &requiredNum)
	return model.Pink{
		Id:         0,
		Uid:        0,
		OrderId:    "",
		OrderIdKey: "",
		TotalNum:   0,
		TotalPrice: 0,
		Cid:        0,
		Pid:        0,
		People:     0,
		CurrentNum: 0,
		Price:      0,
		AddTime:    "",
		StopTime:   "",
		KId:        0,
		IsTpl:      0,
		IsRefund:   0,
		Status:     0,
	}, nil
}

func main() {
	// 解析淘口令
	//fmt.Printf("解析出的拼团信息: ID=%s, 商品名称=%s, 价格=%.2f, 成团人数=%d\n",
	//	parsedGroup.ID, parsedGroup.ProductName, parsedGroup.Price, parsedGroup.RequiredNum)

}
