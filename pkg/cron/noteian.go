package cron

import (
	"context"
	"goOrigin/pkg"
	"time"
)

type NoteIan struct {
	Trick *time.Ticker
}

func (n *NoteIan) Exec(ctx context.Context, info pkg.JobMessageInfo) error {
	//TODO implement me
	panic("implement me")
}

func (n *NoteIan) Stop(ctx context.Context, kill chan struct{}) error {
	//TODO implement me
	panic("implement me")
}

func RegisterNoteIan() error {
	task := &NoteIan{
		time.NewTicker(24 * time.Hour),
	}
	QueueCron = append(QueueCron, task)
	return nil
}
