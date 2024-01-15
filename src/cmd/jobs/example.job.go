package jobs

import (
	"context"
	"fmt"
	"time"

	"github.com/hibiken/asynq"
)

type ExampleJob struct{}

func (m *ExampleJob) Signature() string {
	return "example:job"
}

func (m *ExampleJob) Options() []asynq.Option {
	return []asynq.Option{
		asynq.ProcessIn(5 * time.Second),
	}
}

func (m *ExampleJob) Handle(ctx context.Context, t *asynq.Task) error {
	// data := json.Unmarshal(t.Payload())
	fmt.Println("job example is run")

	return nil
}
