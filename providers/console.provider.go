package providers

import (
	cmd "gtimekeeper/src/cmd/console"

	"github.com/backend-timedoor/gtimekeeper-framework/base/console"
	"github.com/backend-timedoor/gtimekeeper-framework/base/contracts"
)

type ConsoleServiceProvider struct{}

func (p *ConsoleServiceProvider) Boot() {}

func (p *ConsoleServiceProvider) Register() {
	console.BootConsole([]contracts.Commands{
		&cmd.ExampleCommand{},
		//new comment here
	})
}
