package main

func sortColors(nums []int) {
	left, right, i := 0, len(nums)-1, 0

	for i <= right {
		if nums[i] == 0 {
			nums[left], nums[i] = nums[i], nums[left]
			left++
		} else if nums[i] == 2 {
			nums[i], nums[right] = nums[right], nums[i]
			right--
			i--
		}
	}
}
