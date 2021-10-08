package opsHandlers

import "github.com/gin-gonic/gin"

// @Summary 创建发布计划
// @Description
// @Tags plan
// @Accept json
// @Param plan body model.CreatePlanRequestInfo true "计划"
// @Success 200  object model.PlanResponseInfo ""
// @Router /v1/ops/plan [post]
func Plan(c *gin.Context) {

}

// @Summary 查看发布计划详情
// @Description
// @Tags plan
// @Accept json
// @Param plan query model.QueryPlanRequestInfo false "状态"
// @Success 200  object model.QueryPlanResponseInfo ""
// @Router /v1/ops/plan [get]
func FindPlan(c *gin.Context) {

}

// @Summary 删除发布计划
// @Description 不支持批量删除
// @Tags plan
// @Accept json
// @Param plan_id query int false "plan id"
// @Success 200 {string} string "{
//    "retcode":0,
//    "message":"",
//    "data":{
//			"plan_id": 1
//    }
//}"
// @Router /v1/ops/{plan_id} [delete]
func DeletePlan() {

}

// @Summary 执行发布计划
// @Description
// @Tags plan
// @Param plan_id query int false "plan id"
// @Success 200 {string} string "{
//    "retcode":0,
//    "message":"",
//    "data":{
//			"plan_id": 1
//    }
//}"
// @Router /v1/ops/exec/{plan_id} [get]
func ExecPlan() {

}

// @Summary 执行发布计划会滚
// @Description
// @Tags plan
// @Param plan_id query int false "plan id"
// @Success 200 {string} string "{
//    "retcode":0,
//    "message":"",
//    "data":{
//			"plan_id": 1
//    }
//}"
// @Router /v1/ops/exec/{version_id} [get]
func RevertPlan() {

}

// @Summary 主任务生成
// @Description
// @Tags MainJob
// @Param plan body model.CreateMainJobRequestInfo false "job"
// @Success 200  object model.CreateMainJobResInfo ""
// @Router /v1/ops/exec/main_job [post]
func MainJob() {

}

// @Summary 主任务执行
// @Description
// @Tags MainJob
// @Param plan query int true "main job"
// @Router /v1/ops/exec/{main_job_id} [get]
func ExecMainJob() {

}

// @Summary 主任务暂停
// @Description
// @Tags MainJob
// @Param plan query int true "main job"
// @Router /v1/ops/stop/{main_job_id} [get]
func StopMainJob() {

}

// @Summary 主任务回滚
// @Description
// @Tags MainJob
// @Param plan query int true "main job"
// @Router /v1/ops/revert/{version} [get]
func RevertMainJob() {

}

// @Summary 主任务查看
// @Description
// @Tags MainJob
// @Param job_id query int true "job"
// @Success 200  object model.QueryMainJobResInfo ""
// @Router /v1/ops/main_job/{job_id} [get]
func QueryMainJob() {

}

// @Summary 子任务创建
// @Description
// @Tags SubJob
// @Param sub_job body model.CreateSubJobReqInfo true "job"
// @Router /v1/ops/sub_job [post]
func CreateSubJob() {

}

// @Summary 子任务更新
// @Description
// @Tags SubJob
// @Param sub_job body model.CreateSubJobReqInfo true "job"
// @Router /v1/ops/sub_job/ [put]
func UpdateSubJob() {

}

// @Summary 子任务删除
// @Description
// @Tags SubJob
// @Param sub_job_id query int  true "sub_job"
// @Router /v1/ops/sub_job/{id} [delete]
func DeleteSubJob() {

}

// @Summary 子任务详情
// @Description
// @Tags SubJob
// @Param sub_job_id query model.QuerySubJobReqInfo true "sub_job"
// @Success 200  object model.QuerySubJobResInfo ""
// @Router /v1/ops/sub_job/{id} [get]
func QuerySubJob() {

}

// @Summary 子任务暂停
// @Description
// @Tags SubJob
// @Param plan query int true "sub job"
// @Router /v1/ops/stop/{sub_job_id} [get]
func StopSubMainJob() {

}

// @Summary 子任务回滚
// @Description
// @Tags SubJob
// @Param plan query int true "sub job"
// @Router /v1/ops/revert/{version} [get]
func RevertSubMainJob() {

}


