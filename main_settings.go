package echargeSDK

type Environment string

const (
	Prodution    Environment = "https://api.asaas.com/"
	Homologation Environment = "https://api-sandbox.asaas.com/"
)

type Echarge struct {
	key         string
	environment Environment
	header      map[string]string
}

func (e *Echarge) SetKey(key string) *Echarge {
	e.key = key
	return e
}

func (e *Echarge) SetEnvironment(env Environment) *Echarge {
	e.environment = env
	return e
}

func (e *Echarge) SetHeader() *Echarge {
	e.header = map[string]string{
		"accept":       "application/json",
		"content-type": "application/json",
		"access_token": e.key,
	}
	return e
}
