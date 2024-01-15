package user

import (
	"github.com/labstack/echo/v4"
	"gtimekeeper/src/modules/user"
	"net/http"
	"strconv"

	"github.com/backend-timedoor/gtimekeeper-framework/utils/helper"
	"github.com/backend-timedoor/gtimekeeper-framework/utils/paginate"
)

type HandlerV1User struct {
	UserService user.ServiceUserInterface
}

func (h *HandlerV1User) Boot(route *echo.Group) {
	h.UserService = &user.ServiceUser{}

	route.GET("/user", h.findAll)
	route.POST("/user", h.store)
	route.GET("/user/:id", h.find)
	route.PATCH("/user/:id", h.update)
	route.DELETE("/user/:id", h.delete)
}

func (h *HandlerV1User) findAll(c echo.Context) error {
	var request user.QueryRequest
	var response []user.Response

	if err := c.Bind(&request); err != nil {
		return helper.ErrorResponse(http.StatusBadRequest, helper.Resp{
			"message": err.Error(),
		})
	}

	users, pagination, err := h.UserService.FindAll(request)
	if err != nil {
		return helper.ErrorResponse(http.StatusInternalServerError, helper.Resp{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, &paginate.PaginationResponse{
		Data:       helper.Clone(&response, &users),
		Pagination: pagination,
	})
}

func (h *HandlerV1User) store(c echo.Context) error {
	var request user.CreateUserRequest
	var response user.Response

	if err := helper.Validate(c, &request); err != nil {
		return err
	}

	user, err := h.UserService.Store(request)
	if err != nil {
		return helper.ErrorResponse(http.StatusInternalServerError, helper.Resp{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, helper.Clone(&response, &user))
}

func (h *HandlerV1User) find(c echo.Context) error {
	var response user.Response

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return helper.ErrorResponse(http.StatusBadRequest, helper.Resp{
			"message": err.Error(),
		})
	}

	user, err := h.UserService.FindById(id)
	if err != nil {
		return helper.ErrorResponse(http.StatusInternalServerError, helper.Resp{
			"message": err.Error(),
		})
	}

	return helper.SuccessResponse(c, http.StatusOK, helper.Clone(&response, &user))
}

func (h *HandlerV1User) update(c echo.Context) error {
	var request user.UpdateUserRequest
	var response user.Response

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return helper.ErrorResponse(http.StatusBadRequest, helper.Resp{
			"message": err.Error(),
		})
	}

	if err := helper.Validate(c, &request); err != nil {
		return err
	}

	user, err := h.UserService.Update(request, id)
	if err != nil {
		return helper.ErrorResponse(http.StatusInternalServerError, helper.Resp{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, helper.Clone(&response, &user))
}

func (h *HandlerV1User) delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return helper.ErrorResponse(http.StatusBadRequest, helper.Resp{
			"message": err.Error(),
		})
	}

	err = h.UserService.Delete(id)
	if err != nil {
		return helper.ErrorResponse(http.StatusInternalServerError, helper.Resp{
			"message": err.Error(),
		})
	}
	return c.NoContent(http.StatusOK)
}
