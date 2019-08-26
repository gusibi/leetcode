/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func maxArea(height []int) int {
	i, j := 0, len(height)-1
	sum := 0
	for i < j {
		if height[i] < height[j] {
			sum = max(height[i]*(j-i), sum)
			i++
		} else {
			sum = max(height[j]*(j-i), sum)
			j--
		}
	}
	return sum
}

