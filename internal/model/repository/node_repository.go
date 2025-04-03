package repository

import (
	"encoding/json"
	"goOrigin/internal/model/dao"
	"goOrigin/internal/model/entity"
)

func ToTnode(node *entity.NodeEntity) *dao.TNode {
	// 序列化 Tags 和 DependIDs 为 JSON 字符串
	tagsJSON, _ := json.Marshal(node.Tags)
	dependJSON, _ := json.Marshal(node.DependIDs)

	return &dao.TNode{
		Name:      node.Name,
		Content:   node.Content,
		ParentID:  node.ParentID,
		DependIDs: string(dependJSON),
		Done:      node.Done,
		Status:    node.Status,
		Note:      node.Note,
		Region:    node.Region,
		Tags:      string(tagsJSON),
	}
}

func ToNodeEntity(tnode *dao.TNode) *entity.NodeEntity {
	var tags []string
	var dependIDs []uint

	// 反序列化 Tags
	if tnode.Tags != "" {
		_ = json.Unmarshal([]byte(tnode.Tags), &tags)
	}

	// 反序列化 DependIDs
	if tnode.DependIDs != "" {
		_ = json.Unmarshal([]byte(tnode.DependIDs), &dependIDs)
	}

	return &entity.NodeEntity{
		ID:        tnode.ID,
		Name:      tnode.Name,
		Content:   tnode.Content,
		ParentID:  tnode.ParentID,
		DependIDs: dependIDs,
		Done:      tnode.Done,
		Status:    tnode.Status,
		Note:      tnode.Note,
		Region:    tnode.Region,
		Tags:      tags,
	}
}

func ToNodeEntities(nodes []*dao.TNode) []*entity.NodeEntity {
	result := make([]*entity.NodeEntity, 0, len(nodes))
	for _, t := range nodes {
		result = append(result, ToNodeEntity(t))
	}
	return result
}
