package queue

import (
	"fmt"
	"runtime"
)

var (
	QUEUE_PROTO    = "amqp"
	QUEUE_ACCOUNT  = "account"
	QUEUE_PASSWORD = "password"
	QUEUE_SERVER   = "127.0.0.1"
	QUEUE_PORT     = "5672"

	QUEUE_DEV_PROTO    = "amqp"
	QUEUE_DEV_ACCOUNT  = "account"
	QUEUE_DEV_PASSWORD = "password"
	QUEUE_DEV_SERVER   = "127.0.0.1"
	QUEUE_DEV_PORT     = "5672"
)

func getQueueEndPoint() (connUrl string) {
	connUrl = fmt.Sprintf(
		"%s://%s:%s@%s:%s/",
		QUEUE_PROTO,
		QUEUE_ACCOUNT,
		QUEUE_PASSWORD,
		QUEUE_SERVER,
		QUEUE_PORT,
	)

	if runtime.GOOS == "darwin" {
		connUrl = fmt.Sprintf(
			"%s://%s:%s@%s:%s/",
			QUEUE_DEV_PROTO,
			QUEUE_DEV_ACCOUNT,
			QUEUE_DEV_PASSWORD,
			QUEUE_DEV_SERVER,
			QUEUE_DEV_PORT,
		)
	}

	return
}
