package model

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"goOrigin/internal/params"
)

type Base struct {
	Id         primitive.ObjectID  `json:"id" bson:"_id,omitempty"`
	CreateDate primitive.Timestamp `json:"create_date" bson:"create_date"`
	UpdateDate primitive.Timestamp `json:"update_time" bson:"update_date"`
}

type Ian struct {
	*Base
	Name string `json:"name" bson:"name"`
	Time struct {
		T int64 `json:"T"`
		I int64 `json:"I"`
	} `json:"time"`

	Body   `json:"body" bson:"body"`
	BETre  `json:"BETre" bson:"BETre"`
	Worker `json:"worker" bson:"worker"`
}

type Body struct {
	Weight float32 `json:"weight" bson:"weight"`
}
type BETre struct {
	Core       int `json:"core" bson:"core"`
	Runner     int `json:"runner" bson:"runner"`
	Support    int `json:"support" bson:"support"`
	Squat      int `json:"squat" bson:"squat"`
	EasyBurpee int `json:"easy_burpee" bson:"easyBurpee"`
	Chair      int `json:"chair" bson:"chair"`
	Stretch    int `json:"stretch" bson:"stretch"`
}

type Worker struct {
	Vol1 string `json:"vol1" bson:"vol1"`
	Vol2 string `json:"vol2" bson:"vol2"`
	Vol3 string `json:"vol3" bson:"vol3"`
	Vol4 string `json:"vol4" bson:"vol4"`
}

func (i *Ian) ToString() string {
	data, _ := json.Marshal(i)
	return string(data)
}

func NewIan(req params.CreateIanRequestInfo) *Ian {
	return &Ian{
		Name: req.Name,
		Base: &Base{},
		Body: Body{
			Weight: req.Body.Weight,
		},
		BETre: BETre{
			Core:       req.BETre.Core,
			Runner:     req.BETre.Runner,
			Support:    req.BETre.Support,
			Squat:      req.BETre.Squat,
			EasyBurpee: req.BETre.EasyBurpee,
			Chair:      req.BETre.Chair,
			Stretch:    req.BETre.Stretch,
		},
		Worker: Worker{
			Vol1: req.Worker.Vol1,
			Vol2: req.Worker.Vol2,
			Vol3: req.Worker.Vol3,
			Vol4: req.Worker.Vol4,
		},
	}
}

func (root *Ian) Save() {

}

func (root *Ian) Update() {

}
