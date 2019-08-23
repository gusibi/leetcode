/*
 * @lc app=leetcode id=179 lang=golang
 *
 * [179] Largest Number
 */
import (
	"fmt"
	"strconv"
	"strings"
)

func mergeSort(A []int)[]int{
	if len(A) <=1{
		return A
	}
	mid := len(A) / 2
	left := mergeSort(A[:mid])
	right := mergeSort(A[mid:])
	result := merge(left, right)
	return result
}

func largeThan(a, b int) bool{
	ab := fmt.Sprintf("%d%d", a, b)
	ba := fmt.Sprintf("%d%d", b, a)
	abInt, _ := strconv.Atoi(ab)
	baInt, _ := strconv.Atoi(ba)
	// fmt.Println(abInt, baInt)
	if abInt >= baInt{
		return true
	}else{
		return false
	}
}

func merge(left, right []int) []int{
	results := []int{}
    i, j:=0, 0
	for i<len(left) && j<len(right) {
		// fmt.Println(largeThan(left[i], right[j]) == false)
		if largeThan(left[i], right[j]) == true{
			results = append(results, left[i])
			i++
		}else{
			results = append(results, right[j])
			j++
		}
	}
	results = append(results, left[i:]...)
	results = append(results, right[j:]...)
	return results
}
func largestNumber(nums []int) string {
	// 修改比较方法，把字符大的返回
	results := mergeSort(nums)
	// fmt.Println(results)
	resultsText := []string{}
	for i := range results {
        number := results[i]
        text := strconv.Itoa(number)
        resultsText = append(resultsText, text)
    }

    // Join our string slice.
	result := strings.Join(resultsText, "")
	if strings.HasPrefix(result, "0"){
		result = "0"
	}
	return result
}

