package main

import (
	"fmt"
	"math"
	"signailai/initializers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func init() {
	initializers.OpenConnection()
	initializers.OpenLocalConnection()
	initializers.CreateSchema()

	if initializers.NeedSeed() {
		fmt.Println("Seeding db")
		SeedLocalDB()
	} else {
		fmt.Println("Data is already seeded skipping...")
	}

	//no longer needed
	initializers.DB.Close()
	fmt.Println(getGridDimensions())
	min_x, min_y, max_x, max_y := getGridDimensions()
	x := math.Abs(float64(max_x) - float64(min_x))
	y := math.Abs(float64(max_y) - float64(min_y))
	fmt.Println("x: ", x, " y: ", y)

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
