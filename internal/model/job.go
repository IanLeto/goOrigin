package model

import (
	"goOrigin/internal/db"
	"goOrigin/internal/params"
	"goOrigin/pkg/storage"
)

type Job struct {
	ID       uint
	Target   string
	FilePath string
	Name     string
}

func NewJob(req params.CreateJobRequest) *Job {
	return &Job{
		ID:       req.ID,
		Target:   req.Target,
		Name:     req.Name,
		FilePath: req.FilePath,
	}
}

func (j *Job) ToTable() *db.TJob {
	return &db.TJob{
		Name:     j.Name,
		Target:   j.Target,
		FilePath: j.FilePath,
	}
}
func (j *Job) Create() error {

	return storage.GlobalMySQL.Create(j.ToTable()).Error
}
