package clients

import (
	"github.com/go-resty/resty/v2"
	"sync"
)

type Common struct {
	Token string
}

var HttpClient *resty.Client

var once = &sync.Once{}

func GetHttpClient() *resty.Client {
	once.Do(func() {
		HttpClient = resty.New()
	})
	return HttpClient
}
