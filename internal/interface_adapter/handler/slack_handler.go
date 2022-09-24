package handler

import (
	"slack/internal/usecase"

	"github.com/labstack/echo/v4"
)

type slackHandler struct {
	usecase usecase.UseCase
}

type SlackHandler interface {
	SaveAttendance(echo.Context) error
}

var _ SlackHandler = &slackHandler{}

func NewSlackHandler(u usecase.UseCase) *slackHandler {
	return &slackHandler{
		usecase: u,
	}
}

func (sh *slackHandler) SaveAttendance(c echo.Context) error {
	// echo.Contextを使ったrequestの展開
	value := c.FormValue("sample")

	if err := sh.usecase.SaveAttendance(value); err != nil {
		return err
	}
	return nil
}
