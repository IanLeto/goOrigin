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
	Info []*CreateNodeInfo
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
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Region     string `json:"region"`
	Content    string `json:"content"`
	Depend     string `json:"depend"`
	FatherId   uint   `json:"father_id"`
	FatherName string `json:"father_name"`
	Done       *bool  `json:"done"`
	Note       string `json:"note"`
	Status     string `json:"status"`
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
