package repository

import (
	"goOrigin/internal/model/dao"
	"goOrigin/internal/model/entity"
	"strings"
)

// ToDAO converts a NodeEntity to a TNode (DAO).
func ToDAO(node *entity.NodeEntity) *dao.TNode {
	return &dao.TNode{
		Name:     node.Name,
		Content:  node.Content,
		Depend:   node.Depend,
		FatherID: node.ParentID,
		Done:     node.Done,
		Status:   node.Status,
		Region:   node.Region,
		Note:     node.Note,
	}
}

// ToEntity converts a TNode (DAO) to a NodeEntity.
func ToEntity(tnode *dao.TNode) *entity.NodeEntity {
	return &entity.NodeEntity{
		ID:      tnode.ID,
		Name:    tnode.Name,
		Content: tnode.Content,
		Depend:  tnode.Depend,

		ParentID: tnode.FatherID,
		Done:     tnode.Done,
		Status:   tnode.Status,
		Region:   tnode.Region,
		Note:     tnode.Note,
		Tags:     strings.Split(tnode.Tags, ","), // Convert the string back into a slice of strings
		// Children and Nodes fields handling here...
	}
}
