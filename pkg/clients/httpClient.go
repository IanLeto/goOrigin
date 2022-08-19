package clients

import "gopkg.in/resty.v1"

type Common struct {
	Token string
}
type OriginHttpClient struct {
}

func NewHttpClient() {
	resty.New()
}
