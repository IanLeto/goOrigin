package processor_test

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"testing"
	"time"
)

// ========================== 方案 1：全量存储数据 ==========================
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

	if elapsed >= 6000 {
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

// 方案 1：基于事件数组的方案
func BenchmarkFullCache(b *testing.B) {
	cache := NewFullCacheSuccessRate(60 * time.Second)

	// 创建内存分析文件
	memProfile, _ := os.Create("full_cache_mem.prof")
	defer memProfile.Close()

	// 运行测试
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Update(i%2 == 0) // 模拟交替存储成功/失败
		cache.GetSuccessRate()
	}

	// 记录内存使用情况
	runtime.GC() // 先进行垃圾回收，确保拿到的是实际的内存占用
	_ = pprof.WriteHeapProfile(memProfile)
}

// 方案 2：基于滑动窗口计数法
func BenchmarkSlidingWindow(b *testing.B) {
	window := NewSlidingWindowSuccessRate()

	// 创建内存分析文件
	memProfile, _ := os.Create("sliding_window_mem.prof")
	defer memProfile.Close()

	// 运行测试
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		window.Update(i%2 == 0) // 模拟交替存储成功/失败
		window.GetSuccessRate()
	}

	// 记录内存使用情况
	runtime.GC() // 先进行垃圾回收，确保拿到的是实际的内存占用
	_ = pprof.WriteHeapProfile(memProfile)
}
func TestRun(t *testing.T) {
	cache := NewFullCacheSuccessRate(60 * time.Second)
	window := NewSlidingWindowSuccessRate()

	// 模拟 1000 条数据流
	for i := 0; i < 1000; i++ {
		cache.Update(i%2 == 0)
		window.Update(i%2 == 0)
	}

	fmt.Printf("全量缓存成功率: %.2f%%\n", cache.GetSuccessRate()*100)
	fmt.Printf("滑动窗口成功率: %.2f%%\n", window.GetSuccessRate()*100)
}
