package model

import (
	"goOrigin/internal/params"
	"goOrigin/pkg/storage"
)

type Job struct {
	ID       uint
	Target   string
	FilePath string
}

func NewJob(req params.CreateJobRequest) *Job {
	return &Job{
		ID:       req.ID,
		Target:   req.Target,
		FilePath: req.FilePath,
	}
}

func (j *Job) Create() error {
	return storage.GlobalMySQL.Create(j).Error
}
