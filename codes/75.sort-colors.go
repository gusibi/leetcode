/*
 * @lc app=leetcode id=75 lang=golang
 *
 * [75] Sort Colors
 */
func sortColors1(nums []int) []int {
	results := map[int]int{0: 0, 1: 0, 2: 0}
	for _, color := range nums {
		results[color] = results[color] + 1
	}
	for i, _ := range nums {
		if i < results[0] {
			nums[i] = 0
		} else if i < results[0]+results[1] {
			nums[i] = 1
		} else {
			nums[i] = 2
		}
	}
	return nums
}

func sortColors(nums []int) []int {
	head, tail := 0, len(nums)-1
	i := 0
	for i <= tail {
		if nums[i] == 0 {
			nums[i], nums[head] = nums[head], nums[i]
			head++
		} else if nums[i] == 2 {
			nums[i], nums[tail] = nums[tail], nums[i]
			tail--
			i--
		}
		i++
	}
	return nums
}

