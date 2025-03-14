package entity

type TraceInfoEntity struct {
	Orgsys      string `json:"ceb.biz.orgsys"`
	Cursys      string `json:"ceb.biz.cursys"`
	ServiceCode string `json:"servicecode"`
	ReturnCode  string `json:"returnCode"`
	TraceId     string `json:"trace_id"`
	SpanId      string `json:"span_id"`
}
