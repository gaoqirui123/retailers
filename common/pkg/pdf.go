package pkg

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
)

// AddPdf 调用 GeneratePdf 生成 PDF 文件
func AddPdf() error {
	filename := "invoice_list.pdf"
	err := GeneratePdf(filename)
	if err != nil {
		return fmt.Errorf("生成 PDF 文件 %s 时出错: %w", filename, err)
	}
	return nil
}

// GeneratePdf 生成 PDF 文件，包含文本和图片
func GeneratePdf(filename string) error {
	// 创建一个新的 PDF 对象
	pdf := gofpdf.New("P", "mm", "A4", "")
	// 添加一个新页面
	pdf.AddPage()
	// 设置字体
	pdf.SetFont("Arial", "B", 16)
	// 添加文本内容
	pdf.CellFormat(190, 7, "Welcome to topgoer.com", "0", 0, "CM", false, 0, "")

	// 添加图片
	imagePath := "topgoer.png"
	pdf.ImageOptions(
		imagePath,
		80, 20,
		0, 0,
		false,
		gofpdf.ImageOptions{ImageType: "PNG", ReadDpi: true},
		0,
		"",
	)
	// 保存并关闭 PDF 文件
	err := pdf.OutputFileAndClose(filename)
	if err != nil {
		return fmt.Errorf("保存 PDF 文件 %s 时出错: %w", filename, err)
	}

	return nil
}
