package k8s

import (
	"context"
	"goOrigin/pkg/utils"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k *KubeConn) CreateDeploy(ctx context.Context, ns string, dep *v1.Deployment) (*v1.Deployment, error) {
	deploymentClient := k.ClientSet.AppsV1().Deployments(utils.StrDefault(ns, "default"))
	res, err := deploymentClient.Create(ctx, dep, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return res, err
}

func (k *KubeConn) UpdateDeploy(ctx context.Context, ns, name, image string) (*v1.Deployment, error) {
	deploymentClient := k.ClientSet.AppsV1().Deployments(utils.StrDefault(ns, "default"))
	deployment, err := deploymentClient.Get(ctx, name, metav1.GetOptions{})
	deployment.Spec.Template.Spec.Containers[0].Image = image
	res, err := deploymentClient.Update(ctx, deployment, metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	return res, err
}

func (k *KubeConn) DeleteDeploy(ctx context.Context, ns, name string) (string, error) {
	deploymentClient := k.ClientSet.AppsV1().Deployments(utils.StrDefault(ns, "default"))
	deletePolicy := metav1.DeletePropagationForeground
	err := deploymentClient.Delete(ctx, name, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})
	if err != nil {
		return "", err
	}
	return "ok", err
}
func (k *KubeConn) ListDeploy(ctx context.Context, ns string) ([]map[string]interface{}, error) {
	var (
		res = make([]map[string]interface{}, 0)
	)
	deploymentClient := k.ClientSet.AppsV1().Deployments(utils.StrDefault(ns, "default"))
	deploys, err := deploymentClient.List(ctx, metav1.ListOptions{})
	for _, i := range deploys.Items {
		res = append(res, map[string]interface{}{i.Name: i.Spec})
	}
	if err != nil {
		return nil, err
	}
	return res, err
}

func (k *KubeConn) CreateService(ctx context.Context, ns string, dep *v1.Deployment) (*v1.Deployment, error) {
	//svcClient := k.ClientSet.CoreV1().Services(ns)
	//res, err := svcClient.Create(ctx, corev1.Service{
	//	TypeMeta:   metav1.TypeMeta{},
	//	ObjectMeta: metav1.ObjectMeta{},
	//	Spec:       corev1.ServiceSpec{},
	//	Status:     corev1.ServiceStatus{},
	//}, metav1.CreateOptions{})
	//if err != nil {
	//	return nil, err
	//}
	panic(1)
	//return res, err
}

func (k *KubeConn) GetConfigMapDetail(ctx context.Context, ns, name string) (*corev1.ConfigMap, error) {

	configClient := k.ClientSet.CoreV1().ConfigMaps(ns)
	configMap, err := configClient.Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return configMap, err
}

func (k *KubeConn) LeaderElection(ctx context.Context, name, ns string) {

}

func (k *KubeConn) CreateServiceAccount(ctx context.Context, obj ModelObjectMeta, refs []ModelObjectRef, ns string) (*corev1.ServiceAccount, error) {
	var (
		err              error
		objectReferences []corev1.ObjectReference
	)
	for _, r := range refs {
		objectReferences = append(objectReferences, *r.ToK8sModel())
	}

	account, err := k.ClientSet.CoreV1().ServiceAccounts(ns).Create(ctx, &corev1.ServiceAccount{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ServiceAccount",
			APIVersion: "",
		},
		ObjectMeta: *obj.ToK8sModel(),
		Secrets:    objectReferences,
		//ImagePullSecrets: ,
		AutomountServiceAccountToken: nil,
	}, metav1.CreateOptions{})
	return account, err
}
