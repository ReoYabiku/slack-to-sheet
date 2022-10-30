package gateway

import (
	"slack/internal/usecase"
)

type gasGateway struct {
	gasClient GASClient
}

type GASClient interface {
	SetValue(int, int, string) error
}

var _ usecase.GASGateway = &gasGateway{}

func NewGASGateway(gc GASClient) *gasGateway {
	return &gasGateway{
		gasClient: gc,
	}
}

func (gg *gasGateway) SaveCommentTime(reportType bool, user string) error {
	row := 5 //TODO： 日付から行番号を判断する
	column := getColumn(reportType, user)
	value := "20:45" // TODO: slack APIのrequestのtsをUNIX時間変換して取得する
	if err := gg.gasClient.SetValue(row, column, value); err != nil {
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

func getColumn(reportType bool, user string) int {
	// TODO: ユーザーごと、勤怠ごとに列を変える
	if reportType && user == "yabiku" {
		return 7
	} else {
		return 1
	}
}