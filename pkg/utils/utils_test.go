package utils_test

import (
	"github.com/globalsign/mgo/bson"
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

// TestHttpClient :
func TestViperConfiguration(t *testing.T) {
	suite.Run(t, new(UtilsSuite))
}
