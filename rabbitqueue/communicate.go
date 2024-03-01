package rabbitqueue

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"fmt"
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

func EncodeToBytes(p interface{}) []byte {

	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("uncompressed size (bytes): ", len(buf.Bytes()))
	return buf.Bytes()
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

func PublishMessage(ctx context.Context, message DataModel.User, vhostname string) {
	body, err := json.Marshal(message)
	if err != nil {
		fmt.Println(err)
		return
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
		log.Panicf("%s", erro)
	}
	log.Printf("Message passed to the Queue")
}
