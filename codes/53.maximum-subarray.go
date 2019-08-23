/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */
func maxSubArray(nums []int) int {
	sum, min, max := 0, 0, nums[0]
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
		if sum-min >= max {
			max = sum - min
		}
		if sum < min {
			min = sum
		}
	}
	return max
}

