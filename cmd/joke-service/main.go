package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const MAIN_API = "https://v2.jokeapi.dev/joke/"

func getJoke(c echo.Context) error {
	response, err := http.Get(MAIN_API + "any")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, response)
}

func main() {

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/joke", getJoke)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
