package codes

/*
 * @lc app=leetcode id=122 lang=golang
 *
 * [122] Best Time to Buy and Sell Stock II
 */

import "math"

func maxProfit2(prices []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	// profit0 表示 当前没有持有的收益
	// profit1 表示 当前持有的收益 负无穷是初始值
	profit0, profit1 := 0, int(math.Inf(-1))
	for i := 0; i < len(prices); i++ {
		temp := profit0
		// profit0 表示未持有，可能的情况是前一天持有卖出或者前一天未持有
		profit0 = max(profit0, profit1+prices[i])
		// profit1 表示持有，可能的情况是前一天持有，或者前一天未持有今天买入
		profit1 = max(profit1, temp-prices[i])
	}
	// 最后的状态需要是未持有收益最大所以
	return profit0
}
