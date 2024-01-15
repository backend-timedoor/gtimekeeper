# GTimeKeeper
This is initiation golang project starter.


## Introduction

## Dependencies Documentation (Please read these dependencies docs)
- [Go](https://go.dev/dl/) >= 1.21.4 
- ORM [GORM](https://gorm.io/docs/) = v1.25.5
- Cache Management [Redis](https://redis.io/docs/connect/clients/go/) (extending sequelize function with cache feature)
- Database Migration [Golang Migrate](https://github.com/golang-migrate/migrate)
---

# How To Install

- Open your project folder
- run ```go mod tidy ```
- run ```cp env.example .env```
- adjust `.env` file with you local machine
- how to run the project
  ```sh
  go run .
  ```
  ```sh
  go run main.go
  ```
  ```sh
  make run
  ```

## Architecture
Modular architecture, often referred to as modular design or modularity, is an architectural approach that involves breaking down a system into smaller, self-contained, and interchangeable modules. These modules encapsulate specific functionality, and the interactions between them are well-defined and limited. Each module can be developed, tested, and maintained independently, providing a more scalable, maintainable, and understandable structure for a software system.

## Provider
The provider is main gateway to register features on the app, you can make your own provider and register it on bootstrap/app.go, you add the config on this provider, provider has 2 main function `Boot()` and `Register()`, `Boot()` will be execute first and `Register()` after. default provider for the app.
 - app.provider.go
        
    you can add any setting for the app from this provider, for example we can setting default timezone for the app.
 - config.provider.go

    config the app from here, like db, redis, app name,  etc.
 - console.provider.go

    we can make our own console, and register it in this provider.
 - queue.provider.go

    create queue for executing data and you can register the queue on this provider.
 - scheduler.provider.go

    sometimes we need process run on the background in the specific time, you can make it and, register it on this provider
 - server.provider.go

    this is controller for register, you want to use which protocol on your apps, http or grpc
## Feature
### Server
#### HTTP
Register you HTTP module.
you can start the http server with this script in `main.go`.
```go
go func() {
    app.Server.Http().Run(app.Config.GetString("app.host"));
}()
```
#### GRPC
you can start the grpc server with this script in `main.go`.
```go
go func() {
    go app.Server.Grpc().Run(app.Config.GetString("app.host"));
}()
```
#### Custome Validation
You can add custome validation rule by implement the interface of `Validation`
```go
type Validation interface {
	Signature() string
	Handle(validator.FieldLevel) bool
}
```

and register it on `server.provider.go`

### Database
#### ORM
for ORM we use GORM, for the detail GORM Docs you can see it [here](https://gorm.io/docs/).
#### Migration
we use [Golang Migrate](https://github.com/golang-migrate/migrate) for our base migration, but we extend that for simplify reason. migration file will be located in `database/migrations`.
- create

    to make new file migration you just can run ```go run . gtime migrate:create --name=user ```
- up

    for run migration you just need to run
    ``` go run . gtime migrate:up ```
- down

    for undo the migration run.
    ``` go run . gtime migrate:down ```
    you can add `step=4` for undo the migration
#### MongoDB
set env for mongo first to use MongoDB, for mongo instance you can access `app.DB.Mongo()` 
and other function on go [Mongo](https://www.mongodb.com/docs/drivers/go/current/quick-start/)
here the example how to use it:
- Insert
```go
collection := app.DB.Mongo.Database("catfact").Collection("facts")
_, err := collection.InsertOne(context.TODO(), map[string]any{
    "breed": "Persian",
    "fact":  "Cats are the most popular pet in the United States: There are 88 million pet cats and 74 million dogs.",
})

if err != nil {
    app.Log.Fatal(err)
}
```
- Read
```go
collection := app.DB.Mongo.Database("catfact").Collection("facts")

query := bson.M{}
cursor, err := collection.Find(context.TODO(), query)
if err != nil {
    app.Log.Fatal(err)
}

var results []bson.M
if err = cursor.All(context.TODO(), &results); err != nil {
    app.Log.Fatal(err)
}
```

#### Cache
you can use this function to manage cache on the system 
```go
type Cache interface {
	Push(string, any) error
	Retrieve(string) []string
	Remove(string, int)
	Pop(string) []string
	Get(string, any) any
	Has(string) bool
	Set(string, any, time.Duration) error
	Pull(string, any) any
	Add(string, any, time.Duration) bool
	Remember(string, time.Duration, func() any) (any, error)
	RememberForever(string, func() any) (any, error)
	Forever(string, any) bool
	Forget(string) bool
	Flush() bool
}
```
example code 

```go
app.Cache.Add("test")

app.Cache.Add("test_add", "update", time.Minute * 300);
app.Cache.Push("key", "value here");
```
### Scheduler
you can add scheduler file on `src/cmd/schedules` and register the file in schedule provider.
```go
func (log *ScheduleServiceProvider) Register() {
	app.Schedule = schedule.BootSchedule([]contracts.ScheduleEvent{
		// register you schedule here
	})
}
```
schedule example 
```go
import (
	"fmt"
	"time"

	schedule "github.com/backend-timedoor/gtimekeeper-framework/base/scheduler"
)

type ExampleSchedule struct{}	

func (s *ExampleSchedule) Signature() string {
	return "example:schedule"
}

func (s *ExampleSchedule) Schedule() string {
	return schedule.EveryMinute()
}


func (s *ExampleSchedule) Handle() {
	fmt.Println("Cron job executed at:", time.Now())
}
```

and for run the schedule call this code on `main.go`
```go
go func() {
    app.Schedule.Run()
}()
```
### Queue
you should create a new queue with your own name, then use it to send message or receive message, you can make your own queue file at `src/cmd/jobs`, job file example 
```go
import (
	"context"
	"fmt"
	"time"

	"github.com/hibiken/asynq"
)

type ExampleJob struct{}

func (m *ExampleJob) Signature() string {
	return "example:job"
}

func (m *ExampleJob) Options() []asynq.Option {
	return []asynq.Option{
		asynq.ProcessIn(5 * time.Second),
	}
}


func (j *ExampleJob) Handle(ctx context.Context, t *asynq.Task) error {
	// data := json.Unmarshal(t.Payload())
	fmt.Println("job example is run")

	return nil
}
```
you can put setting the queue on the options function for example you want to delay the queue.
To add a task to the queue, simply create an instance of your job type and pass it, and call this code to add the task
```go
app.Queue.Job(&jobs.ExampleJob{}, map[string]any{
    "ID":       1,
    "Username": "john_doe",
    "Email":    "john.doe@example.com",
    "Age":      25,
})
```
### Console
you can make your own shell command to exec something you need to run just in time or it needed. you just need to make new console file on the `src/cmd/console`, example code:
```go
import (
	"fmt"

	"github.com/urfave/cli/v2"
)

type ExampleCommand struct{}

func (m *ExampleCommand) Signature() string {
	return "example"
}

func (m *ExampleCommand) Flags() []cli.Flag {
	return []cli.Flag{}
}

func (m *ExampleCommand) Handle(c *cli.Context) (err error) {
	fmt.Println("Hello World From Example Command!")

	return nil
}
```

and register it on `console.provider.go`
```go
func (p *ConsoleServiceProvider) Register() {
	console.BootConsole([]contracts.Commands{
		&cmd.ExampleCommand{},
		// new comment here
	})
}
```

and console ready to fire, you can call your command with this prefix shell
```go
go run . gtime your_command_signature
```
### Email
gtime framework already support for email send, you can easily send email, first you need to create instance email 
on `src/_common/mail`, example code:
```go
package mail

import (
	envelop "github.com/backend-timedoor/gtimekeeper-framework/utils/app/email"
	"github.com/jordan-wright/email"
)

type ExampleMail struct {
	SendTo      envelop.SendTo
	Attachments []*email.Attachment
}

func (m *ExampleMail) From() string {
	return "Edwin Diradinata <edwindiradinata@gmail.com>"
}

func (m *ExampleMail) Content(data any) envelop.Content {
	return envelop.Content{
		Subject: "Awesome Subject {Data Here}",
		ReplyTo: []string{"edwindiradinata@gmail.com"},
		Text:    []byte("Text Body is, of course, supported!"),
		HTML:    []byte("<h1>Fancy HTML is supported, too!</h1>"),
	}
}
```
if you want to use template for the email you can add this code to your email instance
```go
func (m *ExampleMail) View() string {
	return "example.html"
}
```
this view instance automatically pointing to folder `src/view/mail` and you can add your template there.
for the detail you can see the docs [here](https://github.com/jordan-wright/email)


