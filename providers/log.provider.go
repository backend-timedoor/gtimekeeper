package providers

import (
	"fmt"
	"runtime/debug"
	"slices"

	"github.com/backend-timedoor/gtimekeeper-framework/app"
	"github.com/backend-timedoor/gtimekeeper-framework/base/log"
	"github.com/sirupsen/logrus"
)

type LogServiceProvider struct{}
type LogHook struct{}

func (p *LogServiceProvider) Boot() {
	app.Log.SetReportCaller(true)

	app.Log.AddHook(&LogHook{})
}

func (p *LogServiceProvider) Register() {
	app.Log = log.New()
}

func (h *LogHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *LogHook) Fire(e *logrus.Entry) error {
	environments := []string{"local", "development"}
	levels := []string{"error", "trace"}

	// err := fmt.Sprintf("%s:%d %s", e.Caller.File, e.Caller.Line, e.Level)
	// fmt.Println(err)

	if slices.Contains(environments, app.Config.GetString("app.env")) {
		if slices.Contains(levels, e.Level.String()) {
			fmt.Println(string(debug.Stack()))
		}
	}

	// 310fd47fefef35594d0ad1ab3e68c8fdFFFFNRAL

	return nil
}
