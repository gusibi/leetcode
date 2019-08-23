package main

// 对于一组不同重量、不可分割的物品，我们需要选择一些装入背包，
// 在满足背包最大重量限制的前提下，背包中物品总重量的最大值是多少？

import "fmt"

func knapsack(weight []int, n, w int) int {
	// weight: 物品重量列表, n: 物品个数
	var status [][]bool
	status = make([][]bool, n)
	for i := 0; i < n; i++ {
		status[i] = make([]bool, w+1)
	}
	status[0][0] = true
	if weight[0] <= w {
		// 如果第一个物品重量不大于总重，可以填入
		status[0][weight[0]] = true
	}
	// 在下面循环中，i为物品数量，j为背包中的物品重量
	for i := 1; i < n; i++ { // i=0 已经设置好 i<n 是因为在循环内要再加一个
		// for 循环其实是在找 第i-1 个放入后所以可能出现的结果
		for j := 0; j <= w; j++ { // 不把第i个物品放入背包
			if status[i-1][j] == true {
				// 因为这里不把第i个放入，所以重量还是和 放入第i-1 个时相同，所以判断第i-1 个有没有放入
				status[i][j] = status[i-1][j]
			}
		}
		for j := 0; j <= w-weight[i]; j++ { // 把第i个放入背包
			// 因为第i个被放入了背包，所以最大重量必须- 第i个物品的重量
			if status[i-1][j] == true {
				// 第i个放入了背包，所以第i个放入后的重量为 第i-1 个的重量 + 第i个的重量（weight[i])
				status[i][j+weight[i]] = true
			}
		}
	}
	fmt.Println(status)
	for i := w; i >= 0; i-- {
		if status[n-1][i] == true {
			return i
		}
	}
	return 0
}

func main() {
	weight := []int{2, 4, 3, 5}
	fmt.Println("start")
	fmt.Println(knapsack(weight, 4, 10))
}
