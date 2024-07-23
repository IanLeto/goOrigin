package k8s_test

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/suite"
	"goOrigin/pkg/k8s"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	ctx context.Context
}

func (s *k8sClientSuite) SetupTest() {
	s.NoError(k8s.InitK8s())
	s.Clientset = fake.NewSimpleClientset()
	s.ctx = context.Background()
}

func (s *k8sClientSuite) TestAPIs() {

}

// 赋予一个pod 读写k8s api权限
func (s *k8sClientSuite) TestInner() {
	var (
		//rabcClient  v1.RbacV1Client
		sa          *corev1.ServiceAccount
		role        *rbacv1.Role
		roleBinding *rbacv1.RoleBinding
		err         error
	)
	//rabcClient = s.Clientset.RbacV1()
	sa, err = k8s.Conn.ClientSet.CoreV1().ServiceAccounts("rbac-demo").Create(s.ctx, &corev1.ServiceAccount{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ServiceAccount",
			APIVersion: "",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:         "hello-reader-writer",
			GenerateName: "pre",
			Namespace:    "rbac-demo",
		},
		Secrets:                      nil,
		ImagePullSecrets:             nil,
		AutomountServiceAccountToken: nil,
	}, metav1.CreateOptions{})
	s.NoError(err)
	role, err = k8s.Conn.ClientSet.RbacV1().Roles("rbac-demo").Create(s.ctx, &rbacv1.Role{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:         "hello-role",
			GenerateName: "",
			Namespace:    "rbac-demo",
		},
		Rules: []rbacv1.PolicyRule{{ // 权限规则
			Verbs:     []string{"*"}, // 所有
			Resources: []string{"pods, deployments"},
			APIGroups: []string{""}, // 所属api 组
		}},
	}, metav1.CreateOptions{})
	s.NoError(err)

	roleBinding, err = k8s.Conn.ClientSet.RbacV1().RoleBindings("rbac-demo").Create(s.ctx, &rbacv1.RoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name:         "rb",
			GenerateName: "",
			Namespace:    "rbac-demo",
		},
		// subject 是一个obj ref or user ；用来告诉k8s 帮谁；这里我们让他绑sa
		Subjects: []rbacv1.Subject{{
			Kind:      "ServiceAccount",
			APIGroup:  "",
			Name:      "hello-reader-writer",
			Namespace: "rbac-demo",
		}},
		RoleRef: rbacv1.RoleRef{
			Kind: "Role",
			Name: "hello-role",
		},
	}, metav1.CreateOptions{})
	s.Equal("hello-reader-writer", sa.Name)
	s.Equal("hello-role", role.Name)
	s.Equal("rb", roleBinding.Name)

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

	s.Clientset.AddReactor("*", func(action clienttesting.Action) (handled bool, ret rest.ResponseWrapper, err error) {
		fmt.Println(1)
		return false, nil, err
	})
	s.Clientset.AddWatchReactor("*", func(action clienttesting.Action) (handled bool, ret watch.Interface, err error) {
		fmt.Println(12)
		return false, nil, err
	})
	var mockClient = k8s.KubeConn{ClientSet: s.Clientset}
	res, err := mockClient.ListDeploy(context.TODO(), "")
	s.NoError(err)
	fmt.Println(res)
}

// Testk8sClient :
func TestK8s(t *testing.T) {
	suite.Run(t, new(k8sClientSuite))
}
