package V1

type CreateRecordReqInfo struct {
	Date   int `json:"date"`
	Weight int `json:"weight"`
}

type CreatRecordResInfo struct {
	*BaseResponseInfo
}
