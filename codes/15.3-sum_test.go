package codes

import (
	"fmt"
	"testing"
)

func Test3Sum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	threeSum(nums)
	fmt.Println(">>>>>>>>>>>>>>>>>>>.")
	nums = []int{0, 0, 0}
	threeSum(nums)
	fmt.Println(">>>>>>>>>>>>>>>>>>>.")
	nums = []int{0, 0, 0, 0}
	threeSum(nums)
	fmt.Println(">>>>>>>>>>>>>>>>>>>.")
	nums = []int{-1, 0, 1, -2, 0, 2, 1}
	threeSum(nums)
}
