package model

type ShadowPriest struct {
	Id         int         `json:"id"`
	Weight     int         `json:"weight"`
	CodingLine int         `json:"coding_line"`
	Task       []DailyTask `json:"task"`
	Cost       int         `json:"cost"`
	Deposit    int         `json:"deposit"`
	DailyTask
	HavingFun
}

type DailyTask struct {
	Expected string `json:"expected"`
	Actual   string `json:"actual"`
}

type HavingFun struct {
	WithWho string `json:"with_who"`
}
