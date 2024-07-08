package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var events []Event

func main() {

	loadConfig()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/x.svg", func(c echo.Context) error {
		fileContent, err := os.ReadFile("template.svg")
		if err != nil {
			return err
		}

		template := string(fileContent)

		template = ProcessCalendar(template)
		template = ProcessWeather(template)
		template = ProcessTodo(template)

		return c.Blob(http.StatusOK, "image/svg+xml", []byte(template))
	})

	bindIP := fmt.Sprintf("%s:%d", Config.Server.IP, Config.Server.Port)

	log.Printf("Starting server on %s", bindIP)
	e.Logger.Fatal(e.Start(bindIP))

}
