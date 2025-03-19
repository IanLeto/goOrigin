package repository

import (
	"goOrigin/internal/model/dao"
	"goOrigin/internal/model/entity"
	"strings"
)

// ToTnode converts a NodeEntity to a TNode (DAO).
func ToTnode(node *entity.NodeEntity) *dao.TNode {
	return &dao.TNode{
		Name:     node.Name,
		Content:  node.Content,
		Depend:   node.Depend,
		ParentID: node.ParentID,
		Done:     node.Done,
		Status:   node.Status,
		Note:     node.Note,
	}
}

// ToNodeEntity converts a TNode (DAO) to a NodeEntity.
func ToNodeEntity(tnode *dao.TNode) *entity.NodeEntity {
	return &entity.NodeEntity{
		ID:      tnode.ID,
		Name:    tnode.Name,
		Content: tnode.Content,
		Depend:  tnode.Depend,

		ParentID: tnode.ParentID,
		Done:     tnode.Done,
		Status:   tnode.Status,
		Note:     tnode.Note,
		Tags:     strings.Split(tnode.Tags, ","), // Convert the string back into a slice of strings
		// Children and Nodes fields handling here...
	}
}
