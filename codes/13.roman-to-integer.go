package codes

/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */
var RomanValue = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	prev, sum := 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		num := RomanValue[s[i]]
		if num >= prev {
			sum = sum + num
		} else {
			sum = sum - num
		}
		prev = num
	}
	return sum
}
