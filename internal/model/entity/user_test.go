package entity_test

import (
	"goOrigin/internal/model/entity"
	"testing"

	"github.com/stretchr/testify/suite"
)

// UserInfoSuite :
type UserInfoSuite struct {
	suite.Suite
	token string
	user  entity.UserEntity
}

func (s *UserInfoSuite) SetupTest() {
	s.token = "eyJhbGciOiJub25lIn0.eyJpc3MiOiAiZGV4IiwgInN1YiI6ICIxMjM0NTY3ODkwIiwgIm5hbWUiOiAiaWFuIiwgImVtYWlsIjogImpvaG4uZG9lQGV4YW1wbGUuY29tIn0.signature"
	s.user = entity.UserEntity{}
}

// TestMarshal :
func (s *UserInfoSuite) TestConfig() {
	s.user.ParseToken(s.token)
	s.Equal("ian", s.user.Name)
}

// TestHttpClient :
func TestViperConfiguration(t *testing.T) {
	suite.Run(t, new(UserInfoSuite))
}
