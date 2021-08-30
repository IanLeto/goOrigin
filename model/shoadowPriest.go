package model

import "time"

type IanUI struct {
	Id         string      `json:"id"`
	Weight     int         `json:"weight"`
	CodingLine int         `json:"coding_line"`
	Task       []DailyTask `json:"task"`
	GetUp      time.Time   `json:"get_up"`
	Sleep      time.Time   `json:"sleep"`
	Cost       int         `json:"cost"`
	Deposit    int         `json:"deposit"`
	DailyTask
	HavingFun
}

type DailyTask struct {
	Expected string
	Actual   string
}

type HavingFun struct {
	WithWho string
}
