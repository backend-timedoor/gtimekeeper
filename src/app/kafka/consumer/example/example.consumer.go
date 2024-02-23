package example

import (
	"context"
	"encoding/json"
	"fmt"

	pkgKafka "gtimekeeper/src/_common/pkg/kafka"

	"github.com/backend-timedoor/gtimekeeper-framework/base/kafka"
	baseKafka "github.com/segmentio/kafka-go"
)

type ConsumerExample struct{}

func (c *ConsumerExample) Config() *[]kafka.ModuleConfig {
	return &[]kafka.ModuleConfig{
		{
			Reader: baseKafka.ReaderConfig{
				Topic: string(pkgKafka.ExampleTopic),
			},
			Handle: c.Handle,
		},
	}
}

// func (c *ConsumerExample) HandleUserCreate(ctx context.Context, m pkgKafka.Message, r *pkgKafka.Reader) error {
// 	if err := errors.New("error"); err != nil {
// 		return err
// 	}
// 	// panic("not implemented")
// 	fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))

// 	// r.CommitMessages(ctx, m)

// 	return nil
// }

func (c *ConsumerExample) Handle(ctx context.Context, m baseKafka.Message, r *baseKafka.Reader) error {
	fmt.Printf("handle default message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), m.Value)
	var data pkgKafka.ExampleStruct

	err := json.Unmarshal(m.Value, &data)

	fmt.Println("data", data)
	return err
}
