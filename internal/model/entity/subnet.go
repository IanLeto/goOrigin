package entity

// SubnetEntity 子网详情的统一结构
type SubnetEntity struct {
	// 子网标识
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	CIDR string `json:"cidr"` // 如 "10.244.0.0/16"
	// 统计信息（可选）
	TotalIPs    int     `json:"totalIPs,omitempty"`    // 该 CIDR 理论可用 IP 数（去除网络/广播可由服务端计算后填充）
	UsedIPs     int     `json:"usedIPs,omitempty"`     // 已分配 IP 数
	FreeIPs     int     `json:"freeIPs,omitempty"`     // 可分配 IP 数
	Utilization float64 `json:"utilization,omitempty"` // 使用率 0~1
	// 关联与来源
	Scope       string            `json:"scope,omitempty"`     // 作用域说明，如 "cluster", "namespace", "node"
	Namespace   string            `json:"namespace,omitempty"` // 若为命名空间级子网则填充
	NodeName    string            `json:"nodeName,omitempty"`  // 若为节点级（PodCIDR per node）可填
	CNI         string            `json:"cni,omitempty"`       // 如 calico, cilium, flannel
	Labels      map[string]string `json:"labels,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
}
