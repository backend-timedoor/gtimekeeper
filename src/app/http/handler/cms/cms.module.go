package cms

import (
	"github.com/labstack/echo/v4"
	"gtimekeeper/src/app/http/handler/cms/user"
)

type ModuleCms struct{}

func (h *ModuleCms) CmsApiV1(route *echo.Group) (*echo.Group, []any) {
	group := route.Group("/cms")

	return group, []any{
		&user.HandlerV1User{},
	}
}
