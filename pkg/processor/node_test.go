package processor_test

import (
	"github.com/stretchr/testify/suite"
	"goOrigin/pkg/cron"
	"testing"
)

// RedisSuite :
type NodeTest struct {
	suite.Suite
	jobs *cron.TaskManager
}

func (s *NodeTest) SetupTest() {
}

// TestMarshal :
func (s *NodeTest) TestConfig() {

}

// TestHttpClient :
func TestNodeTestConfiguration(t *testing.T) {
	suite.Run(t, new(NodeTest))
}
