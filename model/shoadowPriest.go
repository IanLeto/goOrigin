package model

import "time"

type IanUI struct {
	Id         string
	Weight     int
	CodingLine int
	Task       []DailyTask
	GetUp      time.Time
	Sleep      time.Time
	Cost       int
	Deposit    int
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
