package rabbitMQ
import (
	"fmt"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
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


func InitChannel()bool{
	if channel != nil {
		return true
	}
	conn, err := amqp.Dial(fmt.Sprintf("amqp://guest:guest@%s:%d/", Config.Host, Config.Port))
	if err != nil {
		zap.S().Error(err)
		return false
	}
	channel, err = conn.Channel()
	if err != nil {
		zap.S().Error(err)
		return false
	}
	return true
}