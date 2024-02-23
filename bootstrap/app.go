package bootstrap

import (
	"gtimekeeper/providers"

	"github.com/backend-timedoor/gtimekeeper-framework/base/contracts"
	"github.com/backend-timedoor/gtimekeeper-framework/boot"
)

func Boot() {
	boot.Booting([]contracts.ServiceProvider{
		&providers.ConfigServiceProvider{},
		&providers.LogServiceProvider{},
		&providers.AppServiceProvider{},
		&providers.ValidationServiceProvider{},
		&providers.DatabaseServiceProvider{},
		&providers.ConsoleServiceProvider{},
		&providers.ServerServiceProvider{},
		&providers.JobServiceProvider{},
		// &providers.KafkaServiceProvider{},
		&providers.MailServiceProvider{},
	})
}
