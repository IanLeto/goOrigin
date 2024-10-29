package jobs

import (
	"context"
	"goOrigin/pkg"
)

type Demo struct {
}

// 具体做什么
func (d *Demo) Exec(ctx context.Context, info pkg.JobMessageInfo) error {
	//TODO implement me
	panic("implement me")
}

func (d *Demo) Stop(ctx context.Context, kill chan struct{}) error {
	//TODO implement me
	panic("implement me")
}

func (d Demo) Start() {

}
