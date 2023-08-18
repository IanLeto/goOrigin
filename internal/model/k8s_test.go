package model

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"goOrigin/API/V1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

// K8sParamsSuite :
type K8sParamsSuite struct {
	suite.Suite
}

func (s *K8sParamsSuite) SetupTest() {

}

// TestMarshal :
func (s *K8sParamsSuite) TestConvBson() {
	x2 := v1.ObjectMeta{
		Name:                       "",
		GenerateName:               "",
		Namespace:                  "",
		SelfLink:                   "",
		UID:                        "",
		ResourceVersion:            "",
		Generation:                 0,
		CreationTimestamp:          v1.Time{},
		DeletionTimestamp:          nil,
		DeletionGracePeriodSeconds: nil,
		Labels:                     nil,
		Annotations:                nil,
		OwnerReferences:            nil,
		Finalizers:                 nil,
		ManagedFields:              nil,
	}
	fmt.Println(x2)
}

func (s *K8sParamsSuite) TestLocal() {
	x := V1.CreateDeploymentReq{
		Name:                        "ian",
		Namespace:                   "",
		MetadataName:                "ian",
		RepNums:                     0,
		MetadataSelectorLabels:      "",
		SpecSelectorLabels:          "",
		TemplateSelectorMatchLabels: "",
		ContainerImage:              "",
		ContainerName:               "",
		Containers:                  "",
		Content:                     nil,
	}

	res, err := NewDeployParams(&x)
	s.NoError(err)
	fmt.Println(res)

}

// TestHttpClient :
func TestParamsConfiguration(t *testing.T) {
	suite.Run(t, new(K8sParamsSuite))
}
