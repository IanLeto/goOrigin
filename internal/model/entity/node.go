package entity

import (
	"context"
	//dao2 "goOrigin/internal/model/dao"
)

type NodeEntity struct {
	ID        uint          `json:"id"`
	Name      string        `json:"name"`       // 节点名称
	Content   string        `json:"content"`    // 节点描述内容
	ParentID  uint          `json:"parent_id"`  // 父节点 id
	DependIDs []uint        `json:"depend_ids"` // 依赖的其他节点 id（拓扑展示时用）
	Done      bool          `json:"done"`       // 是否完成
	Status    string        `json:"status"`     // 当前状态
	Tags      []string      `json:"tags"`       // 标签
	Note      string        `json:"note"`       // 备注
	Region    string        `json:"region"`     // 区域/域
	Children  []*NodeEntity `json:"children"`   // 子节点（树形结构）
}

type Topo struct {
	*NodeEntity
	Children []*NodeEntity `json:"children"`
}

func GetTopo(ctx context.Context, root *NodeEntity) *NodeEntity {
	return nil
}
