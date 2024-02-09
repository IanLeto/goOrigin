package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"goOrigin/API/V1"
	"goOrigin/internal/model"
	"goOrigin/pkg/k8s"
	"goOrigin/pkg/storage"
	"goOrigin/pkg/utils"
	"io/ioutil"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"strconv"
	"strings"
	"time"
)

func CreateDeployment(c *gin.Context, req *V1.CreateDeploymentReq) (string, error) {

	var (
		res = ""
		err error
	)
	deploy, err := model.NewDeployParams(req)
	if err != nil {
		return "", err
	}
	// 启用mongo 事务
	err = storage.GlobalMongo.Client.UseSession(c, func(sessionContext mongo.SessionContext) error {
		err = sessionContext.StartTransaction()
		if err != nil {
			return err
		}
		data, err := storage.GlobalMongo.DB.Collection("pod").InsertOne(c, &deploy)
		dep, err := k8s.Conn.CreateDeploy(c, "default", deploy.Deployment)
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
func CreateDeploymentV2(c *gin.Context, req *V1.CreateDeploymentReq) (string, error) {

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
	err = storage.GlobalMongo.Client.UseSession(c, func(sessionContext mongo.SessionContext) error {
		err = sessionContext.StartTransaction()
		if err != nil {
			return err
		}
		data, err := storage.GlobalMongo.DB.Collection("pod").InsertOne(c, &deploy)
		dep, err := k8s.Conn.DynamicClient.Resource(deployments).Namespace("default").Create(context.TODO(),
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

func UpdateDeployment(c *gin.Context, req *V1.UpdateDeploymentReq) (string, error) {
	var (
		err error
	)
	dep, err := k8s.Conn.UpdateDeploy(c, req.Name, utils.StrDefault(req.Namespace, "default"), req.Image)
	if err != nil {
		return "", err
	}
	return dep.Name, err
}

func DeleteDeployment(c *gin.Context, ns, name string) (string, error) {
	var (
		err error
	)
	res, err := k8s.Conn.DeleteDeploy(c, utils.StrDefault(ns, "default"), name)
	if err != nil {
		return "", err
	}
	return res, err
}
func ListDeployments(c *gin.Context, ns string) (interface{}, error) {
	var (
		err error
	)
	res, err := k8s.Conn.ListDeploy(c, utils.StrDefault(ns, "default"))
	if err != nil {
		return "", err
	}
	return res, err
}

func GetConfigMapDetail(c *gin.Context, req *V1.GetConfigMapRequestInfo) (interface{}, error) {
	var (
		ns   = req.NS
		name = req.Name
	)
	return k8s.Conn.GetConfigMapDetail(c, ns, name)

}

func CreateDeploymentDynamic(c *gin.Context, req *V1.CreateDeploymentDynamicRequest) (interface{}, error) {
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
	result, err := k8s.Conn.DynamicClient.Resource(deploymentRes).Namespace(req.Namespace).Create(c, deployment, metav1.CreateOptions{})
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
	err = k8s.Conn.DynamicClient.Resource(deploymentRes).Namespace(namespace).Delete(c, name, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})
	logrus.Errorf("delete deploy error: %s", err)
	return err
}

func UpdateDeploymentDynamicRequest(c *gin.Context, req *V1.UpdateDeploymentDynamicRequest) (interface{}, error) {
	var (
		err error
	)
	deploymentRes := schema.GroupVersionResource{
		Group:    "apps",
		Version:  "v1",
		Resource: "deployments",
	}
	deployment, err := k8s.Conn.DynamicClient.Resource(deploymentRes).Namespace(req.Namespace).Get(c, req.Name, metav1.GetOptions{})
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
	_, err = k8s.Conn.DynamicClient.Resource(deploymentRes).Namespace(req.Namespace).Update(c, deployment, metav1.UpdateOptions{})
	return nil, err

ERR:
	{
		return deployment, err
	}
}

type Entry struct {
	Timestamp int64  `json:"timestamp"`
	Date      string `json:"date"`
	Content   int    `json:"content"`
}

func reverseArray(arr []string) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
func reverseArray2(arr []interface{}) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func GetPods(c *gin.Context, req *V1.GetPodRequest) (*V1.GetPodResponse, error) {
	var (
		err error
	)
	config, err := rest.InClusterConfig()
	if err != nil {
		logrus.Errorf("get config error: %s", err)
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		logrus.Errorf("get clientset error: %s", err)
		return nil, err
	}
	pods, err := clientset.CoreV1().Pods(req.Ns).List(c, metav1.ListOptions{})
	if err != nil {
		logrus.Errorf("get pods error: %s", err)
		return nil, err
	}
	res := &V1.GetPodResponse{}
	for _, pod := range pods.Items {
		fmt.Println(pod.Name)
	}
	res.Item = pods.Items
	return res, nil
}

func GetCurrentLogs(c *gin.Context, req *V1.GetLogsReq) (*V1.GetLogsRes, error) {
	var (
		err       error
		conn      = k8s.Conn
		byteLimit = int64(req.LimitByte)
		lineLimit = int64(req.LimitLine)
		//sinceTime = int64(100 * time.Second)
		res = &V1.GetLogsRes{Contents: nil}
	)
	logOptions := &v1.PodLogOptions{
		Container:                    req.Container,
		Timestamps:                   true,       // 是否附带时间戳
		TailLines:                    &lineLimit, // 最大行数限制
		LimitBytes:                   &byteLimit, // 最大字节数限制
		InsecureSkipTLSVerifyBackend: false,
	}

	reader, err := conn.ClientSet.CoreV1().RESTClient().Get().Namespace(req.Ns).Name(req.PodID).Resource(
		"pods").SubResource("log").VersionedParams(logOptions, scheme.ParameterCodec).Stream(c)
	if err != nil {
		panic(err)
	}
	content, err := ioutil.ReadAll(reader) // 没有换行符号？？？
	contents := strings.Split(string(content), "\n")
	if req.Size == 0 {
		req.Size = 100
	}
	var isForward bool = req.FromDate != "" && req.ToDate != ""

	switch {
	case len(contents) == 0: // 无数据 返回空
		break
	case isForward: // 按时间段查询，contents 返回5000行
		break
	case isForward && len(contents) <= req.Size: // 日志少于期望查询数量，无论怎样都会返回所有日志
		contents = contents[0:req.Size]
	case !isForward && req.Location == "begin": // 从头开始查询
		contents = contents[0:req.Size]
	case !isForward && len(contents) <= req.Size: // 没有翻页，且总日志量少于期望数量，直接返回所有日志
		contents = contents[0:req.Size]
	case !isForward && req.Location == "end": // 从尾部
		contents = contents[len(contents)-1-req.Size : len(contents)-1]
	case req.Location == "" && req.FromDate == "" && req.ToDate == "":
		contents = contents[len(contents)-1-req.Size : len(contents)-1]
	default:
		contents = contents[len(contents)-1-req.Size : len(contents)-1]
	}

	lines := contents
	entries := make([]Entry, 0)
	var fromTimestamp, toTimestamp int64
	if req.FromDate != "" {
		fromTime, err := time.Parse(time.RFC3339Nano, req.FromDate)
		fromTimestamp = fromTime.UnixNano()
		if err != nil {
			fmt.Printf("Error parsing from date: %v\n", err)
			return nil, err
		}
		toTimest, err := time.Parse(time.RFC3339Nano, req.ToDate)
		toTimestamp = toTimest.UnixNano()
		if err != nil {
			fmt.Printf("Error parsing from date: %v\n", err)
			return nil, err
		}
	}
	fmt.Println(fromTimestamp)
	count := 0
	// 定义一个函数类型，用于处理不同的条件
	type entryHandler func(timestamp int64, entry Entry) bool
	// 根据条件选择合适的处理方式
	var handleEntry entryHandler
	// 如果向后翻页，
	if isForward && req.Step >= 0 {
		handleEntry = func(timestamp int64, entry Entry) bool {
			// 如果当前数据的时间戳大于等于前端传入的时间片段的最大值,也就是todata
			if timestamp > toTimestamp {
				entries = append(entries, entry)
				return true
			}
			return false
		}
	} else if isForward && req.Step < 0 {
		handleEntry = func(timestamp int64, entry Entry) bool {
			// 因为是向前翻页，所以需要反转数组
			// 如果当前数据的时间戳小于等于结束时间，就返回
			if timestamp < fromTimestamp {
				entries = append(entries, entry)
				return true
			}
			return false
		}
	} else {
		handleEntry = func(timestamp int64, entry Entry) bool {
			entries = append(entries, entry)
			return true
		}
	}
	// 向前翻页，需要反转数组
	if isForward && req.Step < 0 {
		reverseArray(lines)
	}

	for _, line := range lines {
		parts := strings.SplitN(line, " ", 2)
		// 正常的数据格式为：时间戳 内容
		if len(parts) >= 2 {
			date := parts[0]
			content, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Printf("Error converting content to integer: %v\n", err)
				continue
			}
			timestamp, err := time.Parse(time.RFC3339Nano, date)
			if err != nil {
				fmt.Printf("Error parsing date: %v\n", err)
				continue
			}
			entry := Entry{
				Timestamp: timestamp.UnixNano(),
				Date:      date,
				Content:   content,
			}

			if handleEntry(timestamp.UnixNano(), entry) {
				count++
				if count >= req.Size {
					break
				}
			}
		}
	}
	if len(entries) == 0 {
		return res, nil
	}
	var fn = func(arr []Entry) {
		for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	if req.Step < 0 {
		fn(entries)
	}
	for _, entry := range entries {
		var epl Entry = entry
		res.Contents = append(res.Contents, epl)
	}
	res.FromDate = string(entries[0].Date) //
	res.ToDate = string(entries[len(entries)-1].Date)

	return res, err

}

func init() {
	utils.NoError(k8s.InitK8s())
}
