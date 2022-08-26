package model

import "goOrigin/internal/params"

type Model interface {
	ToParams() (*params.RequestParams, error)
}

type MySqlModel interface {
	Create() error
	Update(id uint) error
	Delete(id uint) error
}

type Job struct {
}
type Strategy struct {
}
