package main

import (
	"fmt"
	"log"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type request struct {
	Challenge string 	`json:challenge`
}

func main() {
	e := echo.New()
	e.POST("/slack/challenge", func(c echo.Context) error {
		req := request{}
		if err := c.Bind(&req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		log.Println("got an access!!")
		log.Printf("challenge: %s\n", req.Challenge)
		return c.JSON(http.StatusOK, req)
	})
	e.GET("/", func(c echo.Context) error {
		log.Println("got an access!!")
		return c.String(http.StatusOK, "hoge")
	})
	e.Logger.Fatal(e.Start(":80"))
}
