package jobs

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
)

type ExampleQueue struct{}

func (m *ExampleQueue) Signature() string {
	return "example:job"
}

func (m *ExampleQueue) Options() []asynq.Option {
	return []asynq.Option{
		// asynq.ProcessIn(5 * time.Second),
	}
}

func (m *ExampleQueue) Handle(ctx context.Context, t *asynq.Task) error {
	var data map[string]interface{}
	json.Unmarshal(t.Payload(), &data)
	fmt.Println("example queue executed:", data)

	return nil
}
