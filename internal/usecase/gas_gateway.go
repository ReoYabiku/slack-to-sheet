package usecase

type GASGateway interface {
	SaveCommentTime(bool, string) error
	SaveFinishTime() error
	GetStartTime()
	GetFinishTime()
	GetTotalTime()
	CreateNewSheet()
}
