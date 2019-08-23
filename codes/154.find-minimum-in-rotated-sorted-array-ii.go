/*
 * @lc app=leetcode id=154 lang=golang
 *
 * [154] Find Minimum in Rotated Sorted Array II
 */
func findMin(nums []int) int {
	// 遍历
    for i:= 0; i < len(nums); i++{
		if nums[i] < nums[0]{
			return nums[i]
		}
	}
	return nums[0]
}

