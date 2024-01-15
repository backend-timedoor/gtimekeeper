package providers

import (
	"github.com/backend-timedoor/gtimekeeper-framework/app"
	"github.com/backend-timedoor/gtimekeeper-framework/base/queue"
)

type QueueServiceProvider struct{}

func (log *QueueServiceProvider) Boot() {}

func (log *QueueServiceProvider) Register() {
	app.Queue = queue.BootQueue()
}