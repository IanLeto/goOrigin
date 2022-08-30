package model

import (
	"goOrigin/internal/db"
	"goOrigin/internal/params"
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
	tJob := j.ToTable()
	_, err := tJob.Create()
	if err != nil {
		return err
	}
	j.ID = tJob.ID
	return nil
}

func (j *Job) Update() error {
	tJob := j.ToTable()
	_, err := tJob.Update()
	if err != nil {
		return err
	}
	j.ID = tJob.ID
	return nil
}
