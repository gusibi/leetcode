/*
 * @lc app=leetcode id=275 lang=golang
 *
 * [275] H-Index II
 */
func hIndex1(citations []int) int {
	for i, n:= 0, len(citations); i<n; i++{
		if citations[i] >= n - i{
			return n-i
		}
	}
	return 0 
}

func hIndex(citations []int) int {
	n := len(citations)

	low := 0
	high := n - 1

	for low <= high {
		mid := low + (high-low)/2
		if citations[mid] == n-mid {
			return n - mid
		} else if citations[mid] > n-mid {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return n - low
}

