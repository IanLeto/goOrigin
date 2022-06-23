package model_test

import (
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type JobSuite struct {
	suite.Suite
}

func (s *JobSuite) SetupTest() {
}

func (s *JobSuite) TestSync() {
	_, err := os.OpenFile("test.sh", os.O_WRONLY|os.O_CREATE, 0666)
	s.NoError(err)

}

func TestJob(t *testing.T) {
	suite.Run(t, new(JobSuite))
}
