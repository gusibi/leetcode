/*
 * @lc app=leetcode id=120 lang=golang
 *
 * [120] Triangle
 */

func min(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	minNum := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] < minNum {
			minNum = nums[i]
		}
	}
	return minNum
}

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	// 从倒数第二行开始，计算当前位置 和 下一行的对应位置和下一行的右一位置的值，取值小的写入
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += min([]int{triangle[i+1][j], triangle[i+1][j+1]})
		}
	}
	return triangle[0][0]
}

