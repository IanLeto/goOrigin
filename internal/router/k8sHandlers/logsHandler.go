package k8sHandlers

import (
	"fmt"
	"github.com/cstockton/go-conv"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"goOrigin/internal/params"
	"goOrigin/internal/service"
)

func GetCurrentLogs(c *gin.Context) {
	var (
		req = &params.GetLogsReq{}
		res = &params.GetLogsRes{}
		err error
	)
	req.Container = c.Query("container")
	req.Ns = c.Query("ns")
	req.PodID = c.Query("pod_id")
	req.Cluster = c.Query("cluster")
	req.FromDate = c.Query("from_date")
	req.ToDate = c.Query("to_date")
	req.Size, _ = conv.Int(c.Query("size"))
	req.Step, _ = conv.Int(c.Query("step"))

	res, err = service.GetCurrentLogs(c, req)
	if err != nil {
		logrus.Errorf("%s", err)
		goto ERR
	}
	params.BuildResponse(c, params.BuildInfo(res))
	return

ERR:
	params.BuildErrResponse(c, params.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))

}
