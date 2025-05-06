package utlis

import (
	"github.com/skip2/go-qrcode"
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func GenerateQRCodeWithLogo(orderID string, logoPath string, outputPath string) error {
	// 生成基础二维码
	err := qrcode.WriteFile(orderID, qrcode.Medium, 256, outputPath)
	if err != nil {
		log.Printf("生成二维码时出错: %v", err)
		return err
	}

	// 打开基础二维码图片
	qrFile, err := os.Open(outputPath)
	if err != nil {
		log.Printf("打开二维码文件 %s 时出错: %v", outputPath, err)
		return err
	}
	qrImg, _, err := image.Decode(qrFile)
	qrFile.Close() // 确保关闭文件句柄
	if err != nil {
		log.Printf("解码二维码文件 %s 时出错: %v", outputPath, err)
		return err
	}
	log.Printf("二维码文件 %s 已成功解码", outputPath)

	// 打开中间的小图片（logo）
	logoFile, err := os.Open(logoPath)
	if err != nil {
		log.Printf("打开 logo 文件 %s 时出错: %v", logoPath, err)
		return err
	}
	logoImg, _, err := image.Decode(logoFile)
	logoFile.Close() // 确保关闭文件句柄
	if err != nil {
		log.Printf("解码 logo 文件 %s 时出错: %v", logoPath, err)
		return err
	}
	log.Printf("logo 文件 %s 已成功解码", logoPath)

	// 获取二维码图片的尺寸
	qrBounds := qrImg.Bounds()
	qrWidth := qrBounds.Dx()
	qrHeight := qrBounds.Dy()

	// 获取 logo 图片的尺寸
	logoBounds := logoImg.Bounds()
	logoWidth := logoBounds.Dx()
	logoHeight := logoBounds.Dy()

	// 计算 logo 在二维码中的位置，使其居中
	x := (qrWidth - logoWidth) / 2
	y := (qrHeight - logoHeight) / 2

	// 创建一个新的 RGBA 图片，用于合并二维码和 logo
	newImg := image.NewRGBA(qrBounds)
	draw.Draw(newImg, qrBounds, qrImg, qrBounds.Min, draw.Src)
	draw.Draw(newImg, image.Rect(x, y, x+logoWidth, y+logoHeight), logoImg, logoBounds.Min, draw.Src)

	// 保存合并后的图片
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, newImg)
	if err != nil {
		return err
	}

	return nil
}
