package db

type meta struct {
	ID         uint  `swaggerignore:"true" gorm:"primary_key" json:"id" binding:"-" `
	CreateTime int64 `swaggerignore:"true" gorm:"autoCreateTime;" json:"created_time" binding:"-"`
	ModifyTime int64 `swaggerignore:"true" gorm:"autoUpdateTime;" json:"modify_time" binding:"-"`
	//DeleteTime soft_delete.DeletedAt `swaggerignore:"true" json:"delete_time" binding:"-"`
}

type TJob struct {
	*meta
	Name     string `json:"name"`
	Target   string `json:"target"`
	FilePath string `json:"file_path"`
}
