package main
import "fmt"
/*
Given an array of citations (each citation is a non-negative integer) of a researcher, 
write a function to compute the researcher's h-index.
给定一位研究者论文被引用次数的数组（被引用次数是非负整数）。编写一个方法，计算出研究者的 h 指数。

According to the definition of h-index on Wikipedia: 
"A scientist has index h if h of his/her N papers have at least h citations each, 
and the other N − h papers have no more than h citations each."
h 指数的定义: “h 代表“高引用次数”（high citations），
一名科研人员的 h 指数是指他（她）的 （N 篇论文中）至多有 h 篇论文分别被引用了至少 h 次。
（其余的 N - h 篇论文每篇被引用次数不多于 h 次。）”

Example:

Input: citations = [3,0,6,1,5]
Output: 3 

Explanation: 
[3,0,6,1,5] means the researcher has 5 papers in total and each of them had 
received 3, 0, 6, 1, 5 citations respectively. 
Since the researcher has 3 papers with at least 3 citations each and the remaining 
two with no more than 3 citations each, her h-index is 3.

给定数组表示研究者总共有 5 篇论文，每篇论文相应的被引用了 3, 0, 6, 1, 5 次。
由于研究者有 3 篇论文每篇至少被引用了 3 次，其余两篇论文每篇被引用不多于 3 次，所以她的 h 指数是 3。

说明: 如果 h 有多种可能的值，h 指数是其中最大的那个。
 * @lc app=leetcode id=274 lang=golang
 *
 * [274] H-Index
 */

func partition(A []int, p, r int)int{
	pivot := A[r] // 分区点设为最后一个元素
	i := p
	for j := p; j<= r-1; j++{
		if A[j] < pivot{
			A[i], A[j] = A[j], A[i]
			i = i + 1
		}
	}
	A[i], A[r] = A[r], A[i]
	return i
}

func quickSort(citations []int, p, r int) []int{
	// 使用快排
	// p, r := 0, len(citations) - 1 //p,r 为下标，开始时p从0 开始，r 从n-1开始
	// fmt.Println(citations)
	if p >= r{
		return citations
	}
	q := partition(citations, p, r) // 获取分区点
	quickSort(citations, p, q-1)
	quickSort(citations, q+1, r)
	return citations
}
func hIndex(citations []int) int {
	n := len(citations)
	citations = quickSort(citations, 0, n-1) 
	// fmt.Println(citations)
	for i:= 0; i<n; i++{
		if citations[i] >= n - i{
			return n-i
		}
	}
	return 0
}

func hIndex1(citations []int) int {
	m, max := make(map[int]int), 0
	for _, v := range citations {
		m[v]++
	}
	fmt.Println(m)
	for i, h := 0, len(citations); i <= h; i++ {
		max = i
		fmt.Println(i,h, m[i])
		h = h - m[i]
	}
	return max
}

func main(){
	h := hIndex1([]int{3, 0, 6, 1, 5, 7, 8})
	fmt.Println(h)
}

