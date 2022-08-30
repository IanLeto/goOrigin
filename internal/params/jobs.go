package params

type CreateJobRequest struct {
	ID         uint     `json:"id"`
	Name       string   `json:"name"`
	Type       string   `json:"type"`
	StrategyID uint     `json:"strategy_id"`
	TimeOut    int64    `json:"timeOut"`
	Content    string   `json:"content"`
	FilePath   string   `json:"filePath"`
	Targets    []string `json:"targets"`
}

type UpdateJobRequest struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	StrategyID uint   `json:"strategy_id"`
	TimeOut    int64  `json:"timeOut"`
	Content    string `json:"content"`
	FilePath   string `json:"filePath"`
	Target     string `json:"target"`
}
