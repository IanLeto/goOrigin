package V1

type CreateNodeInfo struct {
	Name       string   `json:"name"`
	Content    string   `json:"content"`
	Depend     string   `json:"depend"`
	FatherName string   `json:"father_name"`
	FatherId   uint     `json:"father_id"`
	Region     string   `json:"region"`
	Children   []string `json:"children"`
	Done       bool     `json:"done"`
	Tags       []string `json:"tags"`
	Note       string   `json:"note"`
}

type CreateNodeRequest struct {
	*CreateNodeInfo
}

type CreateNodesRequest struct {
	Region string `json:"region"`
	Info   string `json:"info"`
}

type CreateNodeResponse struct {
	Name    string `json:"name"`
	Content string `json:"content"`
	Father  string `json:"father"`
	Next    string `json:"next"`
}

type GetNodeRequest struct {
	Name    string `json:"name"`
	Content string `json:"content"`
	Father  string `json:"father"`
	Region  string `json:"region"`
}
type GetNodeResponse struct {
}

type GetNodeListRequest struct {
	Name    string `json:"name"`
	Content string `json:"content"`
	Father  string `json:"father"`
	Region  string `json:"region"`
}
type GetNodeListResponse struct {
}

type GetTopoResponse struct {
	Name    string      `json:"name"`
	Content string      `json:"content"`
	Depend  string      `json:"depend"`
	Done    bool        `json:"done"`
	Tags    []string    `json:"tags"`
	Note    string      `json:"note"`
	Nodes   interface{} `json:"nodes"`
}

type UpdateNodeRequest struct {
	Name     string   `json:"name"`
	Content  string   `json:"content"`
	Depend   string   `json:"depend"`
	FatherId uint     `json:"father_id"`
	Done     *bool    `json:"done"`
	Note     string   `json:"note"`
	Status   string   `json:"status"`
	Tags     []string `json:"tags"`
	Children []string `json:"children"`
}

type SearchNodeRequest struct {
	Name    string `json:"name"`
	Content string `json:"content"`
	Father  string `json:"father"`
	Keyword string `json:"keyword"`
}

type SearchNodeResponse struct {
	Items []interface{} `json:"items"`
}
