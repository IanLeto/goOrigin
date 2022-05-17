package model_test

import (
	"context"
	"github.com/stretchr/testify/suite"
	"goOrigin/agent/model"
	"testing"
)

// ModelSuite :
type ModelSuite struct {
	suite.Suite
	model.Task
}

func (s *ModelSuite) SetupTest() {

}

// TestMarshal :
func (s *ModelSuite) TestConfig() {
	a := &model.AsyncTask{
		ID:      "1001",
		Url:     "",
		Content: "",
		Timeout: 100,
		Ctx:     context.Background(),
		Status:  "",
	}
	_, err := a.Exec2(context.Background())
	s.NoError(err)
}
// TestMarshal :
func (s *ModelSuite) TestConfig2() {
	a := &model.AsyncTask{
		ID:      "1001",
		Url:     "",
		Content: "",
		Timeout: 100,
		Ctx:     context.Background(),
		Status:  "",
	}
	a.Reg()
	model.Query("1001")
}




// TestHttpClient :
func TestCmd(t *testing.T) {
	suite.Run(t, new(ModelSuite))
}


