package main

import (
	"log"
	"net/http"
	"slack/internal/interface_adapter/client"
	"slack/internal/interface_adapter/gateway"
	"slack/internal/interface_adapter/handler"
	"slack/internal/usecase"

	echo "github.com/labstack/echo/v4"
)

type request struct {
	Challenge string `json:"challenge"`
}

func main() {
	gc := client.NewGASClient()
	gg := gateway.NewGASGateway(gc)
	u := usecase.NewUseCase(gg)
	sh := handler.NewSlackHandler(u)

	e := echo.New()
	e.POST("/event", func(c echo.Context) error {
		log.Println("got an access!!")
		req := request{}
		if err := c.Bind(&req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if req.Challenge != "" {
			log.Printf("challenge: %s\n", req.Challenge)
			return c.JSON(http.StatusOK, req)
		} else {
			sh.SaveAttendance(c)
			return nil
		}
	})
	e.GET("/", func(c echo.Context) error {
		log.Println("got an access!!")
		return c.String(http.StatusOK, "hoge")
	})
	e.Logger.Fatal(e.Start(":80"))
}
