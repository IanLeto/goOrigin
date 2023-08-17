package k8sHandlers

import (
	"fmt"
	"github.com/cstockton/go-conv"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"goOrigin/internal/params"
	"goOrigin/internal/service"
	"strconv"
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
	paramStr := c.Query("step")
	// 使用 strconv.Atoi 将字符串转换为整数
	req.Step, err = strconv.Atoi(paramStr)
	req.Location, _ = conv.String(c.Query("location"))
	req.Step, _ = conv.Int(c.Query("step2"))
	req.LimitByte, _ = conv.Int(c.Query("limit_byte"))
	req.LimitLine, _ = conv.Int(c.Query("limit_line"))

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
