package params

type CreateDeploymentReq struct {
	Name           string
	Namespace      string
	RepNums        int32
	RepSelector    string
	TemplateLabels string
	Containers     string
	Content        *CreateUnStruct `json:"content"`
}

type CreateUnStruct struct {
	ApiVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
		Name string `json:"name"`
	} `json:"metadata"`
	Spec struct {
		Replicas int `json:"replicas"`
		Selector struct {
			MatchLabels string `json:"matchLabels"`
		} `json:"selector"`
		Template struct {
			Metadata struct {
				Labels string `json:"labels"`
			} `json:"metadata"`
			Spec struct {
				Containers []struct {
					Image string `json:"image"`
					Name  string `json:"name"`
					Ports []struct {
						ContainerPort int    `json:"containerPort"`
						Name          string `json:"name"`
						Protocol      string `json:"protocol"`
					} `json:"ports"`
				} `json:"containers"`
			} `json:"spec"`
		} `json:"template"`
	} `json:"spec"`
}

type UpdateDeploymentReq struct {
	Name      string
	Namespace string
	Image     string
}

//  ---------------------------------------configmap-----------------------------------------------------

type CreateConfigMapRequestInfo struct {
	Name string
}

type GetConfigMapRequestInfo struct {
	Name string `json:"name"`
	NS   string `json:"NS"`
	*BaseK8sRequestInfo
}
