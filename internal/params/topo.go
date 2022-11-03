package params

type CreateNodeRequest struct {
	Name    string `json:"name"`
	Content string `json:"content"`
	Father  string `json:"father"`
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
