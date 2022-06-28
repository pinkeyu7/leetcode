package main

func maxProfit1(prices []int) int {
	lowest := prices[0]
	highest := prices[0]
	profit := 0

	for i := 1; i < len(prices); i++ {
		if lowest > prices[i] {
			lowest = prices[i]
			highest = prices[i]
			continue
		}
		if highest < prices[i] {
			highest = prices[i]
			tempProfit := highest - lowest
			if tempProfit > profit {
				profit = tempProfit
			}
		}
	}

	return profit
}
