package utils_test

import (
	"encoding/json"
	"fmt"
	"goOrigin/API/V1"
	"goOrigin/internal/dao/mysql"
	"testing"

	"github.com/stretchr/testify/suite"
)

// JsonSuit :
type JsonSuit struct {
	suite.Suite
}

func (s *JsonSuit) SetupTest() {

}

// TestMarshal :
func (s *JsonSuit) TestStruct() {
	var (
		res []byte
		err error
	)
	res, err = json.MarshalIndent(mysql.TRecord{}, "", "    ")
	s.NoError(err)
	fmt.Println(string(res))

	res, err = json.MarshalIndent(V1.CreateIanRecordRequest{}, "", "    ")
	s.NoError(err)
	fmt.Println(string(res))

}

// TestHttpClient :
func TestJsonConfiguration(t *testing.T) {
	suite.Run(t, new(JsonSuit))
}
