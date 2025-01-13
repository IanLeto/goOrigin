package dao

type TTransProject struct {
	*Meta          `json:"*_meta,omitempty"`
	ProjectName    string `json:"project_name"`
	TransDimension string `json:"TransDimension" gorm:"type:varchar(100)"`
	Dimension      string `json:"PredefinedDimensions" gorm:"type:varchar(100)"`
}

type TTransCode struct {
	*Meta       `json:"*_meta,omitempty"`
	TransCode   string `json:"trans_code"`
	TransCodeCn string `json:"trans_code_cn"`
}
