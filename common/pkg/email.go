package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/utils"
	"strconv"
)

// 邮箱配置
const (
	smtpServer   = "smtp.qq.com"       //SMTP 服务器的地址
	smtpPort     = "587"               //SMTP 服务器的地址和端口
	smtpUsername = "2474179612@qq.com" //SMTP 用户名，是发件人邮箱地址
	smtpPassword = "dbixyvhdjaj"       // SMTP 密码
	recipient    = "2159685130@qq.com" //是收件人邮箱地址列表
)

func SendEmail(context string) (string, error) {
	emailConfig := EmailConfig()
	emailConn := utils.NewEMail(emailConfig)
	//接受者
	toEmail := recipient //c.GetString("to_email")
	subject := "Hi"      //c.GetString("subject") //标题
	text := context      //c.GetString("text")       //内容
	emailConn.To = []string{toEmail}
	emailConn.Subject = subject
	emailConn.Text = text
	error := emailConn.Send()
	if error != nil {
		fmt.Println(error.Error())
		fmt.Println("123456789o87654")
	}
	return "成功", nil
}

// 发送者
func EmailConfig() string { //QQ邮箱
	emailUserName := smtpUsername
	emailPassword := smtpPassword
	emailPort := smtpPort
	prot, _ := strconv.Atoi(emailPort)
	emailHost := smtpServer
	emailForm := smtpUsername

	Config := map[string]interface{}{
		"Username": emailUserName,
		"Password": emailPassword,
		"Host":     emailHost,
		"Port":     prot,
		"Form":     emailForm,
	}
	emailConfigStr, _ := json.Marshal(Config)
	return string(emailConfigStr)
}
