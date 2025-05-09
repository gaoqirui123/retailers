package pkg

import (
	"common/global"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v4/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

func CreateClient() (_result *dysmsapi20170525.Client, _err error) {
	Conf := global.NaCos.Aliyun
	config := &openapi.Config{
		// 必填，请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_ID。
		AccessKeyId: tea.String(Conf.AccessKeyId),
		// 必填，请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_SECRET。
		AccessKeySecret: tea.String(Conf.AccessKeySecret),
	}
	// Endpoint 请参考 https://api.aliyun.com/product/Dysmsapi
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

func SendSms(phone, code string) (res *dysmsapi20170525.SendSmsResponse, _err error) {
	client, _err := CreateClient()
	if _err != nil {
		return res, _err
	}

	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		SignName:      tea.String("阿里云短信测试"),
		TemplateCode:  tea.String("SMS_154950909"),
		PhoneNumbers:  tea.String(phone),
		TemplateParam: tea.String("{\"code\":" + code + "}"),
	}
	runtime := &util.RuntimeOptions{}

	// 复制代码运行请自行打印 API 的返回值
	res, _err = client.SendSmsWithOptions(sendSmsRequest, runtime)
	if _err != nil {
		return nil, _err
	}

	return res, nil

}
