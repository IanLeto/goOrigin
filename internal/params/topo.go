package params

type CreateNodeInfo struct {
	Name     string   `json:"name"`
	Content  string   `json:"content"`
	Depend   string   `json:"depend"`
	Father   string   `json:"father"`
	FatherId string   `json:"father_id"`
	Children []string `json:"children"`
	Done     bool     `json:"done"`
	Tags     []string `json:"tags"`
	Note     string   `json:"note"`
}

type CreateNodeRequest struct {
	*CreateNodeInfo
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
