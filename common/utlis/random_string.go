package utlis

import (
	"common/global"
	"math/rand"
	"time"
)

// 生成邀请码
func GenerateInviteCode() string {
	const Charset = "qwertyuiopasdfghjklzxcvbnm1234567890QWERTYUIOPASDFGHJKLZXCVBNM"
	rand.Seed(time.Now().UnixNano())
	lens := global.Config.Log
	code := make([]byte, lens)
	for i := range code {
		code[i] = Charset[rand.Intn(len(Charset))]
	}
	return string(code)
}
