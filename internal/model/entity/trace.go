package entity

type TraceEntity struct {
	Gid               string `json:"gid"`
	Cid               string `json:"cid"`
	Pid               string `json:"pid"`
	SystemName        string `json:"system_name"`
	Unitcode          string `json:"unitcode"`
	InstanceZone      string `json:"instance_zone"`
	Timestamp         string `json:"timestamp"`
	LocalApp          string `json:"local.app"`
	TraceId           string `json:"traceId"`
	SpanId            string `json:"spanId"`
	BusinessId        string `json:"businessId"`
	SpanKind          string `json:"spanKind"`
	ResultCode        string `json:"result.code"`
	CurrentThreadName string `json:"current.thread.name"`
	Cost              string `json:"time.cost.milliseconds"`
	ReqUrl            string `json:"req.url"`
	ReqMethod         string `json:"method"`
	Error             string `json:"error"`
	ReqSize           string `json:"req.size.bytes"`
	RespSize          string `json:"resp.size.bytes"`
	RemoteHost        string `json:"remote.host"`
	RemotePort        string `json:"remote.port"`
	SysBaggage        string `json:"sys.baggage"`
	BizBaggage        string `json:"biz.baggage"`
}

type SpanInfoEntity struct {
	TraceId       string                  `json:"traceId"`
	SpanId        string                  `json:"spanId"`
	Ctimestamp    string                  `json:"ctimestamp"`
	InstanceZone  string                  `json:"instance_zone"` // instance_zone
	ContainerInfo SpanContainerInfoEntity `json:"container"`
}

type SpanContainerInfoEntity struct {
	PodID         string `json:"pod_id"`         // pod_id
	ContainerID   string `json:"container_id"`   // container_id
	ContainerName string `json:"container_name"` // container_name
	NameSpace     string `json:"namespace"`      // namespace
	AZ            string `json:"az"`             // az
	Cluster       string `json:"cluster"`        // cluster
	Project       string `json:"project"`        // project
	ResName       string `json:"resName"`        // res_name
	ResKind       string `json:"resKind"`        // res_kind
	Stack         string `json:"stack"`          // stack
	Application   string `json:"application"`    // application
}

type SpanTransTypeInfoEntity struct {
	TransType     string `json:"trans_type"`
	TransTypeCode string `json:"trans_type_code"`
	TransChannel  string `json:"trans_channel"`
	RetCode       string `json:"ret_code"`
}

type KafkaLogEntity struct {
	// === 链路标识 ===
	Gid string `json:"ceb.trace.gid,omitempty"`
	Lid string `json:"ceb.trace.lid,omitempty"`
	Pid string `json:"ceb.trace.pid,omitempty"`

	// === 系统部署信息 ===
	SysName         string `json:"sysName,omitempty"`           // 系统名
	Unitcode        string `json:"unitcode,omitempty"`          // 单元编号
	InstanceZone    string `json:"instanceZone,omitempty"`      // 实例区域
	LocalApp        string `json:"local.app,omitempty"`         // 本地服务名
	LocalClientIP   string `json:"local.client.ip,omitempty"`   // 本地客户端 IP
	LocalClientPort int64  `json:"local.client.port,omitempty"` // 本地客户端端口

	// === 链路追踪与线程信息 ===
	TraceId    string `json:"traceId,omitempty"`             // 链路 Trace ID
	SpanID     string `json:"spanId,omitempty"`              // 当前调用 Span ID
	SpanKind   string `json:"span.kind,omitempty"`           // Span 类型（client/server）
	ThreadName string `json:"current.thread.name,omitempty"` // 当前线程名

	// === 时间信息 ===
	Timestamp   string `json:"timestamp,omitempty"`               // 原始时间戳字符串
	Time        int64  `json:"time,omitempty"`                    // Unix 时间戳（毫秒）
	TimesCostMs int64  `json:"times.cost.milliseconds,omitempty"` // 总耗时

	// === 业务标识 ===
	BusinessId string `json:"businessId,omitempty"` // 业务 ID

	// === 请求信息 ===
	ReqURL             string `json:"request.url,omitempty"`          // 请求 URL
	Method             string `json:"method,omitempty"`               // 请求方法
	ReqParam           string `json:"req.parameter,omitempty"`        // 请求参数
	ReqSizeBytes       int64  `json:"req.size.bytes,omitempty"`       // 请求字节大小
	ReqSize            int64  `json:"req.size,omitempty"`             // 请求大小（冗余字段）
	ReqSerializeTime   int64  `json:"req.serialize.time,omitempty"`   // 请求序列化时间
	ReqDeserializeTime int64  `json:"req.deserialize.time,omitempty"` // 请求反序列化时间

	// === 响应信息 ===
	RespSizeBytes       int64  `json:"resp.size.bytes,omitempty"`       // 响应字节大小
	RespSize            int64  `json:"resp.size,omitempty"`             // 响应大小（冗余字段）
	RespSerializeTime   int64  `json:"resp.serialize.time,omitempty"`   // 响应序列化时间
	RespDeserializeTime int64  `json:"resp.deserialize.time,omitempty"` // 响应反序列化时间
	ResultCode          string `json:"result.code,omitempty"`           // 响应结果码
	Error               string `json:"error,omitempty"`                 // 错误信息

	// === 数据库信息 ===
	DbType           string `json:"db.type,omitempty"`                   // 数据库类型
	DatabaseType     string `json:"database.type,omitempty"`             // 数据库类型（冗余）
	DatabaseName     string `json:"database.name,omitempty"`             // 数据库名称
	DatabaseEndpoint string `json:"database.endpoint,omitempty"`         // 数据库连接地址
	Sql              string `json:"sql,omitempty"`                       // 执行 SQL
	SqlParam         string `json:"sql.parameter,omitempty"`             // SQL 参数
	ConnEstabSpan    string `json:"connection.establish.span,omitempty"` // 连接建立耗时
	DbExecCost       string `json:"db.execute.cost,omitempty"`           // SQL 执行耗时

	// === 网络信息 ===
	RemoteHost  string `json:"remote.host,omitempty"`  // 远端主机
	RemotePort  string `json:"remote.port,omitempty"`  // 远端端口
	RemoteIP    string `json:"remote.ip,omitempty"`    // 远端 IP
	RemotePodId string `json:"remote.podId,omitempty"` // 远端 Pod ID
	RemoteApp   string `json:"remote.app,omitempty"`   // 远端 App 名称

	// === 协议与服务 ===
	Protocol    string `json:"protocol,omitempty"`         // 通信协议
	Service     string `json:"service,omitempty"`          // 服务名
	MethodParam string `json:"method.parameter,omitempty"` // 方法参数
	InvokeType  string `json:"invoke.type,omitempty"`      // 调用类型（sync/async）

	// === 调用性能 ===
	ClientElapseTime   int64 `json:"client.elapse.time,omitempty"`    // 客户端耗时
	ClientConnTime     int64 `json:"client.conn.time,omitempty"`      // 客户端连接时间
	BizImplTime        int64 `json:"biz.impl.time,omitempty"`         // 业务实现耗时
	ServerPoolWaitTime int64 `json:"server.pool.wait.time,omitempty"` // 线程池等待时间

	// === 时间阶段分析 ===
	PhaseTimeCost         string `json:"phase.time.cost,omitempty"`          // 阶段耗时
	SpecialTimeMark       string `json:"special.time.mark,omitempty"`        // 特殊时间标记
	ServerPhaseTimeCost   string `json:"server.phase.time.cost,omitempty"`   // Server 阶段耗时
	ServerSpecialTimeMark string `json:"server.special.time.mark,omitempty"` // Server 特殊时间标记

	// === 日志与消息 ===
	LogType       string `json:"log.type,omitempty"`    // 日志类型
	MessageId     string `json:"msg.id,omitempty"`      // 消息 ID
	MessageTopic  string `json:"msg.topic,omitempty"`   // 消息 Topic
	PoinMessageId string `json:"poin.msg.id,omitempty"` // 指向消息 ID

	// === Baggage 扩展 ===
	SysBaggage string            `json:"sys.baggage,omitempty"` // 系统透传字段
	BizBaggage string            `json:"biz.baggage,omitempty"` // 业务透传字段
	Baggage    string            `json:"baggage,omitempty"`     // 通用透传字段
	SysExpand  map[string]string `json:"sys.expand,omitempty"`  // 系统扩展字段

	// === 自定义业务扩展 ===
	RouterRecord string                  `json:"router.record,omitempty"` // 路由记录
	Trans        SpanTransTypeInfoEntity `json:"biz.expand"`              // 跨服务扩展信息
	ReturnCode   string                  `json:"return_code"`
}
