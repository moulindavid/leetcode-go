package linkedlist

// value are range [1, n]
// len = n + 1
// we need can see each value as pointer
// each value points to a position
// none of them is pointing to 0.

func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]

	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
