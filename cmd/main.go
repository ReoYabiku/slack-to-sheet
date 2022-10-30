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

type Event struct {
	Type string `json:"type"`
	User string `json:"user"`
	Message string `json:"message"`
}

type request struct {
	Challenge string `json:"challenge"`
	Event     Event  `json:"event"`
}

func main() {
	gc := client.NewGASClient()
	gg := gateway.NewGASGateway(gc)
	u := usecase.NewUseCase(gg)
	sh := handler.NewSlackHandler(u)

	e := echo.New()
	// mainが入力を見て判断する必要はない
	// TODO: handlerでrequestによって呼び出す関数を変える
	// TODO: usecaseにReturnChallenge（仮）SaveAttendance（実装済み）を置く
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
			return sh.SaveAttendance(req.Event.Message, req.Event.User)
		}
	})
	e.GET("/", func(c echo.Context) error {
		log.Println("got an access!!")
		return c.String(http.StatusOK, "hoge")
	})
	e.Logger.Fatal(e.Start(":80"))
}
