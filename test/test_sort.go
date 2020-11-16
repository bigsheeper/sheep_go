package main

import (
	"fmt"
	"sort"
	"time"

	"math/rand"
)

func goSort(nums []int64) {
	start := time.Now()

	sort.Slice(nums, func(i int, j int) bool {
		return nums[i] < nums[j]
	})

	duration := time.Since(start)
	fmt.Println(duration)
	fmt.Println(nums)
}

func mySort(nums []int64) {
	start := time.Now()

	qSort(nums, 0, len(nums)-1)

	duration := time.Since(start)
	fmt.Println(duration)
	fmt.Println(nums)
}

func qSort(nums []int64, left int, right int) {
	if left >= right {
		return
	}
	index := partition(nums, left, right)
	qSort(nums, left, index-1)
	qSort(nums, index, right)
}

func partition(nums []int64, left int, right int) int {
	pivot := nums[left]
	for left <= right {
		for nums[left] < pivot {
			left++
		}
		for nums[right] > pivot {
			right--
		}
		if left <= right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}
	return left
}

func testSort() {
	const testNum = 100 * 10000
	nums := make([]int64, testNum)
	for i := 0; i < testNum; i++ {
		nums[i] = rand.Int63() % testNum
	}
	goSort(nums)
	mySort(nums)
}
