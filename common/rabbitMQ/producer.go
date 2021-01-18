package rabbitMQ

import "github.com/streadway/amqp"

func Publish(exchange, routingKey string, msg []byte)bool{
	if nil == channel.Publish(
		exchange,
		routingKey,
		false, // 如果没有对应的queue, 就会丢弃这条消息
		false, //
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg}) {
		return true
	}
	return false
}
