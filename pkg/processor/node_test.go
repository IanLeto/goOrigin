package processor_test

import (
	"context"
	"github.com/stretchr/testify/suite"
	"goOrigin/pkg/processor"
	"testing"
)

// RedisSuite :
type NodeTest struct {
	suite.Suite
	FilterNode processor.FilterTraceNode
	ctx        context.Context
}

func (s *NodeTest) SetupTest() {
	s.FilterNode = processor.FilterTraceNode{}
	s.ctx = context.TODO()

}

// TestMarshal :
func (s *NodeTest) TestConfig() {

}

// TestMarshal :
func (s *NodeTest) TestFilter() {
	testCases := []struct {
		except []byte
		input  []byte
	}{
		{except: []byte("{\"ceb.trace.gid\":\"gid-987654321\",\"ceb.trace.lid\":\"lid-123456789\",\"ceb.trace.pid\":\"pid-543210987\",\"sysName\":\"order-system\",\"unitcode\":\"unit-02\",\"instanceZone\":\"zone-1\",\"timestamp\":\"2024-12-04T14:48:10+08:00\",\"local.app\":\"order-service\",\"traceId\":\"wsS3lVjcmgVreGLet7xtoW2ynr\",\"spanId\":\"1.0.1.0\",\"businessId\":\"order-123456\",\"span.kind\":\"server\",\"result.code\":\"201\",\"time\":1733292189,\"remote.host\":\"172.8.20.64\",\"biz.expand\":{\"ret_code\":\"AAAA\",\"trans_channel\":\"O999\",\"trans_type_code\":\"1001\"}}"), input: []byte("{\"ceb.trace.gid\":\"gid-987654321\",\"ceb.trace.lid\":\"lid-123456789\",\"ceb.trace.pid\":\"pid-543210987\",\"sysName\":\"order-system\",\"unitcode\":\"unit-02\",\"instanceZone\":\"zone-1\",\"timestamp\":\"2024-12-04T14:48:10+08:00\",\"local.app\":\"order-service\",\"traceId\":\"wsS3lVjcmgVreGLet7xtoW2ynr\",\"spanId\":\"1.0.1.0\",\"businessId\":\"order-123456\",\"span.kind\":\"server\",\"result.code\":\"201\",\"time\":1733292189,\"remote.host\":\"172.8.20.64\",\"biz.expand\":{\"ret_code\":\"AAAA\",\"trans_channel\":\"O999\",\"trans_type_code\":\"1001\"}}")}, // 常规输入
		// 输入为nil
		{
			input: []byte("" +
				"{\"ceb.trace.gid\":\"gid-987654321\",\"ceb.trace.lid\":\"lid-123456789\",\"ceb.trace.pid\":\"pid-543210987\",\"sysName\":\"order-system\",\"unitcode\":\"unit-02\",\"instanceZone\":\"zone-1\",\"timestamp\":\"2024-12-04T14:48:10+08:00\",\"local.app\":\"order-service\",\"traceId\":\"wsS3lVjcmgVreGLet7xtoW2ynr\",\"spanId\":\"1.0.1.0\",\"businessId\":\"order-123456\",\"span.kind\":\"server\",\"result.code\":\"201\",\"time\":1733292189,\"remote.host\":\"172.8.20.64\",\"biz.expand\":{\"ret_code\":\"AAAA\",\"trans_channel\":\"O999\",\"trans_type_code\":\"\"}}"),
			except: nil,
		},
		// 输入的数据不计入交易

	}
	for _, v := range testCases {
		res, _ := s.FilterNode.Process(s.ctx, v.input)
		s.Equal(v.except, res)
	}
}

// TestHttpClient :
func TestNodeTestConfiguration(t *testing.T) {
	suite.Run(t, new(NodeTest))
}
