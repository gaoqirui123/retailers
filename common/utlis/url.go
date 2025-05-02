package utlis


import (
	"github.com/skip2/go-qrcode"
	"os"
)

// GenerateQRCode 生成指定内容的二维码并保存为文件
func GenerateQRCode(content string, filePath string) error {
	// 创建二维码
	png, err := qrcode.Encode(content, qrcode.Medium, 256)
	if err != nil {
		return err
	}

	// 将二维码保存到文件
	err = os.WriteFile(filePath, png, 0644)
	if err != nil {
		return err
	}

	return nil
}