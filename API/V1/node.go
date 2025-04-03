package V1

type CreateNodeInfo struct {
	Name     string   `json:"name" binding:"required"`
	Content  string   `json:"content"`
	Depend   string   `json:"depend"`    // JSON 字符串或逗号分隔
	ParentID uint     `json:"parent_id"` // 0 表示根节点
	Done     bool     `json:"done"`
	Tags     []string `json:"tags"`
	Note     string   `json:"note"`
	Status   string   `json:"status"`
}

type CreateNodeRequest struct {
	Region string `json:"region" binding:"required"`
	*CreateNodeInfo
}

type GetNodeRequest struct {
	ID     uint   `form:"id" binding:"required"`
	Region string `form:"region" binding:"required"`
}

type UpdateNodeInfo struct {
	ID       uint     `json:"id" binding:"required"` // 节点 ID
	Name     string   `json:"name"`
	Content  string   `json:"content"`
	Depend   string   `json:"depend"`
	ParentID uint     `json:"parent_id"`
	Done     bool     `json:"done"`
	Tags     []string `json:"tags"`
	Note     string   `json:"note"`
	Status   string   `json:"status"`
}

type UpdateNodeRequest struct {
	Region string `json:"region" binding:"required"`
	*UpdateNodeInfo
}

type DeleteNodeRequest struct {
	ID     uint   `json:"id" binding:"required"`
	Region string `json:"region" binding:"required"`
}
type ListNodeRequest struct {
	Region   string `form:"region" binding:"required"`
	ParentID uint   `form:"parent_id"` // 可选，查某父节点下的子节点
	Status   string `form:"status"`    // 可选，支持状态过滤
	Keyword  string `form:"keyword"`   // 可选，支持 name / content 模糊搜索

	Page     int `form:"page" binding:"required,min=1"`
	PageSize int `form:"page_size" binding:"required,min=1,max=100"`
}

type BatchDeleteNodeRequest struct {
	IDs    []uint `json:"ids" binding:"required"`
	Region string `json:"region" binding:"required"`
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

type SearchNodeResponse struct {
	Items []interface{} `json:"items"`
}
