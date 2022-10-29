package handler

import (
	"slack/internal/usecase"

	"github.com/labstack/echo/v4"
)

type Event struct {
	Type string `json:"type"`
	User string `json:"user"`
	Text string `json:"text"`
}

type Message struct {	
	Event Event `json:"event"`
}

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

	m := Message{}
	if err := c.Bind(m); if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	fmt.Printf("text; %s", m.Event.Text)

	if err := sh.usecase.SaveAttendance(value); err != nil {
		return err
	}
	return nil
}
