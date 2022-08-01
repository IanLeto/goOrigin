package model

import (
	"encoding/json"
	"github.com/fatih/structs"
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
		dep                   = &Deploy{}
		err                   error
		selector              = map[string]string{}
		labels                = map[string]string{}
		containers            []v1.Container
		templateMetadataLabel = map[string]string{}
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
	r := int32(req.Content.Spec.Replicas)
	err = json.Unmarshal([]byte(req.Content.Spec.Template.Metadata.Labels), &templateMetadataLabel)

	dep.ObjectMeta.Name = req.Content.Metadata.Name
	dep.ObjectMeta.Namespace = utils.StrDefault(req.Namespace, "default")
	dep.NS = req.Namespace
	dep.Spec = appsv1.DeploymentSpec{ // pod 中的容器的详细定义
		Replicas: &r, // 副本数量
		Selector: &metav1.LabelSelector{
			MatchLabels:      templateMetadataLabel,
			MatchExpressions: nil,
		},
		Template: v1.PodTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{ // ObjectMeta 元数据信息
				Name:                       "",
				GenerateName:               "",
				Namespace:                  "",
				UID:                        "",
				ResourceVersion:            "",
				Generation:                 0,
				CreationTimestamp:          metav1.Time{},
				DeletionTimestamp:          nil,
				DeletionGracePeriodSeconds: nil,
				Labels:                     nil,
				Annotations:                nil,
				OwnerReferences:            nil,
				Finalizers:                 nil,
				ManagedFields:              nil,
			},
			Spec: v1.PodSpec{
				Volumes:                       nil,
				InitContainers:                nil,
				Containers:                    nil,
				EphemeralContainers:           nil,
				RestartPolicy:                 "",
				TerminationGracePeriodSeconds: nil,
				ActiveDeadlineSeconds:         nil,
				DNSPolicy:                     "",
				NodeSelector:                  nil,
				ServiceAccountName:            "",
				AutomountServiceAccountToken:  nil,
				NodeName:                      "",
				HostNetwork:                   false,
				HostPID:                       false,
				HostIPC:                       false,
				ShareProcessNamespace:         nil,
				SecurityContext:               nil,
				ImagePullSecrets:              nil,
				Hostname:                      "",
				Subdomain:                     "",
				Affinity:                      nil,
				SchedulerName:                 "",
				Tolerations:                   nil,
				HostAliases:                   nil,
				PriorityClassName:             "",
				Priority:                      nil,
				DNSConfig:                     nil,
				ReadinessGates:                nil,
				RuntimeClassName:              nil,
				EnableServiceLinks:            nil,
				PreemptionPolicy:              nil,
				Overhead:                      nil,
				TopologySpreadConstraints:     nil,
				SetHostnameAsFQDN:             nil,
				OS:                            nil,
			},
		}, // 定义容器模板，dep 会以此来生成相关的容器
		Strategy:                appsv1.DeploymentStrategy{}, // 升级策略
		MinReadySeconds:         0,
		RevisionHistoryLimit:    nil,
		Paused:                  false,
		ProgressDeadlineSeconds: nil,
	}
	for _, v := range req.Content.Spec.Template.Spec.Containers {
		containers = append(containers, v1.Container{
			Name:  v.Name,
			Image: v.Image,
			//Command: v.Command,
			//Ports:   v.Ports,
		})
	}

	dep.Spec.Template.Spec = v1.PodSpec{
		Containers: []v1.Container{{
			Name:  "",
			Image: "",
			Ports: nil,
		}},
	}
	return dep, err
}

// NewDeployParamsV2 : 非结构化数据，k8s 结构化处理太多了学习过程中不适合，交给前端处理吧
func NewDeployParamsV2(req *params.CreateDeploymentReq) (map[string]interface{}, error) {
	var (
		err error
	)

	return structs.Map(req.Content), err
}
