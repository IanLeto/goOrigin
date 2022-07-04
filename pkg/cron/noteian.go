package cron

import (
	"context"
	"goOrigin/internal/model"
	"goOrigin/pkg"
	"time"
)

type NoteIan struct {
	Trick *time.Ticker
}

func (n *NoteIan) Exec(ctx context.Context, info pkg.JobMessageInfo) error {
	var (
		err error
	)
	for {
		select {
		case <-n.Trick.C:
			root := model.Ian{
				Body:   model.Body{},
				BETre:  model.BETre{},
				Worker: model.Worker{},
			}
			root.Save()
		case <-ctx.Done():
			return err
		}

	}

}

func (n *NoteIan) Stop(ctx context.Context, kill chan struct{}) error {
	//TODO implement me
	panic("implement me")
}

func RegisterNoteIan() error {
	task := &NoteIan{}
	QueueCron = append(QueueCron, task)
	return nil
}
