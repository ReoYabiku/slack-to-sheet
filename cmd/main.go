package main

import (
	"fmt"
	"log"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.POST("/slack/challenge", func(c echo.Context) error {
		challenge := c.FormValue("challenge")
		fmt.Printf("query: %s", challenge)
		log.Println("got an access!!")
		return c.String(http.StatusOK, challenge)
	})
	e.GET("/", func(c echo.Context) error {
		log.Println("got an access!!")
		return c.String(http.StatusOK, "hoge")
	})
	e.Logger.Fatal(e.Start(":80"))
}
