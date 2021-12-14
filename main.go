package main

import (
	"chat-system/database"
	"fmt"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db, err := database.InitDB()
	if err != nil {
		panic(err)
	}

	fmt.Println(db.Error)

	e.Logger.Fatal(e.Start(":3000"))
}
