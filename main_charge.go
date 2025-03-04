package echargeSDK

import (
	"context"
	"fmt"

	"github.com/monaco-io/request"
)

func (s *Echarge) CreateCharge(body RequestCharge) (*ResponseCharge, error) {
	response := ResponseCharge{}
	url := string(s.environment + "v3/payments")
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

func (s *Echarge) UpdateCharge(param string, body RequestCharge) (*ResponseCharge, error) {
	response := ResponseCharge{}
	url := fmt.Sprintf(string(s.environment+"v3/payments/%s"), param)
	resp := request.
		NewWithContext(context.Background()).
		PUT(url).
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

func (s *Echarge) SearchCharge(param string) (*ResponseCharge, error) {
	response := ResponseCharge{}
	url := fmt.Sprintf(string(s.environment+"v3/payments/%s"), param)
	resp := request.
		NewWithContext(context.Background()).
		GET(url).
		AddHeader(s.header).
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

func (s *Echarge) DeleteCharge(param string) error {
	url := fmt.Sprintf(string(s.environment+"v3/payments/%s"), param)
	resp := request.
		NewWithContext(context.Background()).
		DELETE(url).
		AddHeader(s.header).
		Send()

	defer resp.Close()

	switch resp.Response().StatusCode {
	case 400:
		return resp.Error()
	case 401:
		return resp.Error()
	case 403:
		return resp.Error()
	case 404:
		return resp.Error()
	case 429:
		return resp.Error()
	default:
		return nil
	}
}
