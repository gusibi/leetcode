/*
 * @lc app=leetcode id=167 lang=golang
 *
 * [167] Two Sum II - Input array is sorted
 */
func twoSumBad(numbers []int, target int) []int {
	var results []int
	for i := 0; i < len(numbers); i++{
		if numbers[i]> target{
			break
		}
		for j:= i+1; j<len(numbers); j++{
			if numbers[i] + numbers[j] == target{
				results = []int{i+1, j+1}
			}else if numbers[i] + numbers[j] > target{
				break
			}
		}
	}
	return results
}

func twoSum(numbers []int, target int) []int{
	low, high := 0, len(numbers) -1
	for low < high{
		sum := numbers[low] + numbers[high]
		if sum == target{
			return []int{low+1, high+1}
		}else if sum < target{
			low++
		}else if sum > target{
			high--
		}
	}
	return []int{-1, -1}
}

