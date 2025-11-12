package V1

// 通用基础类型
type IPFamily string

type SubnetType string

// 子网详情的统一结构
type Subnet struct {
	// 子网标识
	ID       string     `json:"id,omitempty"`
	Name     string     `json:"name,omitempty"`
	Type     SubnetType `json:"type,omitempty"`     // Pod / Service / Node / Custom
	CIDR     string     `json:"cidr"`               // 如 "10.244.0.0/16"
	IPFamily IPFamily   `json:"ipFamily,omitempty"` // IPv4/IPv6/DualStack
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

// 错误与分页（可选）
type Page struct {
	Page     int `json:"page,omitempty"`
	PageSize int `json:"pageSize,omitempty"`
}

type PageResult struct {
	Total    int `json:"total"`
	Page     int `json:"page,omitempty"`
	PageSize int `json:"pageSize,omitempty"`
}

// ============ 1) 集群子网查询 ============

type QueryClusterSubnetRequest struct {
	ClusterID  string            `json:"clusterID"`            // 必填：集群标识
	Types      []SubnetType      `json:"types,omitempty"`      // 过滤：仅返回某些类型的子网（Pod/Service/Node/Custom）
	IPFamilies []IPFamily        `json:"ipFamilies,omitempty"` // 过滤：IPv4/IPv6/DualStack
	CNI        string            `json:"cni,omitempty"`        // 过滤：CNI 名称
	WithStats  bool              `json:"withStats,omitempty"`  // 是否计算使用统计
	Page       *Page             `json:"page,omitempty"`       // 可选分页
	Labels     map[string]string `json:"labels,omitempty"`     // 过滤：匹配子网标签
}

type QueryClusterSubnetResponse struct {
	Subnets    []Subnet    `json:"subnets"`
	Pagination *PageResult `json:"pagination,omitempty"`
}

// ============ 2) 命名空间子网查询 ============

type QueryNsSubnetRequest struct {
	ClusterID  string            `json:"clusterID"`            // 必填
	Namespace  string            `json:"namespace"`            // 必填：目标命名空间
	Types      []SubnetType      `json:"types,omitempty"`      // 过滤
	IPFamilies []IPFamily        `json:"ipFamilies,omitempty"` // 过滤
	WithStats  bool              `json:"withStats,omitempty"`  // 是否计算使用统计
	Page       *Page             `json:"page,omitempty"`
	Labels     map[string]string `json:"labels,omitempty"` // 过滤
}

type QueryNsSubnetResponse struct {
	Namespace  string      `json:"namespace"`
	Subnets    []Subnet    `json:"subnets"`
	Pagination *PageResult `json:"pagination,omitempty"`
}

// ============ 3) 空闲子网查询 ============

type QueryAvailableSubnetRequest struct {
	ClusterID      string       `json:"clusterID"`                // 必填
	PreferredTypes []SubnetType `json:"preferredTypes,omitempty"` // 希望的子网类型（例如仅 Pod 子网）
	IPFamilies     []IPFamily   `json:"ipFamilies,omitempty"`     // IPv4/IPv6/DualStack
	// 容量需求：至少需要可用 IP 数（用于容量规划/分配）
	MinFreeIPs int `json:"minFreeIPs,omitempty"`
	// 附加筛选
	Namespace string            `json:"namespace,omitempty"` // 指定命名空间可用子网（若实现支持 namespace-scoped 子网）
	CNI       string            `json:"cni,omitempty"`
	Labels    map[string]string `json:"labels,omitempty"`
	Page      *Page             `json:"page,omitempty"`
	// 是否返回排序（例如按可用 IP 数降序）
	SortBy    string `json:"sortBy,omitempty"` // "freeIPs" | "utilization" | "name"
	SortDesc  bool   `json:"sortDesc,omitempty"`
	WithStats bool   `json:"withStats,omitempty"` // 通常空闲子网查询需要统计，默认 true 亦可
}

type QueryAvailableSubnetResponse struct {
	Subnets    []Subnet    `json:"subnets"` // 满足过滤条件且有可用容量的子网
	Pagination *PageResult `json:"pagination,omitempty"`
}
