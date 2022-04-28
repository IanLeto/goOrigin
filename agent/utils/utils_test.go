package utils_test

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"goOrigin/pkg/utils"
	"os/exec"
	"testing"
)

// HttpClientSuite :
type HttpClientSuite struct {
	suite.Suite
}

func (s *HttpClientSuite) SetupTest() {
	fmt.Println(exec.Command("/bin/bash", utils.GetFilePath("template/test.sh")).CombinedOutput())
}

// TestMarshal :
func (s *HttpClientSuite) TestConfig() {

}

// TestHttpClient :
func TestCmd(t *testing.T) {
	suite.Run(t, new(HttpClientSuite))
}
