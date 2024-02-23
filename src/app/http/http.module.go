package http

import (
	"context"
	"encoding/json"
	"gtimekeeper/src/app/http/handler/auth"
	"gtimekeeper/src/app/http/handler/cms"
	"net/http"

	pkgKafka "gtimekeeper/src/_common/pkg/kafka"

	"github.com/backend-timedoor/gtimekeeper-framework/app"
	"github.com/backend-timedoor/gtimekeeper-framework/utils/helper"
	"github.com/labstack/echo/v4"
	"github.com/segmentio/kafka-go"
)

type ModuleHttp struct{}

func (h *ModuleHttp) ApiV1(route *echo.Group) (*echo.Group, []any) {
	group := route.Group("/api/v1")

	group.GET("/ping", func(c echo.Context) error {
		data := pkgKafka.ExampleStruct{
			ID:   1,
			Name: "test",
		}

		d, err := json.Marshal(data)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.Resp{
				"message": "error marshal data " + err.Error(),
			})
		}
		app.Kafka.Produce(context.Background(), kafka.Message{
			Topic: string(pkgKafka.ExampleTopic),
			Value: []byte(d),
		})

		return c.JSON(http.StatusOK, helper.Resp{
			"message": "pong",
		})
	})

	return group, []any{
		&cms.ModuleCms{},
		&auth.ModuleAuth{},
		// other modules
	}
}
