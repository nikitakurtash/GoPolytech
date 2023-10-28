package main

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	h := 1
	for h < n/3 {
		h = 3*h + 1
	}

	for h >= 1 {
		for i := h; i < n; i++ {
			for j := i; j >= h && nums[j] < nums[j-h]; j -= h {
				nums[j], nums[j-h] = nums[j-h], nums[j]
			}
		}
		h /= 3
	}
	return nums[len(nums)-k]
}
