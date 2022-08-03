package k8s_test

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/suite"
	"goOrigin/pkg/k8s"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/rest"
	clienttesting "k8s.io/client-go/testing"
	"testing"
)

// k8sClientSuite :
type k8sClientSuite struct {
	suite.Suite
	*fake.Clientset
}

func (s *k8sClientSuite) SetupTest() {
	s.Clientset = fake.NewSimpleClientset()

}

// TestMarshal :
func (s *k8sClientSuite) TestConfig() {
	watcherStarted := make(chan struct{})
	s.Clientset.PrependWatchReactor("*", func(action clienttesting.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		fmt.Println(gvr)
		ns := action.GetNamespace()
		fmt.Println(ns)
		watchRes, err := s.Clientset.Tracker().Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		close(watcherStarted)
		return true, watchRes, nil
	})

	s.Clientset.AddProxyReactor("*", func(action clienttesting.Action) (handled bool, ret rest.ResponseWrapper, err error) {
		fmt.Println(1)
		return false, nil, err
	})
	s.Clientset.AddWatchReactor("*", func(action clienttesting.Action) (handled bool, ret watch.Interface, err error) {
		fmt.Println(12)
		return false, nil, err
	})
	//s.Clientset.PrependReactor("*", "*", func(action clienttesting.Action) (handled bool, ret runtime.Object, err error) {
	//	gvr := action.GetResource()
	//	fmt.Println(gvr)
	//	ns := action.GetNamespace()
	//	fmt.Println(ns)
	//	return true, nil, nil
	//})
	var mockClient = k8s.KubeConn{Client: s.Clientset}
	res, err := mockClient.ListDeploy(context.TODO(), "")
	s.NoError(err)
	fmt.Println(res)
}

// Testk8sClient :
func TestK8s(t *testing.T) {
	suite.Run(t, new(k8sClientSuite))
}
