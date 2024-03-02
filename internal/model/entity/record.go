package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Record struct {
	ID         primitive.ObjectID  `json:"id" bson:"_id,omitempty"`
	CreateDate primitive.Timestamp `json:"create_date" bson:"create_date"`
	UpdateDate primitive.Timestamp `json:"update_time" bson:"update_date"`
	Name       string              `json:"name" bson:"name"`
	Weight     float32             `json:"weight" bson:"weight"`
	IsFuck     bool                `json:"is_fuck"`
	Vol1       string              `json:"vol1" bson:"vol1"`
	Vol2       string              `json:"vol2" bson:"vol2"`
	Vol3       string              `json:"vol3" bson:"vol3"`
	Vol4       string              `json:"vol4" bson:"vol4"`
	Cost       int                 `json:"cost" bson:"cost"`
	Content    string              `json:"content" bson:"content"`
	Region     string              `json:"region" bson:"region"`
	Retire     int                 `json:"retire" bson:"retire"`
}
