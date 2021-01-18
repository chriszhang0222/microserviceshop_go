package rabbitMQ

import "go.uber.org/zap"

var done chan bool

func StartConsume(qName, cName string, callback func (msg []byte)bool){
	msgs,err := channel.Consume(qName, cName, true,  //自动应答
		false, // 非唯一的消费者
		false, // rabbitMQ只能设置为false
		false, // noWait, false表示会阻塞直到有消息过来
		nil)
	if err != nil {
		zap.S().Error(err)
		return
	}
	done = make(chan bool)
	go func() {
		for d := range msgs {
			processErr := callback(d.Body)
			if processErr {
				// TODO: 将任务写入错误队列，待后续处理
			}
		}
	}()
	<-done
	channel.Close()
}

func StopConsume() {
	done <- true
}
