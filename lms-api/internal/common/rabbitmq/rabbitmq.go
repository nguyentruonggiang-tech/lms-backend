package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"

	"lms-api/internal/common/env"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	Conn *amqp.Connection
}

func NewRabbitMQ(e *env.Env) *RabbitMQ {
	conn, err := amqp.DialConfig(e.RabbitMQURL, amqp.Config{
		Properties: amqp.Table{
			"connection_name": "lms-api",
		},
	})
	if err != nil {
		panic(fmt.Sprintf("❌ [RABBITMQ] failed to connect: %v", err))
	}

	fmt.Println("✅ [RABBITMQ] Connected successfully")
	return &RabbitMQ{Conn: conn}
}

func (r *RabbitMQ) Send(ctx context.Context, queueName string, payload any) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	ch, err := r.Conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		return err
	}

	return ch.PublishWithContext(ctx, "", q.Name, false, false, amqp.Publishing{
		ContentType:   "application/json",
		CorrelationId: randomString(32),
		DeliveryMode:  amqp.Persistent,
		Body:          body,
	})
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := range bytes {
		bytes[i] = byte(rand.Intn(26) + 65)
	}
	return string(bytes)
}
