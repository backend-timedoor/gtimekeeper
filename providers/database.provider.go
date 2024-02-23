package providers

import (
	"fmt"
	"time"

	"github.com/backend-timedoor/gtimekeeper-framework/app"
	"github.com/backend-timedoor/gtimekeeper-framework/base/database"
	"github.com/backend-timedoor/gtimekeeper-framework/base/database/drivers"
	"github.com/backend-timedoor/gtimekeeper-framework/base/database/redis"
	"gorm.io/gorm"
)

type DatabaseServiceProvider struct{}

func (p *DatabaseServiceProvider) Boot() {}

func (p *DatabaseServiceProvider) Register() {
	app.DB = database.New(&database.Config{
		Driver: &drivers.PgsqlDriver{
			Host:     app.Config.GetString("DB_HOST", "localhost"),
			Port:     app.Config.GetInt("DB_PORT", 5432),
			Database: app.Config.GetString("DB_DATABASE", "gtimekeeper"),
			Username: app.Config.GetString("DB_USERNAME", "root"),
			Password: app.Config.GetString("DB_PASSWORD", ""),
		},
		GormConfig: &gorm.Config{
			NowFunc: func() time.Time { //you need to add this for postgres
				return time.Now().In(time.Local)
			},
		},
		Mongo: app.Config.GetString("MONGO_DB_URI", "mongodb://localhost:27017"),
		Redis: &redis.Config{
			Addr:     fmt.Sprintf("%s:%s", app.Config.Env("REDIS_HOST"), app.Config.Env("REDIS_PORT")),
			Password: app.Config.GetString("REDIS_PASSWORD", ""),
			DB:       app.Config.GetInt("REDIS_DB", 0),
			Prefix:   app.Config.GetString("REDIS_PREFIX", "gtimekeeper"),
		},
	})
}
