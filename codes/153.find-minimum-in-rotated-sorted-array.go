/*
 * @lc app=leetcode id=153 lang=golang
 *
 * [153] Find Minimum in Rotated Sorted Array
 */
func findMin(nums []int) int{
	// 二分查找
	n := len(nums)
	left, right := 0, n -1
	mid := n/2
	for {
		if mid == 0{
			return nums[mid]
		}else if nums[mid] < nums[mid-1]{
			return nums[mid]
		}else if nums[n-1] < nums[n-2] {
			return nums[n-1]
		}else if nums[mid] > nums[n-1]{
			left = mid
			mid = (left + right) / 2
		}else if nums[mid] <= nums[n-1]{
			right = mid
			mid = (left + right) / 2
		}
	}
	return nums[mid]
}
func findMin1(nums []int) int {
	// 遍历
    for i:= 0; i < len(nums); i++{
		if nums[i] < nums[0]{
			return nums[i]
		}
	}
	return nums[0]
}

