package schedules

import (
	"fmt"
	"time"

	schedule "github.com/backend-timedoor/gtimekeeper-framework/base/scheduler"
)

type ExampleSchedule struct{}

func (s *ExampleSchedule) Signature() string {
	return "example:schedule"
}

func (s *ExampleSchedule) Schedule() string {
	return schedule.EveryMinute()
}

func (s *ExampleSchedule) Handle() {
	fmt.Println("Cron job executed at:", time.Now())
}