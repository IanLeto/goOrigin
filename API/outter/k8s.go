package outter

import v1 "k8s.io/api/authorization/v1"

// 对外请求API
type SubjectAccessViewReq struct {
	Url          string `json:"url"`
	User         string `json:"user"`
	Group        string `json:"group"`
	Resource     string `json:"resource"`
	Verb         string `json:"verb"`
	Namespace    string `json:"namespace"`
	ResourceName string `json:"resource_name"`
}

type SubjectAccessReviewRes struct {
	*v1.SubjectAccessReview
}
