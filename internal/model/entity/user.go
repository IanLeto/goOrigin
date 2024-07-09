package entity

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"goOrigin/pkg/utils"
	"io/ioutil"
	v1 "k8s.io/api/authorization/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
	"strings"
	"sync"
	"time"
)

var (
	httpClient *http.Client
	once       sync.Once
)

func getHTTPClient() *http.Client {
	once.Do(func() {
		httpClient = &http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:        100,
				MaxIdleConnsPerHost: 100,
				IdleConnTimeout:     90 * time.Second,
			},
		}
	})
	return httpClient
}

type User interface {
	ToUserEntity(token, url string) VersionUserEntity
	Auth(token, url string) (bool, error)
}

type UserRedis struct {
}

func (u *UserRedis) ToUserEntity(token, url string) VersionUserEntity {
	//TODO implement me
	panic("implement me")
}

func (u *UserRedis) Auth(token, url string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

type UserStr string

func (u *UserStr) ToUserEntity(token, url string) VersionUserEntity {
	host, err := utils.GetDomain(url)
	if err != nil {
		return nil
	}
	switch host {
	case "cp":
		return &CpaasUserEntity{&UserEntity{Token: token, LoginUrl: url}}
	case "cebp":
		return &CebpaasUserEntity{}
	default:
		return &CebpaasUserEntity{}
	}
}

func (u *UserStr) Auth(token, url string) (bool, error) {
	var (
		userEntity VersionUserEntity
	)
	userEntity = u.ToUserEntity(token, url)
	review := userEntity.SubjectReview("", "", "")
	return review.Status.Allowed, nil

}

type VersionUserEntity interface {
	SubjectReview(token, project, url string) v1.SelfSubjectAccessReview
}

type UserEntity struct {
	Iss      string `json:"iss"`
	Sub      string `json:"sub"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Token    string `json:"token"`
	LoginUrl string `json:"login_url"`
}

func (u *UserEntity) ParseToken(token string) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return
	}

	payload := parts[1]
	userInfoJSON, err := base64.RawURLEncoding.DecodeString(payload)
	if err != nil {
		return
	}

	err = json.Unmarshal(userInfoJSON, u)
	if err != nil {
		return
	}
}

func (u *UserEntity) Auth(token, url, project string) (v1.SelfSubjectAccessReview, error) {
	var (
		err error
	)

	requestBody := map[string]interface{}{
		"kind":       "SubjectAccessReview",
		"apiVersion": "authorization.k8s.io/v1",
		"spec": map[string]interface{}{
			"user": u.Name,
			"resourceAttributes": map[string]string{
				"verb":      "delete",
				"resource":  "projects",
				"namespace": project,
			},
		},
	}

	// 将 JSON 参数编码为字节数组
	requestBodyBytes, _ := json.Marshal(requestBody)

	// 创建 POST 请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		panic(err)
	}

	// 设置请求的 Content-Type 为 application/json
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("token", "application/json")

	// 创建 HTTP 客户端
	client := getHTTPClient()

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer func() { _ = resp.Body.Close() }()

	// 读取响应体
	body, _ := ioutil.ReadAll(resp.Body)

	// 打印响应状态码和响应体
	fmt.Println("Response status:", resp.Status)
	fmt.Println("Response body:", string(body))

	return v1.SelfSubjectAccessReview{
		TypeMeta:   metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{},
		Spec:       v1.SelfSubjectAccessReviewSpec{},
		Status: v1.SubjectAccessReviewStatus{
			Allowed:         true,
			Denied:          false,
			Reason:          "",
			EvaluationError: "",
		},
	}, err

}

type CpaasUserEntity struct {
	*UserEntity
}

func (u *CpaasUserEntity) SubjectReview(token, project, url string) v1.SelfSubjectAccessReview {
	res, err := u.UserEntity.Auth(token, url, project)
	if err != nil {
		return v1.SelfSubjectAccessReview{}
	}
	return res
}

func (u *CpaasUserEntity) ToUserEntity(token, url string) VersionUserEntity {
	//TODO implement me
	panic("implement me")
}

func (u *CpaasUserEntity) Auth(token, url string) (bool, error) {
	res, err := u.UserEntity.Auth(token, url, "")
	return res.Status.Allowed, err
}

type CebpaasUserEntity struct{}

func (u *CebpaasUserEntity) SubjectReview(token, project, url string) v1.SelfSubjectAccessReview {
	//TODO implement me
	return v1.SelfSubjectAccessReview{
		TypeMeta:   metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{},
		Spec:       v1.SelfSubjectAccessReviewSpec{},
		Status: v1.SubjectAccessReviewStatus{
			Allowed:         true,
			Denied:          false,
			Reason:          "",
			EvaluationError: "",
		},
	}
}

func (u *CebpaasUserEntity) ToUserEntity(token, url string) VersionUserEntity {
	//TODO implement me
	panic("implement me")
}

func (u *CebpaasUserEntity) Auth(token, url string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

// Base64Encode encodes data to a base64 string without padding
func Base64Encode(data []byte) string {
	return base64.RawURLEncoding.EncodeToString(data)
}

// Base64Decode decodes a base64 string to data
func Base64Decode(data string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(data)
}
