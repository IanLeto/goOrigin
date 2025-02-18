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
	Cluster       string `json:"cluster"`
	SvcName       string `json:"svc_name"`
	TransType     string `json:"trans_type"`
	TransTypeCode string `json:"trans_type_code"`
	TransTypeDesc string `json:"trans_type_desc"`
	TransChannel  string `json:"trans_channel"`
	RetCode       string `json:"ret_code"`
}

type KafkaLogEntity struct {
	Gid string `json:"ceb.trace.gid,omitempty"`

	Lid string `json:"ceb.trace.lid,omitempty"`

	Pid string `json:"ceb.trace.pid,omitempty"`

	SysName string `json:"sysName,omitempty"`

	Unitcode     string `json:"unitcode,omitempty"`
	InstanceZone string `json:"instanceZone,omitempty"`

	Timestamp string `json:"timestamp,omitempty"`

	LocalApp string `json:"local.app,omitempty"`

	TraceId string `json:"traceId,omitempty"`

	SpanID string `json:"spanId,omitempty"`

	BusinessId string `json:"businessId,omitempty"`

	SpanKind string `json:"span.kind,omitempty"`

	ResultCode string `json:"result.code,omitempty"`

	ThreadName string `json:"current.thread.name,omitempty"`

	TimesCostMs int64 `json:"times.cost.milliseconds,omitempty"`

	LogType string `json:"log.type,omitempty"`

	ContainerPodID string `json:"container.podId,omitempty"`

	Time int64 `json:"time,omitempty"`

	ReqURL string `json:"request.url,omitempty"`

	Method string `json:"method,omitempty"`

	Error string `json:"error,omitempty"`

	ReqSizeBytes int64 `json:"req.size.bytes,omitempty"`

	ReqParam string `json:"req.parameter,omitempty"`

	RespSizeBytes int64 `json:"resp.size.bytes,omitempty"`

	RemoteHost string `json:"remote.host,omitempty"`

	RemotePort string `json:"remote.port,omitempty"`

	SysBaggage string `json:"sys.baggage,omitempty"`

	BizBaggage string `json:"biz.baggage,omitempty"`

	SysExpand map[string]string `json:"sys.expand,omitempty"`

	DbType string `json:"db.type,omitempty"`

	DatabaseName string `json:"database.name,omitempty"`

	Sql string `json:"sql,omitempty"`

	SqlParam string `json:"sql.parameter,omitempty"`

	ConnEstabSpan string `json:"connection.establish.span,omitempty"`

	DbExecCost string `json:"db.execute.cost,omitempty"`

	DatabaseType string `json:"database.type,omitempty"`

	DatabaseEndpoint string `json:"database.endpoint,omitempty"`

	Protocol string `json:"protocol,omitempty"`

	Service string `json:"service,omitempty"`

	MethodParam string `json:"method.parameter,omitempty"`

	InvokeType string `json:"invoke.type,omitempty"`

	RouterRecord string `json:"router.record,omitempty"`

	RemoteIP string `json:"remote.ip,omitempty"`

	LocalClientIP string `json:"local.client.ip,omitempty"`

	ReqSize int64 `json:"req.size,omitempty"`

	RespSize int64 `json:"resp.size,omitempty"`

	ClientElapseTime int64 `json:"client.elapse.time,omitempty"`

	LocalClientPort int64 `json:"local.client.port,omitempty"`

	Baggage string `json:"baggage,omitempty"`

	MessageId     string `json:"msg.id,omitempty"`
	MessageTopic  string `json:"msg.topic,omitempty"`
	PoinMessageId string `json:"poin.msg.id,omitempty"`

	BizImplTime         int64 `json:"biz.impl.time,omitempty"`
	ClientConnTime      int64 `json:"client.conn.time,omitempty"`
	ReqDeserializeTime  int64 `json:"req.deserialize.time,omitempty"`
	ReqSerializeTime    int64 `json:"req.serialize.time,omitempty"`
	RespDeserializeTime int64 `json:"resp.deserialize.time,omitempty"`
	RespSerializeTime   int64 `json:"resp.serialize.time,omitempty"`
	ServerPoolWaitTime  int64 `json:"server.pool.wait.time,omitempty"`

	PhaseTimeCost string `json:"phase.time.cost,omitempty"`

	SpecialTimeMark string `json:"special.time.mark,omitempty"`

	ServerPhaseTimeCost string `json:"server.phase.time.cost,omitempty"`

	ServerSpecialTimeMark string `json:"server.special.time.mark,omitempty"`

	RemotePodId string `json:"remote.podId,omitempty"`

	RemoteApp string `json:"remote.app,omitempty"`

	Trans SpanTransTypeInfoEntity `json:"biz.expand"`
}
