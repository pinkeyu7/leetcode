package main

func maxProfit2(prices []int) int {
	current := prices[0]
	profit := 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < current {
			current = prices[i]
		} else {
			profit += prices[i] - current
			current = prices[i]
		}
	}

	return profit
}
