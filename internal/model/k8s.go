package model

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"goOrigin/internal/params"
	"goOrigin/pkg/utils"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Deploy struct {
	Id         primitive.ObjectID  `json:"id" bson:"_id,omitempty"`
	CreateDate primitive.Timestamp `json:"create_date" bson:"create_date"`
	UpdateDate primitive.Timestamp `json:"update_time" bson:"update_date"`
	Name       string              `bson:"name"`
	NS         string              `bson:"ns"`
	*appsv1.Deployment
}

func NewDeployParams(req *params.CreateDeploymentReq) (*Deploy, error) {
	var (
		dep        = &Deploy{}
		err        error
		selector   = map[string]string{}
		labels     = map[string]string{}
		containers []v1.Container
	)

	err = json.Unmarshal([]byte(req.RepSelector), &selector)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(req.TemplateLabels), &labels)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(req.Containers), &containers)

	dep.ObjectMeta.Name = req.Name
	dep.ObjectMeta.Namespace = utils.StrDefault(req.Namespace, "default")
	dep.NS = req.Namespace
	dep.Spec.Replicas = &req.RepNums
	dep.Spec.Selector = &metav1.LabelSelector{MatchLabels: selector}
	dep.Spec.Template = v1.PodTemplateSpec{ObjectMeta: metav1.ObjectMeta{
		Labels: labels,
	}}
	dep.Spec.Template.Spec = v1.PodSpec{
		Containers: containers,
	}
	return dep, err
}
