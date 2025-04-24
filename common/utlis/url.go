package utlis

import (
	"fmt"
	"github.com/nfnt/resize"
	"github.com/skip2/go-qrcode"
	"image"
	"image/png"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

// downloadImage 从网络下载图片并保存到本地
func downloadImage(url, filePath string) error {
	// 发起 HTTP 请求下载图片
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 检查 HTTP 响应状态
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download image: %s", resp.Status)
	}

	// 将图片保存到本地文件
	outFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, resp.Body)
	return err
}

// ChatUrl 生成带有用户头像的二维码图片，并返回图片的访问链接
func ChatUrl(uid int64, imgURL string) string {
	// 将用户ID转换为字符串，用于生成二维码内容
	content := strconv.FormatInt(uid, 10)

	// 生成基础二维码
	qr, err := qrcode.New(content, qrcode.Medium)
	if err != nil {
		return fmt.Sprintf("Failed to generate QR code: %v", err)
	}

	// 创建二维码图片
	qrCodeImage := qr.Image(256)

	// 创建画布
	bounds := qrCodeImage.Bounds()
	outImage := image.NewRGBA(bounds)

	// 将二维码绘制到画布上
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			outImage.Set(x, y, qrCodeImage.At(x, y))
		}
	}

	// 下载头像图片
	tempLogoPath := "D:\\gopath\\src\\retailers\\srv\\distribution_srv\\img\\" + "temp_logo.png"
	if err := downloadImage(imgURL, tempLogoPath); err != nil {
		return fmt.Sprintf("Failed to download logo image: %v", err)
	}

	// 加载用户头像
	logoFile, err := os.Open(tempLogoPath)
	if err != nil {
		return fmt.Sprintf("Failed to open logo image: %v", err)
	}
	defer logoFile.Close()

	logoImage, _, err := image.Decode(logoFile)
	if err != nil {
		return fmt.Sprintf("Failed to decode logo image: %v", err)
	}

	// 调整头像大小
	logoImage = resize.Resize(64, 64, logoImage, resize.Lanczos3)

	// 计算头像在二维码中的位置
	logoX := (bounds.Max.X - logoImage.Bounds().Max.X) / 2
	logoY := (bounds.Max.Y - logoImage.Bounds().Max.Y) / 2

	// 将头像绘制到二维码上
	for y := 0; y < logoImage.Bounds().Max.Y; y++ {
		for x := 0; x < logoImage.Bounds().Max.X; x++ {
			outImage.Set(logoX+x, logoY+y, logoImage.At(x, y))
		}
	}

	// 生成文件名
	fileName := "invite_qr_" + strconv.FormatInt(uid, 10) + ".png"
	filePath := filepath.Join("images", fileName)

	// 确保保存路径存在
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return fmt.Sprintf("Failed to create directory: %v", err)
	}

	// 保存最终的二维码图片
	outFile, err := os.Create(filePath)
	if err != nil {
		return fmt.Sprintf("Failed to create file: %v", err)
	}
	defer outFile.Close()

	if err := png.Encode(outFile, outImage); err != nil {
		return fmt.Sprintf("Failed to encode image: %v", err)
	}

	// 删除临时头像文件
	os.Remove(tempLogoPath)

	// 返回图片的访问链接
	imageUrl := "http://example.com/images/" + fileName
	return imageUrl
}
