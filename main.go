package main

import (
	"chat-system/api"
	"chat-system/database"
	"chat-system/service"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db, err := database.InitDB()
	if err != nil {
		panic(err)
	}

	info := service.NewInfoService(db)
	api.RegisterAPIHandler(info, e)

	e.Logger.Fatal(e.Start(":3000"))
}
