package simple

import (
	"fmt"
)

// Publish 是一个函数，用于向 RabbitMQ 队列发送消息

func Publish(data string) {
	// 创建一个 RabbitMQ 实例
	// NewRabbitMQSimple 是一个函数，用于初始化一个简单的 RabbitMQ 连接
	// 参数 "kuteng" 是队列的名称
	rabbitmq := NewRabbitMQSimple("kuteng")

	// 调用 PublishSimple 方法，向队列发送消息

	rabbitmq.PublishSimple(data)

	// 打印发送成功的提示信息
	fmt.Println("发送成功！")
}
