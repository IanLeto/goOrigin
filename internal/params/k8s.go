package params

type CreateDeploymentReq struct {
	Name                        string `json:"name"`
	Namespace                   string `json:"namespace"`
	MetadataName                string `json:"metadata_name"`
	RepNums                     int32  `json:"rep_nums"`
	MetadataSelectorLabels      string `json:"metadata_selector_labels"`
	SpecSelectorLabels          string `json:"spec_selector_labels"`
	TemplateSelectorMatchLabels string `json:"template_selector_match_labels"`
	ContainerImage              string `json:"container_images"`
	ContainerName               string `json:"container_name"`

	Containers string          `json:"containers"`
	Content    *CreateUnStruct `json:"content"`
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
				Name   string `json:"name"`
				NS     string `json:"NS"`
			} `json:"metadata"`
			Spec struct {
				Volume []struct {
				}
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
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Image     string `json:"image"`
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

//  ---------------------------------------dynamic-----------------------------------------------------

type CreateDeploymentDynamicRequest struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Object    string `json:"Object"`
}

// UpdateDeploymentDynamicRequest 仅支持单一镜像修改
type UpdateDeploymentDynamicRequest struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Image     string `json:"Image"`
}
