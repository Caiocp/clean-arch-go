package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/caiocp/clean-arch-go/pkg/events"
	amqp "github.com/rabbitmq/amqp091-go"
)

type OrderCreatedHandler struct {
	RabbitMQChannel *amqp.Channel
}

func NewOrderCreatedHandler(rabbitMQChannel *amqp.Channel) *OrderCreatedHandler {
	return &OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (h *OrderCreatedHandler) Handle(event events.EventInterface, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()

	fmt.Println("Order created: ", event.GetPayload())

	jsonOutput, _ := json.Marshal(event.GetPayload())

	rabbitMqMessage := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	h.RabbitMQChannel.Publish("amq.direct", "", false, false, rabbitMqMessage)
}
