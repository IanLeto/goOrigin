package V1

type GetResourceEventReq struct {
	*GetResourceEventReqInfo
}

type GetResourceEventReqInfo struct {
	Cluster   string `json:"cluster"`
	Namespace string `json:"namespace"`
	Kind      string `json:"kind"`
	Name      string `json:"name"`
	From      string `json:"from"`
	To        string `json:"to"`
}
