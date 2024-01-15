package main

import (
	"github.com/backend-timedoor/gtimekeeper-framework/app"
	"gtimekeeper/bootstrap"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	bootstrap.Boot()

	go func() {
		// app.Schedule.Run()
		// app.Queue.Run()
		app.Server.Http().Run(app.Config.GetString("app.host"))
		// go app.Server.Grpc().Run(app.Config.GetString("app.host"))

	}()
	//
	select {}
}
