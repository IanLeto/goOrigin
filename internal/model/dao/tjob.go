package dao

type TJob struct {
	Name      string `json:"name" gorm:"name"`
	Target    string `json:"target" gorm:"target"`
	FilePath  string `json:"file_path" gorm:"file_path"`
	Type      string `json:"type" gorm:"type"`
	ScriptIDs string `json:"script_ids" gorm:"script_ids"`
}

//func (t *TJob) Create() (mysql.DBOpt, error) {
//	return nil, storage.GlobalMySQL.Create(t).Error
//}
//
//func (t *TJob) Update() (mysql.DBOpt, error) {
//	return nil, storage.GlobalMySQL.Model(t).Update(t).Error
//}
//
//func (t *TJob) Delete() (mysql.DBOpt, error) {
//	return nil, storage.GlobalMySQL.Delete(t).Error
//}
//
//func (t *TJob) List() (mysql.DBOpt, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (t *TJob) BatchCreate(dbOpts []mysql.DBOpt) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (t *TJob) BatchUpdate(opts []mysql.DBOpt) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func DeleteJobByID(id uint) error {
//	return storage.GlobalMySQL.Table("t_jobs").Delete(&TJob{}, id).Error
//}
