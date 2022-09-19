package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/api/prometheus/v1"
	"goOrigin/internal/params"
	"goOrigin/pkg/clients"
	"goOrigin/pkg/logger"
	"time"
)

func QueryPromWeight(c *gin.Context, req *params.QueryWeightRequest) (string, error) {
	var (
		err error
		log = logger.Logger
	)
	client, err := clients.NewPromClient()
	if err != nil {
		log.Error(fmt.Sprintf("%s", err))
		return "", err
	}
	api := v1.NewAPI(client)
	query := fmt.Sprintf("%s", req.Metric)
	res, _, err := api.Query(c, query, time.Now())
	if err != nil {
		log.Error(fmt.Sprintf("%s", err))
		return "", err
	}
	return res.String(), err
}
