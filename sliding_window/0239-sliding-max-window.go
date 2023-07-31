package main

func maxSlidingWindow(nums []int, k int) []int {
	left, right := 0, 0

	result := []int{}
	queue := make([]int, 0) //Indexes

	for right < len(nums) {
		for len(queue) != 0 && nums[queue[len(queue)-1]] < nums[right] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, right)

		if left > queue[0] {
			queue = queue[1:]
		}
		if (right + 1) >= k {
			result = append(result, nums[queue[0]])
			left++
		}

		right++
	}
	return result
}
