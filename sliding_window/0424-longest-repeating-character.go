package main

func characterReplacement(s string, k int) int {
	count := make(map[byte]int)
	left, result, max_letter_count := 0, 0, 0

	for right := range s {
		count[s[right]] = 1 + count[s[right]]
		max_letter_count = max(max_letter_count, count[s[right]])

		if (right-left+1)-max_letter_count > k {
			count[s[left]] = count[s[left]] - 1
			left++
		}
		result = max(result, right-left+1)
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
