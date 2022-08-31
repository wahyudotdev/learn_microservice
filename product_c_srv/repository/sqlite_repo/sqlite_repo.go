package sqlite_repo

import (
	"context"
	json "github.com/bytedance/sonic"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
	"product_c_srv/models"
)

type SqliteRepo struct {
	Db *gorm.DB
	Ch *amqp.Channel
}

func New(db *gorm.DB, ch *amqp.Channel) SqliteRepo {
	db.AutoMigrate(&models.Product{})
	_, err := ch.QueueDeclare(
		"product_add", // queue name
		true,          // durable
		false,         // auto delete
		false,         // exclusive
		false,         // no wait
		nil,           // arguments
	)
	if err != nil {
		panic(err)
	}
	return SqliteRepo{
		Db: db,
		Ch: ch,
	}
}

func (s SqliteRepo) Add(ctx context.Context, data *models.Product) error {
	data.Id = uuid.New().String()
	if err := s.Db.WithContext(ctx).Create(&data); err.Error != nil {
		return err.Error
	}

	jsonStr, err := json.MarshalString(&data)
	if err != nil {
		return err
	}
	message := amqp.Publishing{
		ContentType:  "application/json",
		Body:         []byte(jsonStr),
		DeliveryMode: amqp.Persistent,
	}
	// Attempt to publish a message to the queue.
	if err := s.Ch.Publish(
		"",            // exchange
		"product_add", // queue name
		false,         // mandatory
		false,         // immediate
		message,       // message to publish
	); err != nil {
		return err
	}
	return nil
}
