package main

func maxProfit(prices []int) int {
	max := 0
	left, right := 0, 1

	for right < len(prices) {

		profit := prices[right] - prices[left]

		if profit > 0 {
			if max < profit {
				max = profit
			}
		} else {
			left = right
		}
		right++
	}
	return max
}
