package trans_type

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"goOrigin/API/V1"
	"goOrigin/internal/logic"
)

func CreateTransInfo(c *gin.Context) {
	var (
		req = &V1.CreateTransInfoReq{}
		res = &V1.CreateTransInfoResponse{}
		err error
	)
	if req.Region == "" {
		c.Set("region", "win")
	} else {
		c.Set("region", req.Region)
	}

	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("%s", err)
		goto ERR
	}

	err = logic.CreateType(c, req.Region, req.Items)
	if err != nil {
		logrus.Errorf("%s", err)
		goto ERR
	}

	V1.BuildResponse(c, V1.BuildInfo(res))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}
