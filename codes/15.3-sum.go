package codes

/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 * 首先对数组进行排序，排序后固定一个数 nums[i]，再使用左右指针指向 nums[i]后面的两端，数字分别为 nums[L]和 nums[R]，计算三个数的和 sum 判断是否满足为 0，满足则添加进结果集
 * 如果 nums[i]大于 00，则三数之和必然无法等于 0，结束循环
 * 如果 nums[i] == nums[i-1]，则说明该数字重复，会导致结果重复，所以应该跳过
 * 当 sum == 0 时，nums[L] == nums[L+1] 则会导致结果重复，应该跳过，L++
 * 当 sum == 0 时，nums[R] == nums[R-1]] 则会导致结果重复，应该跳过，R--
 * 时间复杂度：O(n^2) n 为数组长度
 * https://leetcode-cn.com/problems/3sum/solution/hua-jie-suan-fa-15-san-shu-zhi-he-by-guanpengchn/
 */

func partition15(citations []int, p, r int) int {
	pivot := citations[r] // 使用最后一个元素为分区点
	i := p
	for j := p; j <= r-1; j++ {
		if citations[j] < pivot {
			citations[i], citations[j] = citations[j], citations[i]
			i = i + 1
		}
	}
	citations[i], citations[r] = citations[r], citations[i]
	return i
}

func quickSort15(citations []int, p, r int) []int {
	if p >= r {
		return citations
	}
	q := partition15(citations, p, r) // 获取分区点
	quickSort15(citations, p, q-1)
	quickSort15(citations, q+1, r)
	return citations
}

func threeSum(nums []int) [][]int {
	nums = quickSort15(nums, 0, len(nums)-1)
	// fmt.Println(nums)
	results := make([][]int, 0)
	last := len(nums) - 1
	for i, num := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if num > 0 {
			break
		}
		target := 0 - num
		l, r := i+1, last
		for l < r {
			if nums[l]+nums[r] > target || (r+1 <= last && nums[r] == nums[r+1]) {
				// fmt.Println("r--", nums[r], nums[r+1])
				r--
			} else if nums[l]+nums[r] < target || (l-1 > i && nums[l-1] == nums[l]) {
				// fmt.Println("l++", nums[l-1], nums[l])
				l++
			} else {
				result := []int{num, nums[l], nums[r]}
				results = append(results, result)
				l++
				r--
			}
		}
	}
	// fmt.Println(results)
	return results
}
