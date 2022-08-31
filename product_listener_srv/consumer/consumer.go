package consumer

import (
	json "github.com/bytedance/sonic"
	"github.com/streadway/amqp"
	"log"
	"product_listener_srv/models"
	"product_listener_srv/repository"
)

type Consumer struct {
	r  repository.Repository
	ch *amqp.Channel
}

func New(r repository.Repository, ch *amqp.Channel) *Consumer {
	return &Consumer{
		r:  r,
		ch: ch,
	}
}

func (c Consumer) AddProduct() {
	// Subscribing to QueueService1 for getting messages.
	messages, err := c.ch.Consume(
		"product_add", // queue name
		"",            // consumer
		true,          // auto-ack
		false,         // exclusive
		false,         // no local
		false,         // no wait
		nil,           // arguments
	)
	if err != nil {
		log.Println(err)
	}

	go func() {
		for message := range messages {
			// For example, show received message in a console.
			var data models.Product
			err := json.Unmarshal(message.Body, &data)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(data)
			err = c.r.Add(&data)
			if err != nil {
				log.Println(err)
			}
		}
	}()

}
