package usecase

type GASGateway interface {
	SaveStartTime(string) error
	SaveFinishTime() error
	GetStartTime()
	GetFinishTime()
	GetTotalTime()
	CreateNewSheet()
}
