package entity_test

import (
	"fmt"
	"goOrigin/API/outter"
	"goOrigin/internal/model/entity"
	"goOrigin/pkg/utils"
	"testing"

	"github.com/stretchr/testify/suite"
)

// UserInfoSuite :
type UserInfoSuite struct {
	suite.Suite
	token  string
	user   entity.UserEntity
	cpUser entity.ZpaasUserEntity
}

func (s *UserInfoSuite) SetupTest() {
	s.token = "eyJhbGciOiJub25lIn0.eyJpc3MiOiAiZGV4IiwgInN1YiI6ICIxMjM0NTY3ODkwIiwgIm5hbWUiOiAiaWFuIiwgImVtYWlsIjogImpvaG4uZG9lQGV4YW1wbGUuY29tIn0.signature"
	s.user = entity.UserEntity{}
	s.cpUser = entity.ZpaasUserEntity{}
}

// TestMarshal :
func (s *UserInfoSuite) TestConfig() {
	fmt.Println(utils.ToJson(entity.KafkaLogEntity{}))
	s.user.ParseToken(s.token)
	s.Equal("ian", s.user.Name)
}

func (s *UserInfoSuite) TestSubjectAccessView() {
	s.user.ParseToken(s.token)
	s.Equal("ian", s.user.Name)
	var (
		req = outter.SubjectAccessViewReq{
			Url:          "localhost:8000",
			User:         "ian",
			Group:        "xx",
			Resource:     "pods",
			Verb:         "get",
			Namespace:    "default",
			ResourceName: "",
		}
	)
	res, _, _ := s.cpUser.SubjectReview(req)
	s.Equal(false, res.Status.Allowed)

	var (
		req2 = outter.SubjectAccessViewReq{
			Url:          "localhost:8000",
			User:         "ian",
			Group:        "xx",
			Resource:     "pods",
			Verb:         "get",
			Namespace:    "default",
			ResourceName: "",
		}
	)
	res2, _, _ := s.cpUser.SubjectReview(req2)
	s.Equal(true, res2.Status.Allowed)
}

// TestHttpClient :
func TestViperConfiguration(t *testing.T) {
	suite.Run(t, new(UserInfoSuite))
}
