package main

import (
	_ "github.com/lib/pq"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"openapi/internal/hello"
	"openapi/internal/stockitem"
)

// main is the entry point of the program.
//
// No parameters.
// No return type.
func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello.GetHello)
	e.POST("/stock/items", stockitem.PostStockItem)

	e.Logger.Fatal(e.Start(":3000"))
}
