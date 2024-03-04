package rabbitqueue

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	amqp "github.com/rabbitmq/amqp091-go"
	DataModel "github.com/rohit123sinha456/digitalSignage/model"
)

var conn *amqp.Connection
var ch *amqp.Channel

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func Connect(username string, password string, vhostname string) {
	rabbiturl := strings.Join([]string{"amqp://", username, ":", password, "@localhost:5672/", vhostname}, "")
	conn, err := amqp.Dial(rabbiturl) //"amqp://guest:guest@localhost:5672/"
	failOnError(err, "Failed to connect to RabbitMQ")
	// defer conn.Close()
	ch, err = conn.Channel()
	failOnError(err, "Failed to open a channel")
	// defer ch.Close()
}

func PublishMessage(ctx context.Context, message DataModel.Playlist, vhostname string) error {
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}
	erro := ch.PublishWithContext(
		ctx,
		"PLExchange", // exchange
		"",           // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	if erro != nil {
		return err
	}
	return nil
}
