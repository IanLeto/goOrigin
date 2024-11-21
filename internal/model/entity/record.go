package entity

import "github.com/prometheus/client_golang/prometheus"

type Record struct {
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
	Name       *prometheus.CounterVec
	Weight     *prometheus.GaugeVec
	IsFuck     *prometheus.GaugeVec
	Vol1       *prometheus.CounterVec
	Vol2       *prometheus.CounterVec
	Vol3       *prometheus.CounterVec
	Vol4       *prometheus.CounterVec
	Cost       *prometheus.GaugeVec
	Content    *prometheus.CounterVec
	Region     *prometheus.CounterVec
	Retire     *prometheus.GaugeVec
	Dev        *prometheus.CounterVec
	Coding     *prometheus.CounterVec
	Reading    *prometheus.CounterVec
	Social     *prometheus.CounterVec
	CreateTime *prometheus.GaugeVec
	ModifyTime *prometheus.GaugeVec
}
