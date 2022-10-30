package handler

import (
	"slack/internal/usecase"
)

type slackHandler struct {
	usecase usecase.UseCase
}

type SlackHandler interface {
	SaveAttendance(string, string) error
}

var _ SlackHandler = &slackHandler{}

func NewSlackHandler(u usecase.UseCase) *slackHandler {
	return &slackHandler{
		usecase: u,
	}
}

func (sh *slackHandler) SaveAttendance(message string, user string) error {
	if err := sh.usecase.SaveAttendance(message, user); err != nil {
		return err
	}

	return nil
}
