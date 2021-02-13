package handlers

import (
	"encoding/json"
	"github.com/labstack/echo"
	"gorest/common"
	"gorest/delivery"
	"gorest/entity"
	"gorest/service"
	"io/ioutil"
	"net/http"
	"strconv"
)

type UserHandler struct {
	Service service.UserService
}

func NewUserHandler(e *echo.Echo, service *service.UserService) {
	handler := UserHandler{
		Service: *service,
	}
	e.GET(delivery.Users, handler.GetUsers)
	e.GET(delivery.UsersId, handler.GetUsersId)
	e.GET(delivery.UsersSearch, handler.GetUsersSearch)
	e.POST(delivery.Users, handler.PostUsers)
	e.PATCH(delivery.UsersId, handler.PatchUsers)
	e.DELETE(delivery.UsersId, handler.DeleteUsers)
	e.GET(delivery.UsersCount, handler.GetUsersCount)
}

func (h *UserHandler) GetUsers(c echo.Context) error {
	p := c.QueryParam("isPaged")
	isPaged, err := strconv.ParseBool(p)
	if err != nil || !isPaged {
		res, e := h.Service.FindAll()
		if e != nil {
			return c.JSON(common.GetErrorResponse(e))
		}

		return c.JSON(http.StatusOK, res)
	}

	pageSizeQ := c.QueryParam("pageSize")
	pageNumberQ := c.QueryParam("pageNumber")
	orderQ := c.QueryParam("order")
	ascQ := c.QueryParam("asc")

	pageSize, err := strconv.Atoi(pageSizeQ)
	if err != nil {
		return c.JSON(common.GetErrorResponse(common.InvalidModelError))
	}
	pageNumber, err := strconv.Atoi(pageNumberQ)
	if err != nil {
		return c.JSON(common.GetErrorResponse(common.InvalidModelError))
	}
	if orderQ == "" {
		return c.JSON(common.GetErrorResponse(common.InvalidModelError))
	}
	asc, err := strconv.ParseBool(ascQ)
	if err != nil {
		return c.JSON(common.GetErrorResponse(common.InvalidModelError))
	}

	res, err := h.Service.FindAllPagedAndSorted(pageNumber, pageSize, orderQ, asc)
	if err != nil {
		return c.JSON(common.GetErrorResponse(err))
	}
	if len(res) == 0 {
		return c.JSON(common.GetErrorResponse(common.EntityNotFoundError))
	}

	return c.JSON(http.StatusOK, res)
}

func (h *UserHandler) GetUsersId(c echo.Context) error {
	idP := c.Param("id")
	id, err := strconv.Atoi(idP)
	if err != nil || id < 0 {
		return c.JSON(common.GetErrorResponse(common.InvalidModelError))
	}

	res, err := h.Service.FindByID(uint(id))
	if err != nil {
		return c.JSON(common.GetErrorResponse(err))
	}

	return c.JSON(http.StatusOK, res)
}

func (h *UserHandler) GetUsersSearch(c echo.Context) error {
	username := c.QueryParam("username")
	if username == "" {
		return c.JSON(common.GetErrorResponse(common.InvalidModelError))
	}

	res, err := h.Service.FindByUsername(username)
	if err != nil {
		return c.JSON(common.GetErrorResponse(err))
	}

	return c.JSON(http.StatusOK, res)
}

func (h *UserHandler) GetUsersCount(c echo.Context) error {
	res, err := h.Service.Count()
	if err != nil {
		return c.JSON(common.GetErrorResponse(err))
	}

	return c.JSON(http.StatusOK, res)
}

func (h *UserHandler) PostUsers(c echo.Context) error {
	bodyBytes, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(common.GetErrorResponse(err))
	}

	var u entity.User
	err = json.Unmarshal(bodyBytes, &u)
	if err != nil {
		return c.JSON(common.GetErrorResponse(common.InvalidModelError))
	}

	err = h.Service.Create(&u)
	if err != nil {
		return c.JSON(common.GetErrorResponse(err))
	}

	return c.JSON(http.StatusCreated, &u)
}

func (h *UserHandler) PatchUsers(c echo.Context) error {
	idP := c.Param("id")
	id, err := strconv.Atoi(idP)
	if err != nil {
		return c.JSON(common.GetErrorResponse(common.InvalidModelError))
	}

	bodyBytes, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(common.GetErrorResponse(err))
	}

	var user entity.User
	err = json.Unmarshal(bodyBytes, &user)
	if err != nil {
		return c.JSON(common.GetErrorResponse(common.InvalidModelError))
	}

	user.ID = uint(id)
	err = h.Service.Update(&user)
	if err != nil {
		return c.JSON(common.GetErrorResponse(err))
	}

	return c.JSON(http.StatusOK, &user)
}

func (h *UserHandler) DeleteUsers(c echo.Context) error {
	idP := c.Param("id")
	id, err := strconv.Atoi(idP)
	if err != nil || id < 0 {
		return c.JSON(common.GetErrorResponse(common.InvalidModelError))
	}

	r, err := h.Service.DeleteByID(uint(id))
	if err != nil {
		return c.JSON(common.GetErrorResponse(err))
	}

	return c.JSON(http.StatusNoContent, r)
}
