package entity

import "github.com/prometheus/client_golang/prometheus"

type RecordEntity struct {
	Name       string  `json:"name" bson:"name"`
	Weight     float32 `json:"weight" bson:"weight"`
	IsFuck     bool    `json:"is_fuck"`
	Vol1       string  `json:"vol1" bson:"vol1"`
	Vol2       string  `json:"vol2" bson:"vol2"`
	Vol3       string  `json:"vol3" bson:"vol3"`
	Vol4       string  `json:"vol4" bson:"vol4"`
	Cost       int     `json:"cost" bson:"cost"`
	Content    string  `json:"content" bson:"content"`
	Region     string  `json:"region" bson:"region"`
	Retire     int     `json:"retire" bson:"retire"`
	Dev        string  `json:"dev"`
	Coding     string  `json:"coding"`
	Reading    string  `json:"reading"`
	Social     string  `json:"social"`
	CreateTime int64   `json:"create_time"`
	ModifyTime int64   `json:"update_time"`
}

type RecordMetricEntity struct {
	Weight *prometheus.CounterVec
}

type RecordMetricMessageEntity struct {
	Name   string
	Weight *prometheus.GaugeVec
	IsFuck *prometheus.GaugeVec
	Cost   *prometheus.GaugeVec

	CreateTime *prometheus.GaugeVec
	ModifyTime *prometheus.GaugeVec
}
