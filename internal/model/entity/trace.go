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
	TraceId string `json:"traceId,omitempty"`

	SpanID string `json:"spanId,omitempty"`

	Trans SpanTransTypeInfoEntity `json:"biz.expand"`
}
