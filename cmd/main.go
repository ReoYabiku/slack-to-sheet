package main

import (
	"log"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type request struct {
	Challenge string `json:"challenge"`
}

func slackAuth(c echo.Context) error {
	log.Println("got an access!!")
	req := request{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if req.Challenge != "" {
		log.Printf("challenge: %s\n", req.Challenge)
		return c.JSON(http.StatusOK, req)
	} else {
		return c.JSON(http.StatusOK, request{Challenge: "not challenge"})
	}
}

func main() {
	e := echo.New()
	e.POST("/event", slackAuth)
	e.GET("/", func(c echo.Context) error {
		log.Println("got an access!!")
		return c.String(http.StatusOK, "hoge")
	})
	e.Logger.Fatal(e.Start(":80"))
}
