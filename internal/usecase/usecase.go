package usecase

import "errors"

type useCase struct {
	gasGateway GASGateway
}

type UseCase interface {
	SaveAttendance(string) error
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

func (u *useCase) SaveAttendance(message string) error {
	if isStart(message) {
		if err := u.gasGateway.SaveStartTime(message); err != nil {
			return err
		}
	}
	if isFinish(message) {
		if err := u.gasGateway.SaveFinishTime(); err != nil {
			return err
		}
	}
	return errors.New("This message was not an attendance report.")
}

func (u *useCase) GetStartTime() {}

func (u *useCase) GetFinishTime() {}

func (u *useCase) GetTotalTime() {}

func (u *useCase) CreateNewSheet() {}

func isStart(message string) bool {
	return true
}

func isFinish(message string) bool {
	return true
}
