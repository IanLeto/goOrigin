package model

type QueryPlanRequestInfo struct {
	AccessIds  []int
	ProjectIds []int
	RegionIds  []int
	TemplateId int
}

type QueryPlanResponseInfo struct {
}

type PlanResponseInfo struct {
	PlanId int
}

type CreatePlanRequestInfo struct {
	AccessId   int
	ProjectId  int
	RegionIds  []int
	TemplateId int
}

type CreateMainJobRequestInfo struct {
	*Job
}
