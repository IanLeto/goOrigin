package conf

var Conf = NewConfig()


//func (c *CCApiClient) GetServiceInstance(bizID, limit, start int, ServiceInstanceIds []int) (*CCSearchServiceInstanceResponseData, error) {
//	defer c.SearchHostTimeObserver.Start().Finish()
//	result := struct {
//		APIResponse
//		Data *CCSearchServiceInstanceResponseData `json:"data"`
//	}{}
//	logger := logging.GetStdLogger()
//	response, err := c.Agent().Post("list_service_instance_detail/").BodyProvider(&json.Provider{
//		Payload: &CCSearchServiceInstanceRequest{
//			CommonArgs: c.client.CommonArgs(),
//			BkBizID:    bizID,
//			Limit: CCSearchServiceInstanceRequestMetadataLabelPage{
//				Start: start,
//				Limit: limit,
//				Sort:  "bk_host_id",
//			},
//		}}).Receive(&result, &result)
//
//	if err != nil {
//		c.SearchHostCounter.CounterFails.Inc()
//		logger.Warnf("get service_instance failed: %v, %v", result, err)
//		return nil, err
//	}
//
//	logger.Debugf("get service_instance response: %d, %v", response.StatusCode, result.Message)
//	if result.Data == nil {
//		c.SearchHostCounter.CounterFails.Inc()
//		logger.Warnf("%s query from cc error %d: %v", result.RequestID, result.Code, result.Message)
//		return nil, errors.Wrapf(define.ErrOperationForbidden, result.Message)
//	}
//
//	c.SearchHostCounter.CounterSuccesses.Inc()
//	return result.Data, nil
//}