package providers

import (
	"github.com/backend-timedoor/gtimekeeper-framework/base/configuration"
	"os"

	"github.com/backend-timedoor/gtimekeeper-framework/app"
)

type ConfigServiceProvider struct{}

func (p *ConfigServiceProvider) Boot() {
	initConfig()
}

func (p *ConfigServiceProvider) Register() {
	app.Config = configuration.BootConfig()
}

func initConfig() {
	config := app.Config

	root, err := os.Getwd()
	if err != nil {
		app.Log.Fatal(err)
	}

	// app path
	config.Add("path", map[string]any{
		"root": root,
		"app":  config.Env("APP_PATH", "src/app"),
		"view": config.Env("VIEW_PATH", "src/view"),
		"mail": config.Env("MAIL_PATH", "src/view/mail"),
	})

	// app mail
	config.Add("mail", map[string]any{
		"driver":     config.Env("MAIL_DRIVER", "smtp"),
		"host":       config.Env("MAIL_HOST", "smtp.mailtrap.io"),
		"port":       config.Env("MAIL_PORT", 2525),
		"username":   config.Env("MAIL_USERNAME", ""),
		"password":   config.Env("MAIL_PASSWORD", ""),
		"encryption": config.Env("MAIL_ENCRYPTION", ""),
	})

	// app config
	config.Add("app", map[string]any{
		"name": config.Env("APP_NAME", "Microdemy"),
		"env":  config.Env("APP_ENV", "development"),
		"host": config.Env("APP_HOST", "127.0.0.1:3000"),
	})

	// database config
	defaultDB := config.GetString("DB_CONNECTION", "mysql")
	config.Add("database", map[string]any{
		"connection": defaultDB,
		defaultDB: map[string]any{
			"host":     config.Env("DB_HOST", "127.0.0.1"),
			"port":     config.Env("DB_PORT", "3306"),
			"database": config.Env("DB_DATABASE", "microdemy"),
			"username": config.Env("DB_USERNAME", ""),
			"password": config.Env("DB_PASSWORD", ""),
		},
		"mongo": config.Env("MONGO_DB_URI", "mongodb://localhost:27017"),
		"redis": map[string]any{
			"host":     config.Env("REDIS_HOST", ""),
			"password": config.Env("REDIS_PASSWORD", ""),
			"port":     config.Env("REDIS_PORT", 6379),
			"database": config.Env("REDIS_DB", 0),
			"prefix":   config.Env("REDIS_PREFIX", 0),
		},
	})

	// kafka config
	config.Add("kafka", map[string]any{
		"brokers": config.Env("KAFKA_BROKER", "localhost:9092"),
	})
}
