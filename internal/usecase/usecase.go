package usecase

import (
	"errors"
	"regexp"
)

type useCase struct {
	gasGateway GASGateway
}

type UseCase interface {
	SaveAttendance(string, string) error
	GetStartTime()
	GetFinishTime()
	GetTotalTime()
	CreateNewSheet()
}

var _ UseCase = &useCase{}

func NewUseCase(gg GASGateway) *useCase {
	return &useCase{
		gasGateway: gg,
	}
}

func (u *useCase) SaveAttendance(message string, user string) error {

	// 開始or終了はstring | intで判断形式に変える
	reportType := isStart(message)
	if err := u.gasGateway.SaveCommentTime(reportType, user); err != nil {
		return err
	}
	
	return errors.New("This message was not an attendance report.")
}

func (u *useCase) GetStartTime() {}

func (u *useCase) GetFinishTime() {}

func (u *useCase) GetTotalTime() {}

func (u *useCase) CreateNewSheet() {}

func isStart(message string) bool {
	r := regexp.MustCompile(`開始`)
	return r.MatchString(message)
}