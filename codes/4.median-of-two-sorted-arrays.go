package codes

/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays

 */
// package codes

// import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	if length == 0 {
		return 0.0
	} else if length == 1 {
		if len(nums1) == 1 {
			return float64(nums1[0])
		} else {
			return float64(nums2[0])
		}
	}

	mod := length % 2
	var mid [2]int
	if mod == 1 {
		mid = [2]int{length / 2, length / 2}
	} else {
		mid = [2]int{length/2 - 1, length / 2}
	}
	i, j, count := 0, 0, 0
	var m, n, num int
	var nums []int
	//    fmt.Println("mid: ", mid, "mod: ", mod)
	for i < len(nums1) || j < len(nums2) {
		if i < len(nums1) && j < len(nums2) {
			if nums1[i] < nums2[j] {
				num = nums1[i]
				i++
			} else {
				num = nums2[j]
				j++
			}
		} else if i >= len(nums1) && j < len(nums2) {
			num = nums2[j]
			j++
		} else if i < len(nums1) && j >= len(nums2) {
			num = nums1[i]
			i++
		}
		//    fmt.Println("num: ", num, count)
		nums = append(nums, num)
		if count == mid[0] && mod == 1 {
			return float64(num)
		} else if count == mid[0] && mod == 0 {
			m = num
		} else if count == mid[1] && mod == 0 {
			n = num
		}
		count++
	}
	//    fmt.Println("m: ", m, "n: ", n)
	return float64((m + n)) / 2
}

// func main(){
// 	// nums1 := []int{1, 3}
// 	// nums2 := []int{2}
// 	// fmt.Println(findMedianSortedArrays(nums1, nums2))
// 	nums3 := []int{1, 3}
// 	nums4 := []int{2, 4}
// 	fmt.Println(findMedianSortedArrays(nums3, nums4))

// }
