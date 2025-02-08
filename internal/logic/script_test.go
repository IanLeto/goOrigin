package logic

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"goOrigin/API/V1"
	"io/ioutil"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"goOrigin/config"
)

type NodeEntity struct {
	ID    int
	Name  string
	Value int
}

type ScriptAPISuite struct {
	suite.Suite
	conf *config.Config
}

// Directly sets the value of the field.
func updateNodeDirectly(a1, a2 *NodeEntity) {
	if a1.Value != a2.Value {
		a1.Value = a2.Value
	}
}

// Uses reflection to set the value of the field.
func updateNodeWithReflection(a1, a2 *NodeEntity) {
	valA1 := reflect.ValueOf(a1).Elem()
	valA2 := reflect.ValueOf(a2).Elem()

	for i := 0; i < valA1.NumField(); i++ {
		if valA1.Field(i).Interface() != valA2.Field(i).Interface() {
			valA1.Field(i).Set(valA2.Field(i))
		}
	}
}

func (s *ScriptAPISuite) SetupTest() {
	// Setup logic if needed
}
func CurrentLogs(cluster string, info *V1.GetLogsReqInfo) (*V1.GetLogsRes, error) {
	var (
		byteLimit = int64(info.LimitByte)
		byteLine  = int64(info.LimitLine)
		res       = &V1.GetLogsRes{}
		count     = 0
	)

	logOptions := &v1.PodLogOptions{
		Container:                    info.Container,
		Timestamps:                   true,       // 是否附带时间戳
		TailLines:                    &byteLine,  // 最大行数限制
		LimitBytes:                   &byteLimit, // 最大字节数限制
		InsecureSkipTLSVerifyBackend: false,
	}

	config, err := clientcmd.BuildConfigFromFlags("", "/home/ian/.kube/config")
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		logrus.Errorf("get client error: %s", err)
		return nil, err

	}
	reader, err := client.CoreV1().RESTClient().Get().Namespace(info.Ns).Name(info.PodID).Resource("pods").
		SubResource("log").VersionedParams(logOptions, scheme.ParameterCodec).Stream(context.Background())
	if err != nil {
		logrus.Errorf("get logs error: %s", err)
		return nil, err
	}
	content, err := ioutil.ReadAll(reader)
	contents := strings.Split(string(content), "\n")
	if info.Size == 0 {
		info.Size = 100
	}
	// 是否需要翻页
	var isForward bool = info.FromDate != "" && info.ToDate != ""

	switch {
	case len(contents) == 0: // 无数据 返回空
		break
	case isForward: // 按时间段查询，contents 返回5000行
		break
	case len(contents) <= info.Size: // 日志少于期望查询数量，无论怎样都会返回所有日志
		contents = contents[0:info.Size]
	case !isForward && info.Location == "begin": // 从头开始查询
	case !isForward && info.Location == "end": // 从尾部
		contents = contents[len(contents)-1-info.Size : len(contents)-1]
	case info.Location == "" && info.FromDate == "" && info.ToDate == "":
		contents = contents[len(contents)-1-info.Size : len(contents)-1]
	default:
		contents = contents[len(contents)-1-info.Size : len(contents)-1]
	}

	lines := contents
	entries := make([]Entry, 0)
	var fromTimestamp, toTimestamp int64
	// 如果有时间段，需要解析时间段
	if info.FromDate != "" {
		fromTime, err := time.Parse(time.RFC3339Nano, info.FromDate)
		fromTimestamp = fromTime.UnixNano()
		if err != nil {
			fmt.Printf("Error parsing from date: %v\n", err)
			return nil, err
		}
		toTimest, err := time.Parse(time.RFC3339Nano, info.ToDate)
		toTimestamp = toTimest.UnixNano()
		if err != nil {
			fmt.Printf("Error parsing from date: %v\n", err)
			return nil, err
		}
		res.FromDate = fromTime.Format(time.RFC3339Nano)
		res.FromDate = toTimest.Format(time.RFC3339Nano)
	}

	// 定义一个函数类型，用于处理不同的条件
	type entryHandler func(timestamp int64, entry Entry) bool
	// 根据条件选择合适的处理方式
	var handleEntry entryHandler
	// 如果向后翻页，
	if isForward && info.Step >= 0 {
		handleEntry = func(timestamp int64, entry Entry) bool {
			// 如果当前数据的时间戳大于等于前端传入的时间片段的最大值,也就是todata
			if timestamp > toTimestamp {
				entries = append(entries, entry)
				return true
			}
			return false
		}
	} else if isForward && info.Step < 0 {
		handleEntry = func(timestamp int64, entry Entry) bool {
			// 因为是向前翻页，所以需要反转数组
			// 如果当前数据的时间戳小于等于结束时间，就返回
			if timestamp < fromTimestamp {
				entries = append(entries, entry)
				return true
			}
			return false
		}
	} else {
		handleEntry = func(timestamp int64, entry Entry) bool {
			entries = append(entries, entry)
			return true
		}
	}
	// 向前翻页，需要反转数组
	if isForward && info.Step < 0 {
		reverseArray(lines)
	}

	for _, line := range lines {
		parts := strings.SplitN(line, " ", 2)
		// 正常的数据格式为：时间戳 内容
		if len(parts) >= 2 {
			date := parts[0]
			content := parts[1]
			timestamp, err := time.Parse(time.RFC3339Nano, date)
			if err != nil {
				fmt.Printf("Error parsing date: %v\n", err)
				continue
			}
			entry := Entry{
				Timestamp: timestamp.UnixNano(),
				Date:      date,
				Content:   content,
			}

			if handleEntry(timestamp.UnixNano(), entry) {
				count++
				if count >= info.Size {
					break
				}
			}
		}
	}

	var fn = func(arr []Entry) {
		for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	if info.Step < 0 {
		fn(entries)
	}
	for _, entry := range entries {
		var epl Entry = entry
		res.Items = append(res.Items, epl)
	}
	res.FromDate = string(entries[0].Date) //
	res.ToDate = string(entries[len(entries)-1].Date)

	return res, err

}
func (s *ScriptAPISuite) TestConfig() {
	res, err := CurrentLogs("cluster", &V1.GetLogsReqInfo{
		Container: "tool-filebeat",
		PodID:     "tool-bc48db55f-bdd5x",
		Ns:        "default",
		Size:      10,
		LimitByte: 1000000,
		LimitLine: 10000,
	})
	s.NoError(err)
	s.NotNil(res)
	for i, item := range res.Items {
		fmt.Println(i, item)
	}
	fmt.Println(res.FromDate, res.ToDate)
}
func (s *ScriptAPISuite) TestForward() {
	res, err := CurrentLogs("cluster", &V1.GetLogsReqInfo{
		Container: "tool-filebeat",
		PodID:     "tool-6656fdb57c-grgp8",
		Ns:        "default",
		Size:      10,
		LimitByte: 10000,
		LimitLine: 10000,
		//Step:      -1,
		Location: "end",
		//FromDate: "2024-04-13T15:34:48.964686598Z",
		//ToDate:   "2024-04-13T15:34:52.964823124Z",
	})
	s.NoError(err)
	s.NotNil(res)
	for i, item := range res.Items {
		fmt.Println(i, item)
	}
	fmt.Println(res.FromDate, res.ToDate)
}

// Add performance tests here.
func (s *ScriptAPISuite) TestUpdateNodeDirectly() {
	a1 := &NodeEntity{Value: 1}
	a2 := &NodeEntity{Value: 2}

	s.Run("UpdateNodeDirectly", func() {
		updateNodeDirectly(a1, a2)
		s.Equal(2, a1.Value)
	})
}

func (s *ScriptAPISuite) TestUpdateNodeWithReflection() {
	a1 := &NodeEntity{Value: 1}
	a2 := &NodeEntity{Value: 2}

	s.Run("UpdateNodeWithReflection", func() {
		updateNodeWithReflection(a1, a2)
		s.Equal(2, a1.Value)
	})
}

// Benchmark for direct field access.
func BenchmarkUpdateNodeDirectly(b *testing.B) {
	a1 := &NodeEntity{Value: 1}
	a2 := &NodeEntity{Value: 2}

	for i := 0; i < b.N; i++ {
		updateNodeDirectly(a1, a2)
	}
}

// Benchmark for field access using reflection.
func BenchmarkUpdateNodeWithReflection(b *testing.B) {
	a1 := &NodeEntity{Value: 1}
	a2 := &NodeEntity{Value: 2}

	for i := 0; i < b.N; i++ {
		updateNodeWithReflection(a1, a2)
	}
}

type FullCacheSuccessRate struct {
	mu     sync.Mutex
	events []bool
	window time.Duration
}

func NewFullCacheSuccessRate(window time.Duration) *FullCacheSuccessRate {
	return &FullCacheSuccessRate{
		events: make([]bool, 0),
		window: window,
	}
}

func (f *FullCacheSuccessRate) Update(isSuccess bool) {
	f.mu.Lock()
	defer f.mu.Unlock()

	now := time.Now()
	f.events = append(f.events, isSuccess)

	// 清理超过窗口的数据
	cutoff := now.Add(-f.window)
	newEvents := make([]bool, 0, len(f.events))
	for _, event := range f.events {
		if now.Sub(cutoff) <= f.window {
			newEvents = append(newEvents, event)
		}
	}
	f.events = newEvents
}

func (f *FullCacheSuccessRate) GetSuccessRate() float64 {
	f.mu.Lock()
	defer f.mu.Unlock()

	if len(f.events) == 0 {
		return 0.0
	}

	successCount := 0
	for _, event := range f.events {
		if event {
			successCount++
		}
	}
	return float64(successCount) / float64(len(f.events))
}

// ========================== 方案 2：滑动窗口计数法 ==========================
type SlidingWindowSuccessRate struct {
	mu             sync.Mutex
	buckets        [60][2]int // [成功数, 总请求数]
	lastUpdateTime int64
}

func NewSlidingWindowSuccessRate() *SlidingWindowSuccessRate {
	return &SlidingWindowSuccessRate{
		lastUpdateTime: time.Now().Unix(),
	}
}

func (s *SlidingWindowSuccessRate) Update(isSuccess bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now().Unix()
	elapsed := now - s.lastUpdateTime

	if elapsed >= 60 {
		// 超过窗口范围，直接清空所有桶
		for i := range s.buckets {
			s.buckets[i] = [2]int{0, 0}
		}
	} else {
		// 仅清除过期的数据
		for i := int64(0); i < elapsed; i++ {
			s.buckets[(s.lastUpdateTime+i)%60] = [2]int{0, 0}
		}
	}

	s.lastUpdateTime = now
	index := now % 60
	s.buckets[index][1]++ // 总请求数 +1
	if isSuccess {
		s.buckets[index][0]++ // 成功请求数 +1
	}
}

func (s *SlidingWindowSuccessRate) GetSuccessRate() float64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	totalSuccess := 0
	totalCount := 0

	for _, bucket := range s.buckets {
		totalSuccess += bucket[0]
		totalCount += bucket[1]
	}

	if totalCount == 0 {
		return 0.0
	}
	return float64(totalSuccess) / float64(totalCount)
}

// ========================== Benchmark 测试 ==========================
func BenchmarkFullCache(b *testing.B) {
	cache := NewFullCacheSuccessRate(60 * time.Second)
	for i := 0; i < b.N; i++ {
		cache.Update(i%2 == 0) // 交替存储成功/失败
		cache.GetSuccessRate()
	}
}

func BenchmarkSlidingWindow(b *testing.B) {
	window := NewSlidingWindowSuccessRate()
	for i := 0; i < b.N; i++ {
		window.Update(i%2 == 0) // 交替存储成功/失败
		window.GetSuccessRate()
	}
}

// 打印odametric 计算公式
func (s *ScriptAPISuite) TestOdaMetric() {
	OdaSuccessAndFailedRateMetric(nil, "region", &V1.SuccessRateReqInfo{})
}
func TestScriptConfiguration(t *testing.T) {
	suite.Run(t, new(ScriptAPISuite))
}
