package utils_test

import (
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
func (s *UtilsSuite) TestConfig() {
	//res := utils.ConvBson(map[string][]map[string]string{
	//	"$or": {
	//		{"severname": "test3"}, {"version": "ianLiuUpdate"},
	//	},
	//})
	//s.Equal(bson.M{
	//	"$or": bson.M{"severname": "test3", "version": "ianLiuUpdate"},
	//}, res)
}

// TestHttpClient :
func TestViperConfiguration(t *testing.T) {
	suite.Run(t, new(UtilsSuite))
}
