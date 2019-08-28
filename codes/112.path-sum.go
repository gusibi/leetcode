package codes

import "fmt"

/*
 * @lc app=leetcode id=112 lang=golang
 *
 * [112] Path Sum
 *
 * https://leetcode.com/problems/path-sum/description/
 *
 * algorithms
 * Easy (37.14%)
 * Total Accepted:    292.6K
 * Total Submissions: 786.8K
 * Testcase Example:  '[5,4,8,11,null,13,4,7,2,null,null,null,1]\n22'
 *
 * Given a binary tree and a sum, determine if the tree has a root-to-leaf path
 * such that adding up all the values along the path equals the given sum.
 *
 * Note: A leaf is a node with no children.
 *
 * Example:
 *
 * Given the below binary tree and sum = 22,
 *
 *
 * ⁠     5
 * ⁠    / \
 * ⁠   4   8
 * ⁠  /   / \
 * ⁠ 11  13  4
 * ⁠/  \      \
 * 7   2      1
 *
 *
 * return true, as there exist a root-to-leaf path 5->4->11->2 which sum is 22.
 *
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getNodes(root *TreeNode) (leafs, tNodes []*TreeNode) {
	nodes := []*TreeNode{root}
	tNodes = append(tNodes, nil) // 0 位置设置nil root 在index 1 的位置 方便计算
	for {
		var cNodes []*TreeNode
		hasChild := false
		for index := range nodes {
			node := nodes[index]
			// fmt.Println("node:<<<<<", node)
			tNodes = append(tNodes, node) // 所有节点都保存
			if node == nil {
				// 如果是nil 则左右子树都设置为nil
				cNodes = append(cNodes, nil)
				cNodes = append(cNodes, nil)
			} else {
				// 如果不为nil
				cNodes = append(cNodes, node.Left)
				cNodes = append(cNodes, node.Right)
				if node.Left != nil || node.Right != nil {
					hasChild = true
				} else {
					// 如果没有子节点，说明是叶子节点
					// fmt.Println("找到了leaf", node)
					leafs = append(leafs, node)
				}
			}
		}
		if hasChild {
			// 如果还有子节点，需要继续检查
			nodes = cNodes
		} else {
			break
		}
	}
	return leafs, tNodes
}

func getPathSum(nodes []*TreeNode, index int) (sum int) {
	for index > 0 {
		node := nodes[index]
		fmt.Print("index: ", index, node)
		if node != nil {
			sum = sum + node.Val
		}
		index = index / 2
	}
	fmt.Println("sum: ", sum)
	return sum
}

func hasPathSum(root *TreeNode, sum int) bool {
	// 找到叶子节点，从下往上
	leafs, nodes := getNodes(root)
	leafsMap := make(map[*TreeNode]int)
	for index, leaf := range leafs {
		fmt.Println("map:>>>", leaf)
		leafsMap[leaf] = index
	}
	fmt.Println("nodes:>>>", nodes)
	tLength := len(nodes)
	for i := tLength - 1; i >= 0; i-- {
		node := nodes[i]
		_, ok := leafsMap[node]
		if ok {
			fmt.Println("leaf:>>", node)
			if getPathSum(nodes, i) == sum {
				return true
			}
		}
	}
	return false
}

func main() {
	values := []int{0, 5, 4, 8, 11, 0, 13, 4, 7, 2, 0, 0, 0, 0, 0, 1}
	// values := []int{0, 1, 2, 0}
	// values := []int{0, 1}
	var nodes []*TreeNode
	for index := range values {
		value := values[index]
		node := &TreeNode{Val: value}
		// fmt.Println(index, node)
		nodes = append(nodes, node)
		if index > 1 {
			pindex := index / 2
			pnode := nodes[pindex]
			if index%2 == 0 {
				pnode.Left = node
			} else {
				pnode.Right = node
			}
		}
	}
	tree := nodes[1]
	fmt.Println(tree, tree.Left, tree.Right)
	hasPathSum(tree, 22)
}
