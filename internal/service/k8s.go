package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"goOrigin/internal/model"
	"goOrigin/internal/params"
	"goOrigin/pkg/k8s"
	"goOrigin/pkg/storage"
	"goOrigin/pkg/utils"
)

func CreateDeployment(c *gin.Context, req *params.CreateDeploymentReq) (string, error) {

	var (
		res = ""
		err error
	)
	deploy, err := model.NewDeployParams(req)
	err = storage.Mongo.Client.UseSession(c, func(sessionContext mongo.SessionContext) error {
		err = sessionContext.StartTransaction()
		if err != nil {
			return err
		}
		data, err := storage.Mongo.DB.Collection("pod").InsertOne(c, &deploy)
		dep, err := k8s.K8S.CreateDeploy(c, "default", deploy.Deployment)
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

func UpdateDeployment(c *gin.Context, req *params.UpdateDeploymentReq) (string, error) {
	var (
		err error
	)
	dep, err := k8s.K8S.UpdateDeploy(c, req.Name, utils.StrDefault(req.Namespace, "default"), req.Image)
	if err != nil {
		return "", err
	}
	return dep.Name, err
}

func DeleteDeployment(c *gin.Context, ns, name string) (string, error) {
	var (
		err error
	)
	res, err := k8s.K8S.DeleteDeploy(c, utils.StrDefault(ns, "default"), name)
	if err != nil {
		return "", err
	}
	return res, err
}
func ListDeployments(c *gin.Context, ns string) (interface{}, error) {
	var (
		err error
	)
	res, err := k8s.K8S.ListDeploy(c, utils.StrDefault(ns, "default"))
	if err != nil {
		return "", err
	}
	return res, err
}
