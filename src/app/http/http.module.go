package http

import (
	"github.com/labstack/echo/v4"
	"gtimekeeper/src/app/http/handler/auth"
	"gtimekeeper/src/app/http/handler/cms"
)

type ModuleHttp struct{}

func (h *ModuleHttp) ApiV1(route *echo.Group) (*echo.Group, []any) {
	group := route.Group("/api/v1")

	return group, []any{
		&cms.ModuleCms{},
		&auth.ModuleAuth{},
		// other modules
	}
}
