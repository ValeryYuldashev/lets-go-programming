package main

func findKthLargest(nums []int, k int) int {
	lenght := len(nums)

	for i := 1; i < lenght; i++ {
		j := i - 1
		v := nums[i]
		for j >= 0 && nums[j] < v {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = v
	}

	return nums[k]
}
