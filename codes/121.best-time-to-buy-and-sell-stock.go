package codes

/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 * 记录【今天之前买入的最小值】
 * 计算【今天之前最小值买入，今天卖出的获利】，也即【今天卖出的最大获利】
 * 比较【每天的最大获利】，取最大值即可
 */

func maxItem(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	m := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] > m {
			m = nums[i]
		}
	}
	return m
}
func maxProfit1(prices []int) int {
	profit := 0
	for i := 0; i < len(prices)-1; i++ {
		maxPrice := maxItem(prices[i+1 : len(prices)])
		if maxPrice-prices[i] > profit {
			profit = maxPrice - prices[i]
		}
	}
	return profit
}
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minPrice := prices[0]
	profit := 0
	for i := 0; i < len(prices); i++ {
		currProfit := prices[i] - minPrice
		if currProfit > profit {
			profit = currProfit
		}
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}
	return profit
}
