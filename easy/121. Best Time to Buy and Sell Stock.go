package goLeetCode

func MaxProfit(prices []int) int {
	minPrice, maxProfit := prices[0], []int{0, 0}
	for day := 2; day <= len(prices); day++ {
		if prices[day-1]-minPrice > maxProfit[day-1] {
			maxProfit = append(maxProfit, prices[day-1]-minPrice)
		} else {
			maxProfit = append(maxProfit, maxProfit[day-1])
		}
		if prices[day-1] < minPrice {
			minPrice = prices[day-1]
		}
	}
	return maxProfit[len(prices)]
}
