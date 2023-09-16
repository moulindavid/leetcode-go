package main

func twoSum(nums []int, target int) []int {
	complement_map := make(map[int]int)

	for idx, num := range nums {

		if val, found := complement_map[target-num]; found {
			return []int{val, idx}
		}

		complement_map[num] = idx
	}
	return nil
}
