package params

type CreateDeploymentReq struct {
	Name           string
	Namespace      string
	RepNums        int32
	RepSelector    string
	TemplateLabels string
	Containers     string
}

type UpdateDeploymentReq struct {
	Name      string
	Namespace string
	Image     string
}
