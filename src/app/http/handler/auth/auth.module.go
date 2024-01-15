package auth

import (
	"github.com/labstack/echo/v4"
)

type ModuleAuth struct{}

func (h *ModuleAuth) ApiV1(route *echo.Group) (*echo.Group, []any) {
	group := route.Group("api/v1/auth")

	return group, []any{
		//
	}
}
