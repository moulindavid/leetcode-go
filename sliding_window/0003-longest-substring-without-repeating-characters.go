package main

func lengthOfLongestSubstring(s string) int {

	longest := 0
	char_map := make(map[byte]bool)
	l := 0

	for r := range s {
		for char_map[s[r]] {
			delete(char_map, s[l])
			l++
		}
		char_map[s[r]] = true
		longest = max(longest, r-l+1)
	}

	return longest
}

func max(actual, other int) int {
	if actual < other {
		return other
	}
	return actual
}
