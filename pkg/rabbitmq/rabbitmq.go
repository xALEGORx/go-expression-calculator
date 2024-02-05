package rabbitmq

import (
	"crypto/tls"

	"github.com/streadway/amqp"
	"github.com/xALEGORx/go-expression-calculator/pkg/config"
)

type ibroker struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

var broker *ibroker

func Init() error {
	cfg := new(tls.Config)
	cfg.InsecureSkipVerify = true

	conn, err := amqp.DialTLS(config.Get().RabbitURL, cfg)
	if err != nil {
		return err
	}
	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	broker = &ibroker{
		Connection: conn,
		Channel:    ch,
	}

	return nil
}

func Get() *ibroker {
	return broker
}
