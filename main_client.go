package echargeSDK

import (
	"context"
	"fmt"

	"github.com/monaco-io/request"
)

func (s *Echarge) CreateClient(body RequestClient) (*ResponseClient, error) {
	response := ResponseClient{}
	url := string(s.environment + "v3/customers")
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

func (s *Echarge) ListClient(param string) (*ResponseClient, error) {
	response := ResponseClient{}
	url := fmt.Sprintf(string(s.environment+"v3/customers/cpfCnpj=%s"), param)
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

func (s *Echarge) UpdateClient(param string, body RequestClient) (*ResponseClient, error) {
	response := ResponseClient{}
	url := fmt.Sprintf(string(s.environment+"v3/customers/%s"), param)
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

func (s *Echarge) DeleteClient(param string) error {
	response := ResponseClient{}
	url := fmt.Sprintf(string(s.environment+"v3/customers/%s"), param)
	resp := request.
		NewWithContext(context.Background()).
		PUT(url).
		AddHeader(s.header).
		Send().
		Scan(&response)

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

func (s *Echarge) RestoreClient(param string) (*ResponseClient, error) {
	response := ResponseClient{}
	url := fmt.Sprintf(string(s.environment+"v3/customers/%s/restore"), param)
	resp := request.
		NewWithContext(context.Background()).
		POST(url).
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
