package queue

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

func PushQueue(queuename string, obj interface{}) (err error) {
	data, err := json.Marshal(obj)
	if err != nil {
		log.Println(err)
	}

	conn_url := getQueueEndPoint()
	conn, err := amqp.Dial(conn_url)
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	ch, _ := conn.Channel()
	if err != nil {
		log.Println(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queuename,
		false,
		false,
		false,
		false,
		nil,
	)

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(data),
		})
	return
}
