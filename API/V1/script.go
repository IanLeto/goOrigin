package V1

type CreateScriptRequest struct {
	ID         int      `json:"id"`
	Name       string   `json:"name"`
	Comment    string   `json:"comment"`
	Type       string   `json:"type"`
	Content    string   `json:"content"`
	File       string   `json:"file"`
	Uploader   string   `json:"uploader"`
	CreateTime int      `json:"create_time"`
	UpdateTime int      `json:"update_time"`
	System     string   `json:"system"`
	IsFile     bool     `json:"isFile"`
	Timeout    int      `json:"timeout"`
	Tags       []string `json:"tags"`
	UsedTime   int      `json:"used_time"`
}

type QueryScriptRequest struct {
	Name string `json:"name"`
	Type string `json:"type"`
	File string `json:"file"`
	Key  string `json:"key"`
	Tags string `json:"tags"`
}

type QueryScriptListResponse struct {
	Infos []*QueryScriptListResponseInfo `json:"infos"`
}

type QueryScriptListResponseInfo struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Comment    string   `json:"comment"`
	Type       string   `json:"type"`
	Content    string   `json:"content"`
	File       string   `json:"file"`
	Uploader   string   `json:"uploader"`
	CreateTime int      `json:"create_time"`
	UpdateTime int      `json:"update_time"`
	System     string   `json:"system"`
	IsFile     bool     `json:"isFile"`
	Timeout    int      `json:"timeout"`
	Tags       []string `json:"tags"`
}
