package model

const EsTopo = "topo"     // 知识拓扑
const EsNode = "node"     // 拓扑节点
const EsIanRecord = "ian" // ian 状态记录
const EsScript = "script" // 脚本

type Entity interface {
	ToMySQLTable() (interface{}, error)
}
