package codes

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */
func twoSum(nums []int, target int) []int {
	// 使用map
	valueIndex := make(map[int]int)
	result := []int{0, 0}
	for i, num := range nums {
		diff := target - num
		// fmt.Println(valueIndex, diff)
		if index, ok := valueIndex[diff]; ok {
			// fmt.Println("index: ", index, "ok: ", ok)
			result[0], result[1] = index, i
			return result
		}
		valueIndex[num] = i
	}
	return nil
}
