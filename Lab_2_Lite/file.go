package main

import "sort"

//Example of tested file
func findKthLargest(nums []int, k int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
    	return nums[k-1]
}
