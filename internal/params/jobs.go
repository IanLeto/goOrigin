package params

import "goOrigin/internal/model"

type RequestParams interface {
	ToModel() (*model.Model, error)
}

type CreateJobRequest struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	StrategyID uint   `json:"strategy_id"`
	TimeOut    int64  `json:"timeOut"`
	Content    string `json:"content"`
}
