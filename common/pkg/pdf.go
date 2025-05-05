package pkg

import (
	"common/model"
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"log"
)

// GenerateInvoiceImage 生成包含发票信息的图片并保存为PDF到指定位置
func GenerateInvoiceImage(invoices *model.InvoiceApplication) error {

	// 创建一个新的PDF对象
	pdf := gofpdf.New("P", "mm", "A4", "")
	// 添加一个页面
	pdf.AddPage()
	pdf.AddFont("方正粗宋简体", "B", "../../common/fonts/方正粗宋简体_2.ttf")

	pdf.SetFont("方正粗宋简体", "B", 16)
	// 设置字体
	//pdf.SetFont("../../common/fonts/方正粗宋简体_2.TTF", "B", 16)

	// 写入标题
	pdf.Cell(40, 10, "发票")
	pdf.Ln(20)

	// 发票信息
	identificationNumber := invoices.TaxpayerIdentificationNumber
	pdf.Cell(40, 10, "发票编号: "+identificationNumber)
	pdf.Ln(10)
	reviewTime := invoices.ReviewTime.Format("20060102150405")
	pdf.Cell(40, 10, "开票日期: "+reviewTime)
	pdf.Ln(10)
	title := invoices.InvoiceTitle
	pdf.Cell(40, 10, "客户名称: "+title)
	pdf.Ln(10)

	// 表格头部
	pdf.SetFont("../../common/fonts/方正粗宋简体_2.TTF", "B", 12)
	pdf.Cell(40, 10, "商品名称")
	pdf.Cell(40, 10, "数量")
	pdf.Cell(40, 10, "单价")
	pdf.Cell(40, 10, "金额")
	pdf.Ln(10)
	invoiceAmount := invoices.InvoiceAmount
	sprintf := fmt.Sprintf("%.2f", invoiceAmount)

	// 表格内容
	pdf.SetFont("../../common/fonts/方正粗宋简体_2.TTF", "", 12)
	pdf.Cell(40, 10, "商品1")
	pdf.Cell(40, 10, "2")
	pdf.Cell(40, 10, "100.00")
	pdf.Cell(40, 10, sprintf)
	pdf.Ln(10)

	// 总计
	pdf.SetFont("../../common/fonts/方正粗宋简体_2.TTF", "B", 12)
	pdf.Cell(40, 10, "总计金额: "+sprintf)

	// 保存PDF到指定位置，使用完整的文件名
	pdfPath := "../common/picture/invoice_audit_result.pdf"
	err := pdf.OutputFileAndClose(pdfPath)
	if err != nil {
		log.Fatalf("生成PDF文件时出错: %v", err)
	}
	return nil
}
