package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/api/prometheus/v1"
	"goOrigin/API/V1"
	"goOrigin/pkg/clients"
	"time"
)

func QueryPromWeight(c *gin.Context, req *V1.QueryWeightRequest) (string, error) {
	var (
		err error
	)
	client, err := clients.NewPromClient()
	if err != nil {
		logger.Error(fmt.Sprintf("%s", err))
		return "", err
	}
	api := v1.NewAPI(client)
	query := fmt.Sprintf("%s", req.Metric)
	res, _, err := api.Query(c, query, time.Now())
	if err != nil {
		logger.Error(fmt.Sprintf("%s", err))
		return "", err
	}
	return res.String(), err
}
