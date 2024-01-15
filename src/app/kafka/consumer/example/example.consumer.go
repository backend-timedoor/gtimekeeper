package example

import (
	"fmt"

	"github.com/segmentio/kafka-go"
)

type ConsumerExample struct{}

func (c *ConsumerExample) Topic() string {
	return "example-topic"
}

func (c *ConsumerExample) Group() string {
	return "example-group-id" // use project service id for group id
}

func (c *ConsumerExample) Handle(m kafka.Message) {
	// fmt.Println("Message received:", message)
	fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
}
