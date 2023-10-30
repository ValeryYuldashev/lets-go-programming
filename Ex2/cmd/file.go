package main

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] > pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSort(a []int, min int, max int) []int {
	if min < max {
		a, pivot := partition(a, min, max)
		a = quickSort(a, min, pivot-1)
		a = quickSort(a, pivot+1, max)
	}
	return a
}

func findKthLargest(nums []int, k int) int {
	quickSort(nums, 0, len(nums)-1)
	return nums[k-1]
}
