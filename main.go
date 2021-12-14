package main

import (
	"chat-system/api"
	"chat-system/database"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db, err := database.InitDB()
	if err != nil {
		panic(err)
	}

	api.RegisterAPIHandler(db, e)

	e.Logger.Fatal(e.Start(":3000"))
}
