package utlis

import (
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"image"
	"os"
)

func DecodeQRCode(filePath string) (string, error) {
	// 打开二维码图片文件
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 解码图片
	img, _, err := image.Decode(file)
	if err != nil {
		return "", err
	}

	// 将图片转换为 gozxing 可处理的格式
	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		return "", err
	}

	// 创建二维码阅读器
	reader := qrcode.NewQRCodeReader()

	// 解码二维码
	result, err := reader.Decode(bmp, nil)
	if err != nil {
		return "", err
	}

	return result.GetText(), nil
}
