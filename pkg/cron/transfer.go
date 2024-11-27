package cron

import "context"

type Transfer struct {
	Id string
}

func (t *Transfer) Exec(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (t *Transfer) GetName() string {
	//TODO implement me
	panic("implement me")
}

func TransferCornFactory() error {
	GTM.AddJob(&Transfer{Id: ""})
	return nil
}
