package queue

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"github.com/Lanquill/Forge/pkg/db"
	"github.com/streadway/amqp"
)

var conn *amqp.Connection

func InitializeQueue() {
	var err error
	c := make(chan *amqp.Error)
	go func() {
		err := <-c
		log.Println("reconnect: " + err.Error())
		InitializeQueue()
	}()

	conn, err = amqp.Dial(db.RabbitMQConnection)
	if err != nil {
		panic("cannot connect")
	}
	conn.NotifyClose(c)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func ServiceMQ(chName string, serviceMessage map[string]interface{}) int {

	serviceMessageJSON, err := json.Marshal(serviceMessage)
	if err != nil {
		failOnError(err, "Failed to JSON encode")
	}

	var score int

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // noWait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	corrId := randomString(32)
	//

	err = ch.Publish(
		"",     // exchange
		chName, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: corrId,
			ReplyTo:       q.Name,
			Body:          []byte(serviceMessageJSON),
		})
	failOnError(err, "Failed to publish a message")

	if len(msgs) != 0 {
		for d := range msgs {
			if corrId == d.CorrelationId {
				// returnMap["res"] = string(d.Body)
				err = json.Unmarshal(d.Body, &score)
				if err != nil {
					return score
				}
				break
			}
		}
	}

	return score
}
