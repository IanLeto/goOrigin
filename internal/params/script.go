package params

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
}

type QueryScriptRequest struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Comment string `json:"comment"`
	Type    string `json:"type"`
	Content string `json:"content"`
	File    string `json:"file"`
}

type QueryScriptListResponse struct {
	Infos []*QueryScriptListResponseInfo `json:"infos"`
}

type QueryScriptListResponseInfo struct {
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
}
