// package main

// import (
// 	"fmt"
// 	"strconv"
// )
/*
Given a string of numbers and operators, 
return all possible results from computing all the 
different possible ways to group numbers and operators. 
The valid operators are +, - and *.

Example 1:
    Input: "2-1-1"
	Output: [0, 2]

Explanation: 
    ((2-1)-1) = 0 
    (2-(1-1)) = 2

Example 2:
    Input: "2*3-4*5"
    Output: [-34, -14, -10, -10, 10]
Explanation: 
    (2*(3-(4*5))) = -34 
    ((2*3)-(4*5)) = -14 
    ((2*(3-4))*5) = -10 
    (2*((3-4)*5)) = -10 
    (((2*3)-4)*5) = 10
 * @lc app=leetcode id=241 lang=golang
 *
 * [241] Different Ways to Add Parentheses
 */

func diffWaysToCompute(input string) []int {
	ret := make([]int, 0)
    for i:=0; i < len(input); i++ {
        if input[i] == '-' || input[i] == '*' || input[i] == '+' {
			// 如果是操作符，分别处理操作符两边
			part1 := input[0:i]
			part2 := input[i+1:]
			part1Ret := diffWaysToCompute(part1)
			part2Ret := diffWaysToCompute(part2)
			for _, p1 := range part1Ret{
				for _, p2 := range part2Ret{
					c := 0
					switch (input[i]){
					case '+':
						c = p1 + p2
					case '-':
						c = p1 - p2
					case '*':
						c = p1 * p2
					}
					ret = append(ret, c)
				}
			}
		} else{
			if len(input) == 1{
                x, _ := strconv.Atoi(input)
				ret = append(ret, x)
			}
		}
	}
	return ret
}

// func main(){
// 	input := "2-1-1"
// 	fmt.Println(diffWaysToCompute(input))
// 	input = "2*3-4*5"
// 	fmt.Println(diffWaysToCompute(input))
// }

