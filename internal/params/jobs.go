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
	ScriptIDs  []string `json:"script_ids"`
}

type UpdateJobRequest struct {
	ID         uint     `json:"id"`
	Name       string   `json:"name"`
	StrategyID uint     `json:"strategy_id"`
	TimeOut    int64    `json:"timeOut"`
	Content    string   `json:"content"`
	FilePath   string   `json:"filePath"`
	Target     []string `json:"target"`
}

type GetJobResponse struct {
	ID         uint     `json:"id"`
	Name       string   `json:"name"`
	StrategyID uint     `json:"strategy_id"`
	TimeOut    int64    `json:"timeOut"`
	Content    string   `json:"content"`
	FilePath   string   `json:"filePath"`
	Target     []string `json:"target"`
}

type GetJobsResponse struct {
	Infos []*GetJobsResponseInfo
}

type GetJobsResponseInfo struct {
	ID         uint     `json:"id"`
	Name       string   `json:"name"`
	StrategyID uint     `json:"strategy_id"`
	TimeOut    int64    `json:"timeOut"`
	Content    string   `json:"content"`
	FilePath   string   `json:"filePath"`
	Target     []string `json:"target"`
	ScriptIDs  []string `json:"script_ids"`
}

type RunJobRequest struct {
	ID         uint     `json:"id"`
	StrategyID uint     `json:"strategy_id"`
	Target     []string `json:"target"`
	ScriptIDs  []string `json:"script_ids"`
}

type RunJobResponse struct {
	ID         uint     `json:"id"`
	StrategyID uint     `json:"strategy_id"`
	Target     []string `json:"target"`
	ScriptIDs  []string `json:"script_ids"`
}
