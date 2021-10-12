package model

type AddTemplateReq struct {
	ServerName string `json:"server_name"` //server_name 服务名 必填
	Version    string `json:"version"`     //version  版本
	Content    string `json:"content"`     //content  模板内容
	Tag        string `json:"tag"`         //tag  模板唯一标示 默认为 服务名+版本名+zone+set
	Check      bool   `json:"check"`       // check
	Config     string `json:"config"`      //config 配置文件地址
	SetId      string `json:"set_id"`      //set_id  set id
	ZoneId     string `json:"zone_id"`     //zone_id  zone id
	Type       string `json:"type"`        //type  类型 default 为基础模板
}

