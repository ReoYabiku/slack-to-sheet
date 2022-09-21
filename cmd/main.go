package main

import (
	"fmt"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.POST("/slack/challenge", func(c echo.Context) error {
		challenge := c.QueryParam("challenge")
		fmt.Printf("query: %s", challenge)
		return c.String(http.StatusOK, challenge)
	})
	e.Logger.Fatal(e.Start(":80"))
}
