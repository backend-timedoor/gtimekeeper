package main

import (
	"gtimekeeper/bootstrap"
	"runtime"

	"github.com/backend-timedoor/gtimekeeper-framework/app"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	bootstrap.Boot()

	go func() {
		app.Job.Run()
		app.Server.Http().Run(app.Config.GetString("app.host"))
		// app.Server.Grpc().Run(app.Config.GetString("app.host"))

	}()

	select {}
}
