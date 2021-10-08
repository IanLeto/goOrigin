package model

type Job struct {
	Id        int
	PlanId    int
	PlanName  string
	IsSql     bool
	ProductId int
	RegionIds []int
	SubJob    []SubJob
}

type SubJob struct {
	Id        int
	MainJobId int
	Target    string
	Actions   string
	Rank      string
	ModuleId  int
	State     string
	Type      string
}

type Module struct {
}

// main job
type CreateMainJobReqInfo struct {
	PlanId    int
	PlanName  string
	IsSql     bool
	ProductId int
	RegionIds []int
	SubJob    []CreateSubJobReqInfo
}

type CreateMainJobResInfo struct {
	MainJobId int
}

type QueryMainJobReqInfo struct {
}

type QueryMainJobResInfo struct {
	*Job
}

// sub job
type CreateSubJobReqInfo struct {
	MainJobId int
	ModuleId  int
	Target    string
	Actions   string
	Rank      string
	State     string
	Type      string
}

type QuerySubJobReqInfo struct {
	MainJobId int
	ModuleId  int      // 子任务 module 一对一
	Target    []string // 执行目标
	Rank      []string // 顺序
	State     string   // 状态
	Type      []string // 类型
}

type QuerySubJobResInfo struct {
	*SubJob
}
