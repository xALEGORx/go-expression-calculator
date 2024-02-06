package rabbitmq

import (
	"crypto/tls"
	"github.com/streadway/amqp"
)

type IBroker struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	Queue      string
}

var broker *IBroker

func Init(dsn string, queue string) (*IBroker, error) {
	cfg := new(tls.Config)
	cfg.InsecureSkipVerify = true

	conn, err := amqp.DialTLS(dsn, cfg)
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
		Queue:      queue,
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

func (b *IBroker) ConnQueue() (<-chan amqp.Delivery, error) {
	messages, err := b.Channel.Consume(
		b.Queue, // queue name
		"",      // consumer
		true,    // auto-ack
		false,   // exclusive
		false,   // no local
		false,   // no wait
		nil,     // arguments
	)
	if err != nil {
		return nil, err
	}
	return messages, nil
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
