package pkg

import (
	"common/model"
	"github.com/fogleman/gg"
	"github.com/jung-kurt/gofpdf"
)

// GenerateInvoiceImage 生成包含发票信息的图片并保存为 PDF 到指定位置
func GenerateInvoiceImage(invoices *model.InvoiceApplication) error {
	// 创建一个新的绘图上下文
	dc := gg.NewContext(1000, 500) // 根据发票数量调整高度
	// 设置背景颜色为白色
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	// 设置字体，使用绝对路径加载字体文件
	fontPath := "D:\\Words\\方正粗宋简体_2.TTF" // 根据实际情况修改
	if err := dc.LoadFontFace(fontPath, 20); err != nil {
		return err
	}

	// 设置文本颜色为黑色
	dc.SetRGB(0, 0, 0)
	format := invoices.ReviewTime.Format("20060102150405")
	// 准备要绘制的文本
	text := "发票编号: " + invoices.TaxpayerIdentificationNumber + "\n审核日期: " + format + "\n审核结果: 审核通过"

	// 计算文本的宽度和高度
	textWidth, textHeight := dc.MeasureMultilineString(text, 2)
	x := (400 - textWidth) / 2
	y := (float64(200*(1)) - textHeight) / 2

	// 在图片上绘制文本
	dc.DrawStringAnchored(text, x, y, 0, 0.5)

	// 保存图片，使用绝对路径
	imagePath := "D:\\gocode\\src\\retailers\\common\\picture\\invoice_audit_result.png"
	if err := dc.SavePNG(imagePath); err != nil {
		return err
	}

	// 创建一个新的 PDF 文档
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// 插入生成的图片到 PDF
	pdf.ImageOptions(imagePath, 10, 10, 190, 0, false, gofpdf.ImageOptions{ImageType: "PNG", ReadDpi: true}, 0, "")

	// 保存 PDF 到指定位置，使用完整的文件名
	pdfPath := "D:\\gocode\\src\\retailers\\common\\picture\\invoice_audit_result.pdf"
	return pdf.OutputFileAndClose(pdfPath)
}
