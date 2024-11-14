package entity

type TraceEntity struct {
	Gid               string `json:"gid"`                    // gid
	Cid               string `json:"cid"`                    // cid
	Pid               string `json:"pid"`                    // pid
	SystemName        string `json:"system_name"`            // system_name
	Unitcode          string `json:"unitcode"`               // unitcode
	InstanceZone      string `json:"instance_zone"`          // instance_zone
	Timestamp         string `json:"timestamp"`              // 采集时间
	LocalApp          string `json:"local.app"`              // 本地应用
	TraceId           string `json:"traceId"`                // trace ip
	SpanId            string `json:"spanId"`                 // span id
	BusinessId        string `json:"businessId"`             // 业务id
	SpanKind          string `json:"spanKind"`               // span kind
	ResultCode        string `json:"result.code"`            // result code
	CurrentThreadName string `json:"current.thread.name"`    // current thread name
	Cost              string `json:"time.cost.milliseconds"` // cost
	ReqUrl            string `json:"req.url"`                // req url
	ReqMethod         string `json:"method"`                 // req method
	Error             string `json:"error"`                  // error
	ReqSize           string `json:"req.size.bytes"`         // req size
	RespSize          string `json:"resp.size.bytes"`        // resp size
	RemoteHost        string `json:"remote.host"`            // remote host
	RemotePort        string `json:"remote.port"`            // remote port
	SysBaggage        string `json:"sys.baggage"`            // sys baggage
	BizBaggage        string `json:"biz.baggage"`            // biz baggage

}

type SpanInfoEntity struct {
	TraceId       string                  `json:"traceId"`       // trace ip
	SpanId        string                  `json:"spanId"`        // span id
	Ctimestamp    string                  `json:"ctimestamp"`    // ctimestamp
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
	Src           string `json:"src"`
	Psrc          string `json:"psrc"`
	TransType     string `json:"trans_type"`
	TransTypeCode string `json:"trans_type_code"`
	TransTypeDesc string `json:"trans_type_desc"`
	TransChannel  string `json:"trans_channel"`
	RetCode       string `json:"ret_code"`
}

type SpanInfoDocEntity struct {
}
