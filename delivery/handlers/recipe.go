package handlers

import (
	"encoding/json"
	"github.com/labstack/echo"
	"gorest/common"
	"gorest/delivery"
	"gorest/delivery/models"
	"gorest/entity"
	"gorest/service"
	"io/ioutil"
	"net/http"
	"strconv"
)

type RecipeHandler struct {
	Service service.RecipeService
}

func NewRecipeHandler(e *echo.Echo, service *service.RecipeService) {
	handler := RecipeHandler{
		Service: *service,
	}
	e.GET(delivery.Recipes, handler.GetRecipes)
	e.GET(delivery.RecipesId, handler.GetRecipesId)
	e.POST(delivery.RecipesSearch, handler.PostRecipesSearch) // by title, products and or tags
	e.POST(delivery.Recipes, handler.PostRecipes)
	e.POST(delivery.RecipesBatch, handler.PostRecipesBatch)
	e.PUT(delivery.RecipesId, handler.PutRecipeId)
	e.DELETE(delivery.RecipesId, handler.DeleteRecipesId)
	e.GET(delivery.RecipesCount, handler.GetRecipesCount)
}

func (h *RecipeHandler) GetRecipes(c echo.Context) error {
	p := c.QueryParam("isPaged")
	isPaged, _ := strconv.ParseBool(p)
	if !isPaged {
		res, err := h.Service.FindAll()
		if err != nil {
			return c.JSON(common.GetErrorResponse(err))
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
		return c.JSON(common.GetErrorResponse(err))
	}

	res, err := h.Service.
		FindAllPagedAndSorted(pageNumber, pageSize, orderQ, asc)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	return c.JSON(http.StatusOK, res)
}

func (h *RecipeHandler) GetRecipesId(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 0 {
		return c.JSON(http.StatusBadRequest, nil)
	}

	res, err := h.Service.FindByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, nil)
	}

	return c.JSON(http.StatusOK, res)
}

func (h *RecipeHandler) PostRecipesSearch(c echo.Context) error {
	bytes, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(common.GetErrorResponse(err))
	}

	var request models.RecipesSearchRequest
	err = json.Unmarshal(bytes, &request)
	if err != nil {
		return c.JSON(common.GetErrorResponse(err))
	}

	res, err := h.Service.Search(request)
	if err != nil {
		return c.JSON(common.GetErrorResponse(err))
	}
	return c.JSON(http.StatusOK, res)
}

func (h *RecipeHandler) PostRecipes(c echo.Context) error {
	bytes, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(common.GetErrorResponse(err))
	}

	var recipe entity.Recipe
	err = json.Unmarshal(bytes, &recipe)
	if err != nil {
		return c.JSON(common.GetErrorResponse(err))
	}

	err = h.Service.Create(&recipe)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, recipe)
}

func (h *RecipeHandler) PostRecipesBatch(c echo.Context) error {
	bytes, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(common.GetErrorResponse(err))
	}

	var recipes []entity.Recipe
	err = json.Unmarshal(bytes, &recipes)
	if err != nil {
		return c.JSON(common.GetErrorResponse(err))
	}

	err = h.Service.CreateBatch(&recipes)
	if err != nil {
		return c.JSON(common.GetErrorResponse(err))
	}

	return c.JSON(http.StatusOK, recipes)
}

func (h *RecipeHandler) PutRecipeId(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(common.GetErrorResponse(err))
	}

	bytes, err := ioutil.ReadAll(c.Request().Body)

	var recipe entity.Recipe
	err = json.Unmarshal(bytes, &recipe)
	if err != nil {
		return c.JSON(common.GetErrorResponse(err))
	}
	recipe.ID = uint(id)

	err = h.Service.Update(&recipe)
	if err != nil {
		return c.JSON(common.GetErrorResponse(err))
	}

	return c.JSON(http.StatusOK, &recipe)
}

func (h *RecipeHandler) DeleteRecipesId(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(common.GetErrorResponse(err))
	}

	res, err := h.Service.DeleteByID(uint(id))
	if err != nil {
		return c.JSON(common.GetErrorResponse(err))
	}

	return c.JSON(http.StatusOK, res)
}

func (h *RecipeHandler) GetRecipesCount(c echo.Context) error {
	r, err := h.Service.Count()
	if err != nil {
		return c.JSON(common.GetErrorResponse(err))
	}

	return c.JSON(http.StatusOK, r)
}
