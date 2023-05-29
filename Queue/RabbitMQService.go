package queue

import (
	"fmt"

	"github.com/streadway/amqp"
)

type rabbitmqService struct {
	Conn *amqp.Connection
	Ch   *amqp.Channel
}

func NewRabbitMQService() *rabbitmqService {
	new_conn := CreateConn()
	new_ch := CreateChannel(*new_conn)
	return &rabbitmqService{Conn: new_conn, Ch: new_ch}
}

func CreateConn() *amqp.Connection {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return conn
}

func CreateChannel(conn amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return ch
}
