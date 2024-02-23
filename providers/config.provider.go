package providers

import (
	"os"

	"github.com/backend-timedoor/gtimekeeper-framework/app"
	"github.com/backend-timedoor/gtimekeeper-framework/base/config"
)

type ConfigServiceProvider struct{}

func (p *ConfigServiceProvider) Boot() {
	initConfig()
}

func (p *ConfigServiceProvider) Register() {
	app.Config = config.New(&config.Configuration{})
}

func initConfig() {
	pwd, _ := os.Getwd()
	config := app.Config

	// app path
	config.Add("path", map[string]any{
		"root": pwd,
		"app":  config.Env("APP_PATH", "src/app"),
		"view": config.Env("VIEW_PATH", "src/view"),
		"mail": config.Env("MAIL_PATH", "src/view/mail"),
	})

	// app config
	config.Add("app", map[string]any{
		"name": config.Env("APP_NAME", "GtTimeKeeper"),
		"env":  config.Env("APP_ENV", "development"),
		"host": config.Env("APP_HOST", "127.0.0.1:3000"),
	})
}
