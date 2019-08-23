/*
 * @lc app=leetcode id=42 lang=golang
 *
 * [42] Trapping Rain Water
 * 先找最高点
 * 找左侧高点
 */
// package main

// import "fmt"

func max(height []int) int{
	if len(height) == 1{
		return height[0]
	}
	m := height[0]
	for _, n := range(height){
		if n>=m{
			m = n
		}
	}
	return m
}

func min(a, b int) int{
	if (a<b){
		return a
	}else{
		return b
	}
}

func trap(height []int) int {
    if len(height) <= 2{
		return 0
	}
	// total 总水量
	// left 左侧最高点
	// right 右侧最高点
	total, left, right := 0, 0, max(height)
	for i, n := range(height){
		// fmt.Println(total, left, right)
		if n > left{
			left = n
		}
		if n == right && i!=len(height) -1{
			// 如果计算到了右侧最高点，且不是最后一个，重新计算右侧最高点
			right = max(height[i+1:])
		}
		mint := min(left, right)
		// 计算每个柱子的水量
		if n < mint{
			// fmt.Println(">>>>>>>>>>>>>>>>>", i, mint-n)
			total += mint-n
		}
	}
	return total
}

// func main(){
// 	height := []int{0,1,0,2,1,0,1,3,2,1,2,1}
// 	total := trap(height)
// 	fmt.Println(total)
// }

