package kafka

import "github.com/backend-timedoor/gtimekeeper-framework/base/kafka"

const (
	ExampleTopic kafka.KafkaTopic = "example-topic"
)

type ExampleStruct struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
