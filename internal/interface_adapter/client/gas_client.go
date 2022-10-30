package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"slack/internal/interface_adapter/gateway"
)

type gasClient struct{}

type RequestBody struct {
	Row    int    `json:"row"`
	Column int    `json:"column"`
	Value  string `json:"value"`
}

var _ gateway.GASClient = &gasClient{}

func NewGASClient() *gasClient {
	return &gasClient{}
}

func (gc *gasClient) SetValue(row int, column int, value string) error {
	requestBody := &RequestBody{
		Row:    row,
		Column: column,
		Value:  value,
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	endpoint := "https://script.google.com/macros/s/AKfycbyf3TLco4AZ-_lGOo1ra57u7x0a5PUxJIgeq6eTvZS-paMImhN12k50Z7JFAIleuapK/exec?task=select"
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonString))
	if err != nil {
		fmt.Printf("error: %v", err)
	}


	req.Header.Set("Content-Type", "application-json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	fmt.Printf("%v", string(byteArray))
	// TODO: 適切なエラーを返す

	return nil
}
