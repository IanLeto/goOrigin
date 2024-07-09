package entity

import (
	"encoding/base64"
	"encoding/json"
	"goOrigin/pkg/utils"
	v1 "k8s.io/api/authorization/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

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
		return nil
	}
}

func (u *UserStr) Auth(token, url string) (bool, error) {
	var (
		userEntity VersionUserEntity
	)
	userEntity = u.ToUserEntity(token, url)
	review := userEntity.SubjectReview()
	return review.Status.Allowed, nil

}

type VersionUserEntity interface {
	SubjectReview() v1.SelfSubjectAccessReview
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

func (u *UserEntity) Auth() {

	panic(111)
	// 定义请求的 JSON 参数
	//	requestBody := map[string]interface{}{
	//		"kind":       "SubjectAccessReview",
	//		"apiVersion": "authorization.k8s.io/v1",
	//		"spec": map[string]interface{}{
	//			"user": "tool-readonly-user",
	//			"resourceAttributes": map[string]string{
	//				"namespace": "default",
	//				"verb":      "delete",
	//				"resource":  "pods",
	//			},
	//		},
	//	}
	//
	//	// 将 JSON 参数编码为字节数组
	//	requestBodyBytes, _ := json.Marshal(requestBody)
	//
	//	// 创建 POST 请求
	//	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	// 设置请求的 Content-Type 为 application/json
	//	req.Header.Set("Content-Type", "application/json")
	//
	//	// 创建 HTTP 客户端
	//	client := &http.Client{}
	//
	//	// 发送请求
	//	resp, err := client.Do(req)
	//	if err != nil {
	//		panic(err)
	//	}
	//	defer resp.Body.Close()
	//
	//	// 读取响应体
	//	body, _ := ioutil.ReadAll(resp.Body)
	//
	//	// 打印响应状态码和响应体
	//	fmt.Println("Response status:", resp.Status)
	//	fmt.Println("Response body:", string(body))
}

type CpaasUserEntity struct {
	*UserEntity
}

func (u *CpaasUserEntity) SubjectReview() v1.SelfSubjectAccessReview {
	panic(111)
}

func (u *CpaasUserEntity) ToUserEntity(token, url string) VersionUserEntity {
	//TODO implement me
	panic("implement me")
}

func (u *CpaasUserEntity) Auth(token, url string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

type CebpaasUserEntity struct{}

func (u *CebpaasUserEntity) SubjectReview() v1.SelfSubjectAccessReview {
	//TODO implement me
	return v1.SelfSubjectAccessReview{
		TypeMeta:   metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{},
		Spec:       v1.SelfSubjectAccessReviewSpec{},
		Status: v1.SubjectAccessReviewStatus{
			Allowed:         false,
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
