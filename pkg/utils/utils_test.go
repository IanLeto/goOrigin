package utils_test

import (
	"encoding/json"
	"fmt"
	"github.com/globalsign/mgo/bson"
	"goOrigin/API/V1"
	"goOrigin/pkg/utils"
	"testing"

	"github.com/stretchr/testify/suite"
)

// UtilsSuite :
type UtilsSuite struct {
	suite.Suite
}

func (s *UtilsSuite) SetupTest() {

}

// TestMarshal :
func (s *UtilsSuite) TestConvBson() {
	s.Equal(bson.M{"key": bson.M{"jk": "value"}}, utils.ConvBsonNoErr(map[string]map[string]interface{}{"key": {"jk": "value"}}))
}

// TestMarshal :
func (s *UtilsSuite) TestJson() {
	res, _ := json.MarshalIndent(V1.CreateIanRecordRequest{CreateIanRecordRequestInfo: &V1.CreateIanRecordRequestInfo{}}, "", "  ")
	fmt.Println(string(res))

}

// TestHttpClient :
func TestViperConfiguration(t *testing.T) {
	suite.Run(t, new(UtilsSuite))
}
