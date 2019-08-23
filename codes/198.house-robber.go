/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 */
// I: O(n), S = O(3n), T = O(3n)
func rob(nums []int) int {
	currSum, preSum := 0, 0
	// 计算抢当前这个收益大还是下一个收益大
	for _, num := range nums {
		temp := currSum
		if preSum+num > currSum {
			currSum = preSum + num
		}
		preSum = temp
	}
	return currSum
}

