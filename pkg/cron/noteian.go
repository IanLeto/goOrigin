package cron

import (
	"context"
	"github.com/sirupsen/logrus"
	"goOrigin/internal/model"
	"goOrigin/pkg"
	"goOrigin/pkg/storage"
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
			_, err := storage.GlobalRedisCon.Client.LPush("noteian", "1").Result()
			if err != nil {
				logrus.Errorf("error in reging redis %s", err)
			}
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
	task := &NoteIan{
		time.NewTicker(24 * time.Hour),
	}
	QueueCron = append(QueueCron, task)
	return nil
}
