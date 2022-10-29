package gateway

import (
	"slack/internal/usecase"
)

type gasGateway struct {
	gasClient GASClient
}

type GASClient interface {
	Hoge(string) error
}

var _ usecase.GASGateway = &gasGateway{}

func NewGASGateway(gc GASClient) *gasGateway {
	return &gasGateway{
		gasClient: gc,
	}
}

func (gg *gasGateway) SaveStartTime(message string) error {
	if err := gg.gasClient.Hoge(message); err != nil {
		return err
	}
	return nil
}

func (gg *gasGateway) SaveFinishTime() error {
	return nil
}
func (gg *gasGateway) GetStartTime()   {}
func (gg *gasGateway) GetFinishTime()  {}
func (gg *gasGateway) GetTotalTime()   {}
func (gg *gasGateway) CreateNewSheet() {}
