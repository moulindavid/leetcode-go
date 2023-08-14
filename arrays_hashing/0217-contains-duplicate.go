package main

func containsDuplicate(nums []int) bool {
	nums_map := make(map[int]bool)

	for _, num := range nums {
		if _, ok := nums_map[num]; ok {
			return true
		} else {
			nums_map[num] = true
		}
	}

	return false
}
