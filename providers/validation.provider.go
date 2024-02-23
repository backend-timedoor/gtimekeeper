package providers

import (
	"github.com/backend-timedoor/gtimekeeper-framework/app"
	"github.com/backend-timedoor/gtimekeeper-framework/base/validation"
)

type ValidationServiceProvider struct{}

func (p *ValidationServiceProvider) Boot() {}

func (p *ValidationServiceProvider) Register() {
	app.Validation = validation.New()
}
