package client

import (
	"fmt"
	"slack/internal/interface_adapter/gateway"
)

type gasClient struct{}

var _ gateway.GASClient = &gasClient{}

func NewGASClient() *gasClient {
	return &gasClient{}
}

func (gc *gasClient) Hoge(message string) error {
	fmt.Println(message)
	return nil
}
