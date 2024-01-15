package providers

import (
	"gtimekeeper/src/app/kafka/consumer/example"

	"github.com/backend-timedoor/gtimekeeper-framework/app"
	"github.com/backend-timedoor/gtimekeeper-framework/base/contracts"
	"github.com/backend-timedoor/gtimekeeper-framework/base/kafka"
)

type KafkaServiceProvider struct{}

func (p *KafkaServiceProvider) Boot() {}

func (p *KafkaServiceProvider) Register() {
	consumers := []contracts.KafkaConsumer{
		&example.ConsumerExample{},
	}

	app.Kafka = kafka.BootKafka([]string{
		"example-topic-1",
	}, consumers)
}
