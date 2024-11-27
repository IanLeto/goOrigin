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
	Father   string        `json:"father_name"`
	FatherID uint          `json:"father_id"`
	Done     bool          `json:"done"`
	Status   string        `json:"status"`
	Tags     []string      `json:"tags"`
	Note     string        `json:"note"`
	Region   string        `json:"region"`
	Children []string      `json:"children"`
	Nodes    []*NodeEntity `json:"nodes"`
}

func (n *NodeEntity) MergeWith(other *NodeEntity) {
	if other.ID != 0 && n.ID != other.ID {
		n.ID = other.ID
	}
	if other.Name != "" && n.Name != other.Name {
		n.Name = other.Name
	}
	if other.Content != "" && n.Content != other.Content {
		n.Content = other.Content
	}
	if other.Depend != "" && n.Depend != other.Depend {
		n.Depend = other.Depend
	}
	if other.Father != "" && n.Father != other.Father {
		n.Father = other.Father
	}
	if other.FatherID != 0 && n.FatherID != other.FatherID {
		n.FatherID = other.FatherID
	}
	if other.Status != "" && n.Status != other.Status {
		n.Status = other.Status
	}
	if other.Note != "" && n.Note != other.Note {
		n.Note = other.Note
	}
	if other.Region != "" && n.Region != other.Region {
		n.Region = other.Region
	}
	if len(other.Tags) != 0 && !equalStringSlices(n.Tags, other.Tags) {
		n.Tags = other.Tags
	}
	if len(other.Children) != 0 && !equalStringSlices(n.Children, other.Children) {
		n.Children = other.Children
	}
	if other.Nodes != nil && !equalNodeEntitySlices(n.Nodes, other.Nodes) {
		n.Nodes = other.Nodes
	}
}

// Helper function to compare two slices of strings.
func equalStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Helper function to compare two slices of *NodeEntity.
func equalNodeEntitySlices(a, b []*NodeEntity) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		// This is a simple comparison that can be made more complex if needed.
		// For example, you might want to compare IDs or other identifying fields.
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

//func (n *NodeEntity) ToMySQLTable() (mysql.Table, error) {
//
//	table := dao2.TNode{
//		GetName:     n.GetName,
//		Content:  n.Content,
//		Depend:   n.Depend,
//		Father:   n.Father,
//		FatherID: n.FatherID,
//		Done:     n.Done,
//		Status:   n.Status,
//		Region:   n.Region,
//		Note:     n.Note,
//	}
//	return table, nil
//}

//func NewNodeEntityFromTnode(node *dao2.TNode) *NodeEntity {
//	var (
//		tags []string
//	)
//	value, err := json.Marshal(node.Tags)
//	if err != nil {
//		logrus.Errorf("处理tag 失败 %s", err)
//		return nil
//	}
//	err = json.Unmarshal(value, &tags)
//	if err != nil {
//		logrus.Errorf("处理tag 失败 %s", err)
//		return nil
//	}
//	return &NodeEntity{
//		ID:       node.ID,
//		GetName:     node.GetName,
//		Content:  node.Content,
//		Depend:   node.Depend,
//		Father:   node.Father,
//		FatherID: node.FatherID,
//		Done:     node.Done,
//		Status:   node.Status,
//		Tags:     tags,
//		Note:     node.Note,
//		Region:   node.Region,
//	}
//}

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

//func GetNodeAdapter(c *gin.Context, name, father, region string) ([]*NodeEntity, error) {
//	var (
//		db    *gorm.DB
//		res   []*NodeEntity
//		dbRes []*dao2.TNode
//	)
//	tNode := []*dao2.TNode{
//		{GetName: name, Father: father},
//	}
//	db = mysql.NewMysqlConn(config.Conf.Backend.MysqlConfig.Clusters[region]).Client
//	data, _, err := mysql.GetValues(db, tNode, 100)
//	if err != nil {
//		logrus.Errorf("get node failed by %s", err)
//		return nil, err
//	}
//	values, err := json.Marshal(data)
//	if err != nil {
//		logrus.Errorf("get node failed by %s", err)
//		return nil, err
//	}
//	err = json.Unmarshal(values, &dbRes)
//	if err != nil {
//		goto ERR
//	}
//	for _, v := range dbRes {
//		bytes, err := json.Marshal(v.Tags)
//		if err != nil {
//			goto ERR
//		}
//		ephemeralTags := []string{}
//		err = json.Unmarshal(bytes, &ephemeralTags)
//		if err != nil {
//			goto ERR
//		}
//		res = append(res, &NodeEntity{
//			ID:       v.ID,
//			GetName:     v.GetName,
//			Content:  v.Content,
//			Depend:   v.Depend,
//			Father:   v.Father,
//			FatherID: v.FatherID,
//			Done:     v.Done,
//			Status:   v.Status,
//			Tags:     ephemeralTags,
//			Note:     v.Note,
//			Region:   v.Region,
//		})
//	}
//	return res, err
//ERR:
//	return nil, err
//}

//func (node *NodeEntity) CreateNode(c *gin.Context) (id uint, err error) {
//	var (
//		conn   *elastic2.EsV2Conn
//		father *NodeEntity
//		logger = logger2.NewLogger()
//	)
//	conn = elastic2.EsConns[node.Region]
//	_, err = conn.Client.Info()
//	if err != nil {
//		logger.Error(fmt.Sprintf("初始化 es 失败 %s", err))
//		return 0, err
//	}
//	var (
//		query = map[string]interface{}{}
//	)
//	var (
//		doc                   *elastic2.EsDoc
//		insertResultInfo      *elastic2.InsertResultInfo
//		insertResultInfoValue []byte
//
//		insertInfo      map[string]interface{}
//		insertInfoValue []byte
//		value           []byte
//		source          []byte
//	)
//	switch {
//	case node.Father != "":
//		query = map[string]interface{}{
//			"bool": map[string]interface{}{
//				"must": map[string]interface{}{
//					"term": map[string]interface{}{
//						"ToTNode": node.Father,
//					},
//				},
//			},
//		}
//		goto Query
//	case node.FatherID != 0:
//		query = map[string]interface{}{
//			"bool": map[string]interface{}{
//				"must": map[string]interface{}{
//					"term": map[string]interface{}{
//						"_id": node.ID,
//					},
//				},
//			},
//		}
//		goto Query
//	default:
//		goto Create
//	}
//
//Query:
//	logrus.Debugf("query: %s", func() string {
//		b, _ := json.Marshal(query)
//		return string(b)
//	}())
//	value, err = conn.Search(config.NodeMapping, query)
//	if err != nil {
//		goto ERR
//	}
//	err = json.Unmarshal(value, &doc)
//	if err != nil {
//		logrus.Debugf("query: %s", func() string {
//			b, _ := json.Marshal(value)
//			return string(b)
//		}())
//	}
//	if doc.Hits.Total.Value == 0 {
//		err = errors.New("father node not found")
//		goto Create
//	}
//	source, err = json.Marshal(doc.Hits.Hits[0].Source)
//	if err != nil {
//		goto ERR
//	}
//	err = json.Unmarshal(source, &father)
//
//	if err != nil {
//		goto ERR
//	}
//	node.Father = father.GetName
//	node.FatherID = father.ID
//	if source == nil {
//		err = errors.New("father node not found")
//		goto ERR
//	}
//	goto Create
//
//Create:
//	insertInfoValue, err = json.Marshal(node)
//	if err != nil {
//		goto ERR
//	}
//	err = json.Unmarshal(insertInfoValue, &insertInfo)
//	if err != nil {
//		goto ERR
//	}
//	insertResultInfoValue, err = conn.Create("node", insertInfoValue)
//	if err != nil {
//		goto ERR
//	}
//	err = json.Unmarshal(insertResultInfoValue, &insertResultInfo)
//	id, _ = conv.Uint(insertResultInfo.Id)
//	return id, err
//ERR:
//	{
//		return 0, err
//	}
//}
