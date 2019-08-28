package codes

/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 */

func partition56(intervals [][]int, p, r int) int {
	// 分区，把大于分区的放分区点右边，小于分区的放左边
	pivot := intervals[r]
	i := p
	for j := p; j <= r-1; j++ {
		if intervals[j][0] < pivot[0] {
			intervals[i], intervals[j] = intervals[j], intervals[i]
			i = i + 1
		}
	}
	intervals[i], intervals[r] = intervals[r], intervals[i]
	return i
}

func quickSort56(intervals [][]int, p, r int) [][]int {
	if p >= r {
		return intervals
	}
	q := partition56(intervals, p, r) // 获取分区点
	quickSort56(intervals, p, q-1)
	quickSort56(intervals, q+1, r)
	return intervals
}

func merge56(intervals [][]int) [][]int {
	n := len(intervals)
	if n == 0 {
		return intervals
	}
	intervals = quickSort56(intervals, 0, n-1)
	results := [][]int{}
	left, right := intervals[0][0], intervals[0][1]
	for i := 0; i < n; i++ {
		if right <= intervals[i][1] {
			right = intervals[i][1]
		}
		if i == n-1 {
			results = append(results, []int{left, right})
		} else {
			if intervals[i+1][0] > right {
				results = append(results, []int{left, right})
				left = intervals[i+1][0]
			}
		}
	}
	return results
}
