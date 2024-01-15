package providers

import (
	"github.com/backend-timedoor/gtimekeeper-framework/app"
	"github.com/backend-timedoor/gtimekeeper-framework/base/contracts"
	schedule "github.com/backend-timedoor/gtimekeeper-framework/base/scheduler"
)

type ScheduleServiceProvider struct{}

func (log *ScheduleServiceProvider) Boot() {}

func (log *ScheduleServiceProvider) Register() {
	app.Schedule = schedule.BootSchedule([]contracts.ScheduleEvent{
		// &schedules.ExampleSchedule{},
		// other schedule
	})
}
