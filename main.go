package main

import (
	"github.com/dystopia-systems/alaskalog"
	"github.com/labstack/echo"
	"gorest/db"
	_httpDelivery "gorest/web/handlers"
)

func main() {
	sql, err := db.InitAndMigrate()
	if err != nil {
		alaskalog.Logger.Fatalf("Failed to init db: %v", err)
	}

	e := echo.New()
	e.HideBanner = true

	userRepo := db.NewUserRepository(sql)

	_httpDelivery.NewUserHandler(e, userRepo)

	alaskalog.Logger.Fatal(e.Start(":9001"))
}
