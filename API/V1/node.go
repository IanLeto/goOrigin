package V1

type CreateNodeInfo struct {
	Name     string   `json:"name"`
	Content  string   `json:"content"`
	Depend   string   `json:"depend"`
	ParentId uint     `json:"parent_id"`
	Region   string   `json:"region"`
	Done     bool     `json:"done"`
	Tags     []string `json:"tags"`
	Note     string   `json:"note"`
	Status   string   `json:"status"`
}

type CreateNodeRequest struct {
	Region string `json:"region"`
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

type GetNodesRequest struct {
	Name        string `json:"name"`
	Content     string `json:"content"`
	ParentID    string `json:"parent_id"`
	DeadLine    string `json:"dead_line"`
	Status      string `json:"status"`
	Achievement int    `json:"achievement"`
}
type GetNodeResponse struct {
}

type GetNodeListRequest struct {
	Name    string `json:"name"`
	Content string `json:"content"`
	Father  string `json:"father"`
	Region  string `json:"region"`
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
	ParentID   uint   `json:"parent_id"`
	FatherName string `json:"father_name"`
	Done       *bool  `json:"done"`
	Note       string `json:"note"`
	Status     string `json:"status"`
}

type SearchNodeResponse struct {
	Items []interface{} `json:"items"`
}
