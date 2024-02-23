package providers

import (
	"gtimekeeper/src/cmd/jobs"
	"time"

	"github.com/backend-timedoor/gtimekeeper-framework/app"
	"github.com/backend-timedoor/gtimekeeper-framework/base/contracts"
	"github.com/backend-timedoor/gtimekeeper-framework/base/job"
	"github.com/hibiken/asynq"
)

type JobServiceProvider struct{}

func (log *JobServiceProvider) Boot() {
	app.Job.RegisterSchedule([]contracts.Schedule{
		&jobs.ExampleSchedule{},
		//
	})

	app.Job.RegisterQueue([]contracts.Queue{
		// &jobs.ExampleQueue{},
	})
}

func (log *JobServiceProvider) Register() {
	app.Job = job.New(&job.Config{
		ServerOpt: &asynq.Config{},
		ScheduleOpt: &asynq.SchedulerOpts{
			Location: time.Local,
		},
	})
}
