package processor_test

import (
	"testing"
)

// 使用滑动窗口法查找最大和
func maxSumSlidingWindow(nums []int, k int) int {
	if len(nums) == 0 || k <= 0 {
		return 0
	}

	maxSum := 0
	windowSum := 0

	for i := 0; i < len(nums); i++ {
		windowSum += nums[i]

		if i >= k-1 {
			maxSum = max(maxSum, windowSum)
			windowSum -= nums[i-k+1]
		}
	}

	return maxSum
}

// 使用常规方式查找最大和
func maxSumRegular(nums []int, k int) int {
	if len(nums) == 0 || k <= 0 {
		return 0
	}

	maxSum := 0

	for i := 0; i <= len(nums)-k; i++ {
		windowSum := 0
		for j := i; j < i+k; j++ {
			windowSum += nums[j]
		}
		maxSum = max(maxSum, windowSum)
	}

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 基准测试函数
func BenchmarkMaxSumSlidingWindow(b *testing.B) {
	nums := []int{1, 4, 2, 10, 23, 3, 1, 0, 20}
	k := 4

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxSumSlidingWindow(nums, k)
	}
}

func BenchmarkMaxSumRegular(b *testing.B) {
	nums := []int{1, 4, 2, 10, 23, 3, 1, 0, 20}
	k := 4

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxSumRegular(nums, k)
	}
}
