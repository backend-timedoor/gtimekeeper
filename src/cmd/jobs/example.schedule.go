package jobs

import (
	"context"
	"fmt"
	"time"

	"github.com/hibiken/asynq"
)

type ExampleSchedule struct{}

func (s *ExampleSchedule) Signature() string {
	return "example:schedule"
}

func (m *ExampleSchedule) Options() []asynq.Option {
	return []asynq.Option{
		// asynq.ProcessIn(5 * time.Second),
	}
}

func (s *ExampleSchedule) Schedule() string {
	return "@every 3s"
}

func (s *ExampleSchedule) Handle(ctx context.Context, t *asynq.Task) error {
	fmt.Println("Cron job executed at:", time.Now())

	return nil
}
