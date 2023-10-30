package main

import "math"

func getLast(a []int) int {
	return a[len(a)-1]
}

func contains(a []int, e int) bool {
	for _, v := range a {
		if v == e {
			return true
		}
	}
	return false
}

func findKthLargest(nums []int, k int) int {
	var maxIds []int
	for i := 0; i < k; i++ {
		lastMax := math.MinInt
		var lastMaxId int
		for i1, v1 := range nums {
			if contains(maxIds, i1) {
				continue
			}
			if v1 > lastMax {
				lastMax = v1
				lastMaxId = i1
			}
		}
		maxIds = append(maxIds, lastMaxId)
	}

	return nums[getLast(maxIds)]
}
