package main

func topKFrequent(nums []int, k int) []int {
	count_map := map[int]int{}

	for _, num := range nums {
		count_map[num]++
	}

	count_slice := make([][]int, len(nums)+1)

	for num, count := range count_map {
		count_slice[count] = append(count_slice[count], num)
	}

	var result []int

	for i := len(count_slice) - 1; i > 0; i-- {
		result = append(result, count_slice[i]...)
		if len(result) == k {
			break
		}
	}

	return result
}
