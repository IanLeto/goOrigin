package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"goOrigin/internal/model"
	"goOrigin/internal/params"
	"goOrigin/pkg/k8s"
	"goOrigin/pkg/storage"
	"goOrigin/pkg/utils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func CreateDeployment(c *gin.Context, req *params.CreateDeploymentReq) (string, error) {

	var (
		res = ""
		err error
	)
	deploy, err := model.NewDeployParams(req)
	if err != nil {
		return "", err
	}
	// 启用mongo 事务
	err = storage.Mongo.Client.UseSession(c, func(sessionContext mongo.SessionContext) error {
		err = sessionContext.StartTransaction()
		if err != nil {
			return err
		}
		data, err := storage.Mongo.DB.Collection("pod").InsertOne(c, &deploy)
		dep, err := k8s.K8SConn.CreateDeploy(c, "default", deploy.Deployment)
		if err != nil {
			_ = sessionContext.AbortTransaction(sessionContext)
			return err
		} else {
			_ = sessionContext.AbortTransaction(sessionContext)
		}
		res = fmt.Sprintf("%s-%s", data.InsertedID.(primitive.ObjectID), dep.Name)
		return nil
	})

	return res, err
}
func CreateDeploymentV2(c *gin.Context, req *params.CreateDeploymentReq) (string, error) {

	var (
		res = ""
		err error
	)
	deploy, err := model.NewDeployParamsV2(req)
	if err != nil {
		return "", err
	}
	deployments := schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "deployments"}
	// 启用mongo 事务
	err = storage.Mongo.Client.UseSession(c, func(sessionContext mongo.SessionContext) error {
		err = sessionContext.StartTransaction()
		if err != nil {
			return err
		}
		data, err := storage.Mongo.DB.Collection("pod").InsertOne(c, &deploy)
		dep, err := k8s.K8SConn.DynamicClient.Resource(deployments).Namespace("default").Create(context.TODO(),
			&unstructured.Unstructured{Object: deploy}, metav1.CreateOptions{})

		if err != nil {
			_ = sessionContext.AbortTransaction(sessionContext)
			return err
		} else {
			_ = sessionContext.AbortTransaction(sessionContext)
		}
		res = fmt.Sprintf("%s-%s", data.InsertedID.(primitive.ObjectID), dep)
		return nil
	})

	return res, err
}

func UpdateDeployment(c *gin.Context, req *params.UpdateDeploymentReq) (string, error) {
	var (
		err error
	)
	dep, err := k8s.K8SConn.UpdateDeploy(c, req.Name, utils.StrDefault(req.Namespace, "default"), req.Image)
	if err != nil {
		return "", err
	}
	return dep.Name, err
}

func DeleteDeployment(c *gin.Context, ns, name string) (string, error) {
	var (
		err error
	)
	res, err := k8s.K8SConn.DeleteDeploy(c, utils.StrDefault(ns, "default"), name)
	if err != nil {
		return "", err
	}
	return res, err
}
func ListDeployments(c *gin.Context, ns string) (interface{}, error) {
	var (
		err error
	)
	res, err := k8s.K8SConn.ListDeploy(c, utils.StrDefault(ns, "default"))
	if err != nil {
		return "", err
	}
	return res, err
}

func GetConfigMapDetail(c *gin.Context, req *params.GetConfigMapRequestInfo) (interface{}, error) {
	var (
		ns   = req.NS
		name = req.Name
	)
	return k8s.K8SConn.GetConfigMapDetail(c, ns, name)

}

func CreateDeploymentDynamic(c *gin.Context, req *params.CreateDeploymentDynamicRequest) (interface{}, error) {
	var (
		err error
		obj map[string]interface{}
	)
	deploymentRes := schema.GroupVersionResource{
		Group:    "apps",
		Version:  "v1",
		Resource: "deployments",
	}
	err = json.Unmarshal([]byte(req.Object), &obj)
	deployment := &unstructured.Unstructured{Object: map[string]interface{}{}}
	result, err := k8s.K8SConn.DynamicClient.Resource(deploymentRes).Namespace(req.Namespace).Create(c, deployment, metav1.CreateOptions{})
	return result, err
}
func DeleteDeploymentDynamic(c *gin.Context, name, namespace string) error {
	var (
		err error
	)
	deploymentRes := schema.GroupVersionResource{
		Group:    "apps",
		Version:  "v1",
		Resource: "deployments",
	}

	deletePolicy := metav1.DeletePropagationForeground
	err = k8s.K8SConn.DynamicClient.Resource(deploymentRes).Namespace(namespace).Delete(c, name, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})
	logrus.Errorf("delete deploy error: %s", err)
	return err
}

func UpdateDeploymentDynamicRequest(c *gin.Context, req *params.UpdateDeploymentDynamicRequest) (interface{}, error) {
	var (
		err error
	)
	deploymentRes := schema.GroupVersionResource{
		Group:    "apps",
		Version:  "v1",
		Resource: "deployments",
	}
	deployment, err := k8s.K8SConn.DynamicClient.Resource(deploymentRes).Namespace(req.Namespace).Get(c, req.Name, metav1.GetOptions{})
	containers, found, err := unstructured.NestedSlice(deployment.Object, "spec", "template", "spec", "containers")
	if err != nil || !found || containers == nil {
		logrus.Errorf("deployment containers not found or error in spec: %v", err)
		goto ERR
	}
	err = unstructured.SetNestedField(containers[0].(map[string]interface{}), req.Image, "image")
	if err != nil {
		goto ERR
	}
	err = unstructured.SetNestedField(deployment.Object, containers, "spec", "template", "spec", "containers")
	if err != nil {
		goto ERR
	}
	_, err = k8s.K8SConn.DynamicClient.Resource(deploymentRes).Namespace(req.Namespace).Update(c, deployment, metav1.UpdateOptions{})
	return nil, err

ERR:
	{
		return deployment, err
	}
}
