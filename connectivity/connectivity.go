package connectivity

type Session struct {
	Method string
	Header map[string]string
	Body   interface{}
}

func New() *Session {
	return &Session{}
}
