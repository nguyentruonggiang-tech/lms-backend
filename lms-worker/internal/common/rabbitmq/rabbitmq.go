package rabbitmq

import (
	"context"
	"fmt"

	"lms-worker/internal/common/env"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	Conn *amqp.Connection
}

func NewRabbitMQ(e *env.Env) *RabbitMQ {
	conn, err := amqp.DialConfig(e.RabbitMQURL, amqp.Config{
		Properties: amqp.Table{
			"connection_name": "lms-worker",
		},
	})
	if err != nil {
		panic(fmt.Sprintf("❌ [RABBITMQ] failed to connect: %v", err))
	}

	fmt.Println("✅ [RABBITMQ] Connected successfully")
	return &RabbitMQ{Conn: conn}
}

func (r *RabbitMQ) On(ctx context.Context, queueName string, handler func(context.Context, []byte) error) (err error) {
	ch, err := r.Conn.Channel()
	if err != nil {
		fmt.Println(err)
		return
	}

	q, err := ch.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		defer ch.Close()
		for d := range msgs {
			if err := handler(ctx, d.Body); err != nil {
				d.Nack(false, false)
				continue
			}
			d.Ack(false)
		}
	}()

	return
}
