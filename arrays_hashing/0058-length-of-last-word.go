package main

func lengthOfLastWord(s string) int {
	result := 0

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if result >= 1 {
				return result
			}
		} else {
			result++
		}
	}

	return result
}
