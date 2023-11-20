package clients_test

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"goOrigin/internal/dao/elastic"
	"goOrigin/pkg/utils"
	"net/http"
	"net/http/httputil"
	"testing"
)

func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	utils.NoError(err)
	fmt.Println(string(dump))
	w.Write([]byte("hello world"))

}

type HttpsSuite struct {
	suite.Suite
	conn *elastic.EsV2Conn
}

func (s *HttpsSuite) SetupTest() {
	http.HandleFunc("/", handler)
	http.ListenAndServeTLS(":8080", "/home/ian/workdir/cc/dev/server.crt",
		"/home/ian/workdir/cc/dev/server.key", nil)
}

// TestMarshal :
func (s *HttpsSuite) TestConfig() {

}

// TestHttpClient :
func TestHttpsConfiguration(t *testing.T) {
	suite.Run(t, new(HttpsSuite))
}
