package main

import (
	"log"
	"os"

	"github.com/chawintee/assessment/expense"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	expense.InitDB()
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	log.Fatal(e.Start(os.Getenv("PORT")))
}
