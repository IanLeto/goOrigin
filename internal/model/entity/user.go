package entity

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"goOrigin/API/outter"
	"goOrigin/pkg/utils"
	"io/ioutil"
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
	ToUserEntity(token, url, region string) EnvironmentUserEntity
	Auth(token, url, project, verb string) (bool, error)
}

type UserFromToken string

func (u *UserFromToken) ToUserEntity(token, url, region string) EnvironmentUserEntity {
	var (
		environment EnvironmentUserEntity
	)
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil
	}
	jsonInfo, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil
	}
	ephEntity := DetermineDomainType(url)
	switch entity := ephEntity.(type) {
	case *WinUserEntity:
		json.Unmarshal(jsonInfo, entity)
		environment = entity
	default:
		json.Unmarshal(jsonInfo, entity)
		environment = entity
	}
	return environment
}

func (u *UserFromToken) Auth(token, url, project, verb string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

type EnvironmentUserEntity interface {
	SubjectReview(req outter.SubjectAccessViewReq) (*outter.SubjectAccessReviewRes, int, error)
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

type CpaasUserEntity struct {
	*UserEntity
}

func (u *CpaasUserEntity) SubjectReview(req outter.SubjectAccessViewReq) outter.SubjectAccessReviewRes {
	var (
		err    error
		result outter.SubjectAccessReviewRes
	)

	requestBody := map[string]interface{}{
		"kind":       "SubjectAccessReview",
		"apiVersion": "authorization.k8s.io/v1",
		"spec": map[string]interface{}{
			"user": u.Name,
			"resourceAttributes": map[string]string{
				"verb":     req.Verb,
				"resource": req.Resource,
				"name":     req.ResourceName,
			},
		},
	}

	// 将 JSON 参数编码为字节数组
	requestBodyBytes, _ := json.Marshal(requestBody)

	// 创建 POST 请求
	request, err := http.NewRequest("POST", req.Url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		panic(err)
	}

	// 设置请求的 Content-Type 为 application/json
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("token", "application/json")

	// 创建 HTTP 客户端
	client := getHTTPClient()

	// 发送请求
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer func() { _ = resp.Body.Close() }()

	// 读取响应体
	body, _ := ioutil.ReadAll(resp.Body)

	// 打印响应状态码和响应体
	fmt.Println("Response status:", resp.Status)
	fmt.Println("Response body:", string(body))
	err = json.Unmarshal(body, &result)
	utils.NoError(err)
	return result

}

func (u *CpaasUserEntity) ToUserEntity(token, url string) EnvironmentUserEntity {
	//TODO implement me
	panic("implement me")
}

type WinUserEntity struct {
	*UserEntity
}

func (w *WinUserEntity) SubjectReview(req outter.SubjectAccessViewReq) (*outter.SubjectAccessReviewRes, int, error) {
	//TODO implement me
	panic("implement me")
}

func DetermineDomainType(domain string) EnvironmentUserEntity {
	host, _ := utils.GetDomain(domain)
	if host != "" {
		return &WinUserEntity{}
	}
	return nil
}
