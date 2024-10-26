package main

import (
	"signailai/initializers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func init() {
	initializers.OpenConnection()
}

func main() {

	defer initializers.DB.Close()

	e := echo.New()

	e.GET("/matavimai", func(c echo.Context) error {
		data, _ := GetMeasurements()

		return c.JSON(200, data)
	})

	e.GET("/stiprumai", func(c echo.Context) error {
		data, _ := GetStrengths()

		return c.JSON(200, data)
	})

	e.GET("/naudotojai", func(c echo.Context) error {
		data, _ := GetUsers()

		return c.JSON(200, data)
	})

	e.Logger.Fatal(e.Start(":3000"))
}
