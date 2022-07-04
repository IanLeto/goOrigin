package model

import "goOrigin/internal/params"

type Ian struct {
	Body   `json:"body"`
	BETre  `json:"BETre"`
	Worker `json:"worker"`
}

type Body struct {
	Weight int
	Code   string
}
type BETre struct {
	Core    int `json:"core"`
	Runner  int `json:"runner"`
	Support int `json:"support"`
}

type Worker struct {
}

func NewIan(req params.CreateIanRequestInfo) *Ian {
	return &Ian{}
}

func (root *Ian) Save() {

}

func (root *Ian) Update() {

}
