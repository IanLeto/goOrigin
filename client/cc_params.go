package client

type CCRequestInfo struct {
	Product              string `json:"product"`
	ConfigVersion        string `json:"config_version"`
	ServerVersion        string `json:"server_version"`
	ServerName           string `json:"server_name"`
	Output               bool   `json:"output"`
	ConfigTemplate       string `json:"config_template"`
	ConfigTemplatePath   string `json:"config_template_path"`
	BusinessTemplate     string `json:"business_template"`
	BusinessTemplatePath string `json:"business_template_path"`
	CommonFile           string `json:"common_file"`
	CommonFilePath       string `json:"common_file_path"`
}

type CCResponseInfo struct {
	Count int                 `json:"count"`
	Info  []CCMKCResponseInfo `json:"info"`
}

type CCMKCResponseInfo struct {
	Name          string `json:"name"`
	ConfigContent string `json:"config_content"`
}
