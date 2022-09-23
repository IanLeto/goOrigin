package model

import (
	"goOrigin/internal/db"
	"goOrigin/pkg/storage"
	"strings"
)

type Job struct {
	ID         uint
	Targets    []string
	FilePath   string
	Name       string
	Type       string
	StrategyID uint
	Scripts    BaseScript
}

func (j *Job) ToTable() *db.TJob {
	return &db.TJob{
		Name:     j.Name,
		Target:   strings.Join(j.Targets, ","),
		FilePath: j.FilePath,
		Type:     j.Type,
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

func (j *Job) Delete() error {
	return db.DeleteJobByID(j.ID)
}

func (j *Job) QueryDetail() (*db.TJob, error) {
	tJob := j.ToTable()
	err := storage.GlobalMySQL.Model(tJob).First(tJob).Error
	if err != nil {
		return nil, err
	}
	return tJob, err

}
