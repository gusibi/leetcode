package codes

// import "fmt"
/*
 * @lc app=leetcode id=77 lang=golang
 *
 * [77] Combinations
 */

func combineList(ns []int, k int) [][]int {
	if len(ns) == k {
		return [][]int{ns}
	}
	results := [][]int{}
	for i := 0; i < len(ns); i++ {
		if k == 1 {
			results = append(results, []int{ns[i]})
		} else {
			subResults := combineList(ns[i+1:len(ns)], k-1)
			for _, r := range subResults {
				result := []int{ns[i]}
				result = append(result, r...)
				results = append(results, result)
			}
		}
	}
	return results
}
func combine(n int, k int) [][]int {
	var ns = []int{}
	for i := 1; i <= n; i++ {
		ns = append(ns, i)
	}
	return combineList(ns, k)
}

// func main(){
// 	fmt.Println(combine(3, 2))
// 	fmt.Println(combine(3, 3))
// 	fmt.Println(combine(4, 2))
// 	fmt.Println(combine(4, 3))
// }
