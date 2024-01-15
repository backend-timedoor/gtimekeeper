package providers

import (
	"time"

	"github.com/backend-timedoor/gtimekeeper-framework/app"
)

type AppServiceProvider struct{}

func (p *AppServiceProvider) Boot() {
	timezone := app.Config.GetString("APP_TIMEZONE", "Asia/Makassar")
	loc, _ := time.LoadLocation(timezone)
	time.Local = loc
}

func (p *AppServiceProvider) Register() {}
