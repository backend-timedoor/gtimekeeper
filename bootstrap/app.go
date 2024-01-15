package bootstrap

import (
	"gtimekeeper/providers"

	"github.com/backend-timedoor/gtimekeeper-framework/base/contracts"
	"github.com/backend-timedoor/gtimekeeper-framework/boot"
)

func Boot() {
	boot.Booting([]contracts.ServiceProvider{
		&providers.ConfigServiceProvider{},
		&providers.AppServiceProvider{},
		&providers.ConsoleServiceProvider{},
		&providers.ServerServiceProvider{},
		&providers.QueueServiceProvider{},
		&providers.ScheduleServiceProvider{},
		//&providers.KafkaServiceProvider{},
	})
}
