package rabbitmq

import (
	"crypto/tls"

	"github.com/streadway/amqp"
	"github.com/xALEGORx/go-expression-calculator/pkg/config"
)

type IBroker struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	Queue      string
}

var broker *IBroker

func Init() (*IBroker, error) {
	appConfig := config.Get()
	cfg := new(tls.Config)
	cfg.InsecureSkipVerify = true

	conn, err := amqp.DialTLS(appConfig.RabbitURL, cfg)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	broker = &IBroker{
		Connection: conn,
		Channel:    ch,
		Queue:      appConfig.RabbitQueue,
	}

	return broker, nil
}

func (b *IBroker) InitQueue() error {
	_, err := b.Channel.QueueDeclare(
		b.Queue, // queue name
		true,    // durable
		false,   // auto delete
		false,   // exclusive
		false,   // no wait
		nil,     // arguments
	)

	return err
}

func (b *IBroker) SendTask(task string) error {
	message := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(task),
	}

	err := b.Channel.Publish(
		"",
		b.Queue,
		false,
		false,
		message,
	)

	return err
}

func Get() *IBroker {
	return broker
}
