/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */

var result = map[int]int{
	1: 1,
	2: 2,
}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	left, ok := result[n-1]
	if ok == false {
		left = climbStairs(n - 1)
		result[n-1] = left
	}
	right, ok := result[n-2]
	if ok == false {
		right = climbStairs(n - 2)
		result[n-2] = right
	}
	return left + right
}

