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
	ScriptIDS  []string
}

func (j *Job) ToTable() *db.TJob {
	return &db.TJob{
		Name:      j.Name,
		Target:    strings.Join(j.Targets, ","),
		FilePath:  j.FilePath,
		Type:      j.Type,
		ScriptIDs: strings.Join(j.ScriptIDS, ","),
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

//func QueryList(j []*Job) (*db.TJob, error) {
//	var tJobs []*db.TJob
//	for _, job := range j {
//		tJobs = append(tJobs, job.ToTable())
//	}
//	err := storage.GlobalMySQL.Model(&db.TJob{}).Find(tJobs)
//	if err != nil {
//		return nil, err
//	}
//	return tJob, err
//
//}
