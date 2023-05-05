package params

type CreateRoutineReq struct {
	Content string `json:"content"`
	Extra   string `json:"extra"`
	Status  string `json:"status"`
}

type CreateRoutine struct {
	ID string `json:"ID"`
}

type QueryRoutineReq struct {
	Date string `json:"date"`
}

type QueryRoutineRes struct {
	Content string `json:"content"`
	Extra   string `json:"extra"`
	Status  string `json:"status"`
	Date    string `json:"date"`
}
