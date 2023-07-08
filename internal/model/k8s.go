package model

import (
	"encoding/json"
	"github.com/fatih/structs"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
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

func convJson(s string) map[string]string {
	var res map[string]string
	_ = json.Unmarshal([]byte(s), &res)
	return res
}

func NewDeployParams(req *params.CreateDeploymentReq) (*Deploy, error) {

	var (
		dep = &Deploy{
			Id:         primitive.ObjectID{},
			CreateDate: primitive.Timestamp{},
			UpdateDate: primitive.Timestamp{},
			Name:       "",
			NS:         "",
			Deployment: &appsv1.Deployment{
				TypeMeta:   metav1.TypeMeta{},
				ObjectMeta: metav1.ObjectMeta{},
				Spec:       appsv1.DeploymentSpec{},
				Status:     appsv1.DeploymentStatus{},
			},
		}
	)
	dep.NS = req.Namespace
	dep.Name = req.Name
	dep.Deployment.ObjectMeta = metav1.ObjectMeta{
		Name:   req.MetadataName,
		Labels: convJson(req.MetadataSelectorLabels),
	}
	replicas := cast.ToInt32(req.RepNums)
	dep.Deployment.Spec = appsv1.DeploymentSpec{
		Replicas: &replicas,
		Selector: &metav1.LabelSelector{
			MatchLabels: convJson(req.SpecSelectorLabels),
		},
		Template: v1.PodTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{
				Labels: convJson(req.TemplateSelectorMatchLabels),
			},
			Spec: v1.PodSpec{
				Containers: []v1.Container{{
					Name:  req.ContainerName,
					Image: req.ContainerImage,
				}},
			},
		},
	}
	return dep, nil
}

func NewDeployParamsDetail(req *params.CreateDeploymentReq) (*Deploy, error) {
	var (
		dep                   = &Deploy{}
		err                   error
		containers            []v1.Container
		templateMetadataLabel = map[string]string{}
		specSelectorLabel     = map[string]string{}
	)

	err = json.Unmarshal([]byte(req.Containers), &containers)
	r := int32(req.Content.Spec.Replicas)
	err = json.Unmarshal([]byte(req.Content.Spec.Template.Metadata.Labels), &templateMetadataLabel)
	if err != nil {
		logrus.Errorf("格式化容器规格-模板-元数据-标签 err %s", err)
	}
	err = json.Unmarshal([]byte(req.Content.Spec.Selector.MatchLabels), &specSelectorLabel)
	if err != nil {
		logrus.Errorf("格式化容器规格-Selector-标签 err %s", err)
	}
	dep.ObjectMeta.Name = req.Content.Metadata.Name
	dep.ObjectMeta.Namespace = utils.StrDefault(req.Namespace, "default")
	dep.NS = req.Namespace
	dep.Spec = appsv1.DeploymentSpec{ // pod 中的容器的详细定义
		Replicas: &r, // 副本数量
		Selector: &metav1.LabelSelector{
			MatchLabels:      specSelectorLabel,
			MatchExpressions: nil,
		},
		Template: v1.PodTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{ // ObjectMeta 元数据信息
				Name:         req.Content.Spec.Template.Metadata.Name,  // todo
				GenerateName: "",                                       // prefix
				Namespace:    req.Content.Spec.Template.Metadata.NS,    //
				Labels:       cast.ToStringMapString(req.Content.Spec), // todo
			},
			Spec: v1.PodSpec{ // 数据卷配置
				Volumes: []v1.Volume{
					{
						Name: "", // 数据卷名称
						VolumeSource: v1.VolumeSource{
							HostPath: &v1.HostPathVolumeSource{
								Path: "",  //
								Type: nil, //
							}, // 挂载主机的路径
							EmptyDir: &v1.EmptyDirVolumeSource{
								Medium:    "",
								SizeLimit: nil,
							},
							GCEPersistentDisk:     nil, // 一种 数据类型
							AWSElasticBlockStore:  nil, // 一种 数据类型
							Secret:                nil,
							NFS:                   nil,
							ISCSI:                 nil,
							Glusterfs:             nil,
							PersistentVolumeClaim: nil,
							RBD:                   nil,
							FlexVolume:            nil,
							Cinder:                nil,
							CephFS:                nil,
							Flocker:               nil,
							DownwardAPI:           nil,
							FC:                    nil,
							AzureFile:             nil,
							ConfigMap:             nil,
							VsphereVolume:         nil,
							Quobyte:               nil,
							AzureDisk:             nil,
							PhotonPersistentDisk:  nil,
							Projected:             nil,
							PortworxVolume:        nil,
							ScaleIO:               nil,
							StorageOS:             nil,
							CSI:                   nil,
							Ephemeral:             nil,
						},
					}},
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

type Log struct {
	FromDate string
	ToDate   string
	Content  []string
}
