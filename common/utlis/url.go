package utlis

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"image"
	"image/draw"
	"image/png"
	"os"
)

// resizeImage 函数用于调整图像大小
func resizeImage(img image.Image, width, height int) image.Image {
	resized := image.NewRGBA(image.Rect(0, 0, width, height))
	//draw.NearestNeighbor.Scale(resized, resized.Bounds(), img, img.Bounds(), draw.Over, nil)
	return resized
}

const KeyOrder = "user_%d:img_%d"

func ChatUrl(uid, str int64) {
	// 要编码到二维码中的内容
	content := fmt.Sprintf(KeyOrder, uid, str)
	// 二维码的纠错级别，这里选择高纠错级别以保证添加头像后仍能正常扫描
	errorCorrectionLevel := qrcode.High
	// 二维码的边长（像素）
	size := 256

	// 生成二维码
	qr, err := qrcode.New(content, errorCorrectionLevel)
	if err != nil {
		panic(err)
	}
	qrImg := qr.Image(size)

	// 打开头像文件，这里假设头像文件名为 avatar.jpg，你可按需修改
	avatarFile, err := os.Open(fmt.Sprintf(KeyOrder, uid, str))
	if err != nil {
		panic(err)
	}
	defer avatarFile.Close()

	// 解码头像图片
	avatarImg, _, err := image.Decode(avatarFile)
	if err != nil {
		panic(err)
	}

	// 调整头像大小，这里将头像大小设置为二维码边长的 20%
	avatarWidth := size / 5
	avatarHeight := size / 5
	resizedAvatar := resizeImage(avatarImg, avatarWidth, avatarHeight)

	// 计算头像在二维码中心的位置
	x := (qrImg.Bounds().Dx() - avatarWidth) / 2
	y := (qrImg.Bounds().Dy() - avatarHeight) / 2

	// 创建一个新的图像用于合成二维码和头像
	finalImg := image.NewRGBA(qrImg.Bounds())
	// 将二维码绘制到新图像上
	draw.Draw(finalImg, qrImg.Bounds(), qrImg, image.Point{}, draw.Src)
	// 将调整好大小的头像绘制到二维码中心位置
	draw.Draw(finalImg, image.Rect(x, y, x+avatarWidth, y+avatarHeight), resizedAvatar, image.Point{}, draw.Over)

	// 创建输出文件，这里将合成后的图片保存为 qrcode_with_avatar.png
	outputFile, err := os.Create(fmt.Sprintf(KeyOrder, uid, str))
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	// 将合成后的图像以 PNG 格式编码并保存到文件中
	err = png.Encode(outputFile, finalImg)
	if err != nil {
		panic(err)
	}

}
