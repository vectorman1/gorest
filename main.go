package main

import (
	"github.com/dystopia-systems/alaskalog"
	"github.com/labstack/echo"
	_db "gorest/db"
	_delivery "gorest/delivery/handlers"
	_service "gorest/service"
)

func main() {
	sql, err := _db.InitAndMigrate()
	if err != nil {
		alaskalog.Logger.Fatalf("Failed to init db: %v", err)
	}

	e := echo.New()
	e.HideBanner = true

	userRepo := _db.NewUserRepository(sql)
	recipeRepo := _db.NewRecipeRepository(sql)

	userService := _service.NewUserService(userRepo)
	recipeService := _service.NewRecipeService(recipeRepo, userRepo)

	_delivery.NewUserHandler(e, userService)
	_delivery.NewRecipeHandler(e, recipeService)

	alaskalog.Logger.Fatal(e.Start(":9001"))
}
