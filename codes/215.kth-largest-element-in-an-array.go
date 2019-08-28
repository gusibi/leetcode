package codes

/*
 * @lc app=leetcode id=215 lang=golang
 *
 * [215] Kth Largest Element in an Array
 *
 * https://leetcode.com/problems/kth-largest-element-in-an-array/description/
 *
 * algorithms
 * Medium (46.29%)
 * Total Accepted:    341.2K
 * Total Submissions: 732.6K
 * Testcase Example:  '[3,2,1,5,6,4]\n2'
 *
 * Find the kth largest element in an unsorted array. Note that it is the kth
 * largest element in the sorted order, not the kth distinct element.
 *
 * Example 1:
 *
 *
 * Input: [3,2,1,5,6,4] and k = 2
 * Output: 5
 *
 *
 * Example 2:
 *
 *
 * Input: [3,2,3,1,2,4,5,5,6] and k = 4
 * Output: 4
 *
 * Note:
 * You may assume k is always valid, 1 ≤ k ≤ array's length.
 *
 */
func findKthLargest(nums []int, k int) int {
	// 1. 快排，按索引取值
	// 2. 最大堆对数组建立堆，只是需要找到第k大的数所以没有必要像直接排序方法中那样对所有数进行排序。建立好最大队后每次都提取出最大的数，提取 k 次就得到了第 k 大的数。建立最大堆的时间复杂度为O(n)，提取最大数后每次调整最大队时间复杂为O(logn)，所以总的时间复杂度为O(n+k*logn)。
	return 0
}
