package handlers

import (
	"encoding/json"
	"github.com/labstack/echo"
	"gorest/db"
	"gorest/entity"
	"gorest/web"
	"gorest/web/models"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"strconv"
)

type UserHandler struct {
	Repository db.UserRepository
}

func NewUserHandler(e *echo.Echo, repo *db.UserRepository) {
	handler := UserHandler{
		Repository: *repo,
	}
	e.GET(web.Users, handler.GetUsers)
	e.GET(web.UsersId, handler.GetUsersId)
	e.GET(web.UsersSearch, handler.GetUsersSearch)
	e.POST(web.Users, handler.PostUsers)
	e.PATCH(web.UsersId, handler.PatchUsers)
	e.DELETE(web.UsersId, handler.DeleteUsers)
	e.GET(web.UsersCount, handler.GetUsersCount)
}

func (h *UserHandler) GetUsers(c echo.Context) error {
	p := c.QueryParam("isPaged")
	isPaged, _ := strconv.ParseBool(p)
	if isPaged {
		pageSizeQ := c.QueryParam("pageSize")
		pageNumberQ := c.QueryParam("pageNumber")
		orderQ := c.QueryParam("order")
		ascQ := c.QueryParam("asc")

		pageSize, err := strconv.Atoi(pageSizeQ)
		if err != nil {
			return c.JSON(http.StatusBadRequest, nil)
		}
		pageNumber, err := strconv.Atoi(pageNumberQ)
		if err != nil {
			return c.JSON(http.StatusBadRequest, nil)
		}
		if orderQ == "" {
			return c.JSON(http.StatusBadRequest, nil)
		}
		asc, err := strconv.ParseBool(ascQ)
		if err != nil {
			return c.JSON(http.StatusBadRequest, nil)
		}

		res, err := h.Repository.FindAllPagedAndSorted(int(pageNumber), int(pageSize), orderQ, asc)
		if err != nil {
			return c.JSON(http.StatusBadRequest, nil)
		}
		return c.JSON(http.StatusOK, res)
	}
	res, err := h.Repository.FindAll()
	if err != nil {
		return c.JSON(http.StatusNotFound, nil)
	}
	return c.JSON(http.StatusOK, res)
}

func (h *UserHandler) GetUsersId(c echo.Context) error {
	idP := c.Param("id")
	id, err := strconv.Atoi(idP)
	if err != nil || id < 0 {
		return c.JSON(http.StatusBadRequest, nil)
	}
	res, err := h.Repository.FindByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, nil)
	}
	return c.JSON(http.StatusOK, res)
}

func (h *UserHandler) GetUsersSearch(c echo.Context) error {
	username := c.QueryParam("username")
	if username == "" {
		return c.JSON(http.StatusBadRequest, nil)
	}
	res, err := h.Repository.FindByUsername(username)
	if err != nil {
		return c.JSON(http.StatusNotFound, nil)
	}
	return c.JSON(http.StatusOK, res)
}

func (h *UserHandler) PostUsers(c echo.Context) error {
	bodyBytes, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	var userRequest models.UserRequest
	err = json.Unmarshal(bodyBytes, &userRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	user := &entity.User{
		Username:    userRequest.Username,
		Password:    userRequest.Password,
		Gender:      userRequest.Gender,
		Role:        userRequest.Role,
		AvatarUrl:   userRequest.AvatarUrl,
		Description: userRequest.Description,
		Valid:       userRequest.Valid,
		Recipes:     userRequest.Recipes,
	}
	err = h.Repository.Create(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, &user)
}

func (h *UserHandler) PatchUsers(c echo.Context) error {
	idP := c.Param("id")
	id, err := strconv.Atoi(idP)
	if err != nil || id < 0 {
		return c.JSON(http.StatusBadRequest, nil)
	}

	bodyBytes, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	var userRequest models.UserRequest
	err = json.Unmarshal(bodyBytes, &userRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	user := &entity.User{
		Model:       gorm.Model{ID: uint(id)},
		Username:    userRequest.Username,
		Password:    userRequest.Password,
		Gender:      userRequest.Gender,
		Role:        userRequest.Role,
		AvatarUrl:   userRequest.AvatarUrl,
		Description: userRequest.Description,
		Valid:       userRequest.Valid,
		Recipes:     userRequest.Recipes,
	}

	err = h.Repository.Update(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	return c.JSON(http.StatusOK, nil)
}

func (h *UserHandler) DeleteUsers(c echo.Context) error {
	idP := c.Param("id")
	id, err := strconv.Atoi(idP)
	if err != nil || id < 0 {
		return c.JSON(http.StatusBadRequest, nil)
	}
	err = h.Repository.DeleteByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, nil)
	}
	return c.JSON(http.StatusNoContent, nil)
}

func (h *UserHandler) GetUsersCount(c echo.Context) error {
	res, err := h.Repository.Count()
	if err != nil {
		return c.JSON(http.StatusNotFound, nil)
	}
	return c.JSON(http.StatusOK, res)
}
