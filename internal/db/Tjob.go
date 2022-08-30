package db

import "goOrigin/pkg/storage"

type TJob struct {
	*Meta
	Name     string `json:"name"`
	Target   string `json:"target"`
	FilePath string `json:"file_path"`
}

func (t *TJob) Create() (DBOpt, error) {
	return nil, storage.GlobalMySQL.Create(t).Error
}

func (t *TJob) Update() (DBOpt, error) {
	return nil, storage.GlobalMySQL.Update(t).Error
}

func (t *TJob) Delete() (DBOpt, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TJob) List() (DBOpt, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TJob) BatchCreate(dbOpts []DBOpt) error {
	//TODO implement me
	panic("implement me")
}

func (t *TJob) BatchUpdate(opts []DBOpt) {
	//TODO implement me
	panic("implement me")
}
