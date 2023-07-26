package main

func minWindow(s string, t string) string {
	left, right := 0, 0
	target_frequency := make(map[uint8]int)
	current_frequency := make(map[uint8]int)
	have := 0
	result := ""

	for index := range t {
		target_frequency[t[index]]++
	}

	for right < len(s) {
		current_frequency[s[right]]++

		if target_frequency[s[right]] != 0 && target_frequency[s[right]] == current_frequency[s[right]] {
			have++
		}

		for have == len(target_frequency) {
			if result == "" {
				result = s[left : right+1]
			}

			if right-left+1 < len(result) {
				result = s[left : right+1]
			}

			current_frequency[s[left]]--
			//If we are no longer at the wanted number of given char
			if current_frequency[s[left]] < target_frequency[s[left]] {
				have--
			}
			left++
		}
		right++
	}
	return result
}
