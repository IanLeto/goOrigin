package processor_test

import (
	"context"
	"fmt"
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

type Span struct {
	TraceID    string                 `json:"traceId"`
	SpanID     string                 `json:"spanId"`
	Time       int64                  `json:"time"`
	RemoteHost string                 `json:"remote.host"`
	BizExpand  map[string]interface{} `json:"biz.expand"`
}

func calculateTraceStatus(ctx context.Context, spans []Span, keyword string) (int, int, int) {
	traceMap := make(map[string]bool)
	successTraces := make([]string, 0)
	successCount := 0
	failureCount := 0

	for _, span := range spans {
		transTypeCode, ok := span.BizExpand["trans_type_code"].(string)
		if !ok {
			// 如果 trans_type_code 不存在或类型不是字符串,则忽略该 span
			continue
		}

		if _, exists := traceMap[span.TraceID]; !exists {
			traceMap[span.TraceID] = true
		}

		if transTypeCode != keyword {
			traceMap[span.TraceID] = false
		}
	}

	for traceID, status := range traceMap {
		if status {
			successCount++
			successTraces = append(successTraces, traceID)
		} else {
			failureCount++
		}
	}

	totalCount := len(traceMap)

	// 打印成功的 trace
	fmt.Printf("Successful Traces: %v\n", successTraces)

	return successCount, failureCount, totalCount
}

// 单元测试
func TestCalculateTraceStatus(t *testing.T) {
	testCases := []struct {
		name          string
		spans         []Span
		keyword       string
		ctx           map[string]interface{}
		expectedSucc  int
		expectedFail  int
		expectedTotal int
	}{
		{
			name: "Single trace with matching keyword",
			spans: []Span{
				{TraceID: "trace1", BizExpand: map[string]interface{}{"trans_type_code": "A"}},
				{TraceID: "trace1", BizExpand: map[string]interface{}{"trans_type_code": "A"}},
			},
			keyword:       "A",
			expectedSucc:  1,
			expectedFail:  0,
			expectedTotal: 1,
		},
		{
			name: "Single trace with non-matching keyword",
			spans: []Span{
				{TraceID: "trace1", BizExpand: map[string]interface{}{"trans_type_code": "B"}},
				{TraceID: "trace1", BizExpand: map[string]interface{}{"trans_type_code": "B"}},
			},
			keyword:       "A",
			expectedSucc:  0,
			expectedFail:  1,
			expectedTotal: 1,
		},
		{
			name: "Multiple traces with mixed keywords",
			spans: []Span{
				{TraceID: "trace1", BizExpand: map[string]interface{}{"trans_type_code": "A"}},
				{TraceID: "trace1", BizExpand: map[string]interface{}{"trans_type_code": "A"}},
				{TraceID: "trace2", BizExpand: map[string]interface{}{"trans_type_code": "B"}},
				{TraceID: "trace2", BizExpand: map[string]interface{}{"trans_type_code": "B"}},
				{TraceID: "trace3", BizExpand: map[string]interface{}{"trans_type_code": "A"}},
			},
			keyword:       "A",
			expectedSucc:  2,
			expectedFail:  1,
			expectedTotal: 3,
		},
		{
			name: "Spans with missing trans_type_code",
			spans: []Span{
				{TraceID: "trace1", BizExpand: map[string]interface{}{"trans_type_code": "A"}},
				{TraceID: "trace1", BizExpand: map[string]interface{}{}},
				{TraceID: "trace2", BizExpand: map[string]interface{}{"trans_type_code": "B"}},
			},
			keyword:       "A",
			expectedSucc:  1,
			expectedFail:  1,
			expectedTotal: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			successCount, failureCount, totalCount := calculateTraceStatus(context.TODO(), tc.spans, tc.keyword)

			if successCount != tc.expectedSucc {
				t.Errorf("Expected success count %d, but got %d", tc.expectedSucc, successCount)
			}

			if failureCount != tc.expectedFail {
				t.Errorf("Expected failure count %d, but got %d", tc.expectedFail, failureCount)
			}

			if totalCount != tc.expectedTotal {
				t.Errorf("Expected total count %d, but got %d", tc.expectedTotal, totalCount)
			}
		})
	}
}

// TestHttpClient :
func TestNodeTestConfiguration(t *testing.T) {
	suite.Run(t, new(NodeTest))
}
