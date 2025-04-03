package dao

type TNode struct {
	*Meta

	Name      string `gorm:"type:varchar(255);not null;comment:节点名称"`
	Content   string `gorm:"type:text;comment:节点详细描述"`
	ParentID  uint   `gorm:"type:int unsigned;index;comment:父节点 ID，0 表示根节点"`
	DependIDs string `gorm:"type:text;comment:依赖的其他节点 ID，JSON 数组存储"`
	Done      bool   `gorm:"type:tinyint(1);default:0;comment:是否完成"`
	Status    string `gorm:"type:varchar(64);default:'pending';comment:状态"`
	Region    string `gorm:"type:varchar(64);comment:区域/环境"`
	Note      string `gorm:"type:text;comment:备注信息"`
	Tags      string `gorm:"type:text;comment:标签，JSON 数组存储"`
}
