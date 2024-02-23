package providers

import (
	kPkg "gtimekeeper/src/_common/pkg/kafka"
	"gtimekeeper/src/app/kafka/consumer/example"

	"github.com/backend-timedoor/gtimekeeper-framework/app"
	"github.com/backend-timedoor/gtimekeeper-framework/base/kafka"
)

type KafkaServiceProvider struct{}

func (p *KafkaServiceProvider) Boot() {}

func (p *KafkaServiceProvider) Register() {
	topics := []kafka.Topic{
		{
			Topic:       kPkg.ExampleTopic,
			Partition:   1,
			Replication: 1,
		},
	}

	consumers := []kafka.Consumer{
		&example.ConsumerExample{},
	}

	app.Kafka = kafka.New(&kafka.Config{
		Brokers:          app.Config.GetString("KAFKA_BROKER", "localhost:9092"),
		Topics:           topics,
		Consumers:        consumers,
		ConsumerGroupID:  app.Config.GetString("KAFKA_CONSUMER_GROUP_ID", "gtimekeeper"),
		AutoCommitOffset: true,
		SchemaRegistry: kafka.SchemaRegistry{
			Host: app.Config.GetString("KAFKA_SCHEMA_REGISTRY_HOST", "http://localhost:8081"),
			Schemas: []kafka.Schema{
				{
					Subject: kPkg.ExampleTopic,
					Type:    "example",
					Schema:  "src/app/kafka/schema/create.schema.avsc",
				},
			},
		},
	})
}
