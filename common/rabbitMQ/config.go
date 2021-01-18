package rabbitMQ
import (
	"github.com/streadway/amqp"
)
type Rabbitconfig struct{
	Host string
	Port int
	Routingkey string
	ExchangeName string
	QueueName string
}

var Config = &Rabbitconfig{}
var conn *amqp.Connection
var channel *amqp.Channel
