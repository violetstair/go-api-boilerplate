package boilerplate_queue

import (
	"log"
	"github.com/streadway/amqp"
)

func GetQueue(queuename string, object func([]byte)) {
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

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	go func() {
		for d := range msgs {
			object(d.Body)
		}
	}()

	return
}
