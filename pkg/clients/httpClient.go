package clients

import (
	"github.com/go-resty/resty/v2"
	//"gopkg.in/resty.v1"

	"sync"
)

type Common struct {
	Token string
}

var once = &sync.Once{}

func GetHttpClient() *resty.Client {
	once.Do(func() {
		resty.New()
	})
	return resty.New()
}
