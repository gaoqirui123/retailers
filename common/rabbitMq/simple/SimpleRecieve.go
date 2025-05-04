package simple

// Receive 是一个函数，用于从 RabbitMQ 队列中接收消息

func Receive() {
	// 创建一个 RabbitMQ 实例
	// NewRabbitMQSimple 是一个函数，用于初始化一个简单的 RabbitMQ 连接
	// 参数 "kuteng" 是队列的名称
	rabbitmq := NewRabbitMQSimple("kuteng")

	// 调用 ConsumeSimple 方法，从队列中消费消息
	// ConsumeSimple 方法会持续监听队列，并在收到消息时进行处理
	rabbitmq.ConsumeSimple()
	return
}
