package echargeSDK

import (
	"context"

	"github.com/monaco-io/request"
)

func (s *Echarge) CreateToken(body RequestCreditCard) (*ResponseCreditCard, error) {
	response := ResponseCreditCard{}
	url := string(s.environment + "v3/creditCard/tokenizeCreditCard")
	resp := request.
		NewWithContext(context.Background()).
		POST(url).
		AddHeader(s.header).
		AddJSON(body).
		Send().
		Scan(&response)

	defer resp.Close()

	switch resp.Response().StatusCode {
	case 400:
		return nil, resp.Error()
	case 401:
		return nil, resp.Error()
	case 403:
		return nil, resp.Error()
	case 404:
		return nil, resp.Error()
	case 429:
		return nil, resp.Error()
	default:
		return &response, nil
	}
}
