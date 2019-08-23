/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

func divmod(m, n int) (int, int) {
	return m / n, m % n
}

func isPalindrome(x int) bool {
	if x == 0 {
		return true
	} else if x < 0 || x%10 == 0 {
		return false
	}
	y := 0
	for {
		d, m := divmod(x, 10)
		y = y*10 + m
		if x == y || d == y {
			return true
		} else if y > x {
			return false
		}
		x = d
	}
	return false
}

