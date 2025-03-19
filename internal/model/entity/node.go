package entity

import (
	"context"
	"github.com/sirupsen/logrus"
	//dao2 "goOrigin/internal/model/dao"
)

type NodeEntity struct {
	ID       uint          `json:"id"`
	Name     string        `json:"name"`
	Content  string        `json:"content"`
	Depend   string        `json:"depend"`
	ParentID uint          `json:"parent_id"`
	Done     bool          `json:"done"`
	Status   string        `json:"status"`
	Tags     []string      `json:"tags"`
	Note     string        `json:"note"`
	Region   string        `json:"region"`
	Children []string      `json:"children"`
	Nodes    []*NodeEntity `json:"nodes"`
}

// Epl 接收一个callback ，callback 为递归查询子节点的实现，目前支持 mysql ， searchlight
func (n *NodeEntity) Epl(fn func(entity *NodeEntity) (*NodeEntity, error)) {
	for _, child := range n.Children {
		var epl = &NodeEntity{
			Name: child,
		}
		result, err := fn(epl)
		if err != nil {
			logrus.Errorf("获取topo失败 %s", err)
			return
		}
		n.Nodes = append(n.Nodes, result)
	}
}

// ToNodes todo 获取node的所有节点，并将其转为slice
func (n *NodeEntity) ToNodes() []*NodeEntity {
	var (
		res []*NodeEntity
	)
	for _, v := range n.Nodes {
		if res != nil {
			res = append(res, v)
		}
		if len(v.Nodes) != 0 {
			v.ToNodes()
		}

	}
	return res
}

type Topo struct {
	*NodeEntity
	Children []*NodeEntity `json:"children"`
}

func GetTopo(ctx context.Context, root *NodeEntity) *NodeEntity {
	return nil
}
