package codes

/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 * Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.
 *
 * Example 1:
 *
 * Input: "babad"
 * Output: "bab"
 * Note: "aba" is also a valid answer.
 * Example 2:
 *
 * Input: "cbbd"
 * Output: "bb"
 */

func longestPalindrome(s string) string {
	max := func(a, b int) int {
		if a >= b {
			return a
		} else {
			return b
		}
	}
	if len(s) < 1 {
		return s
	}
	index, longest := 0, 0
	for i := 0; i < len(s); i++ {
		left := getPalindrome(s, i, i)
		right := getPalindrome(s, i, i+1)
		if max(left, right) > longest {
			longest = max(left, right)
			index = i
		}
	}
	start, end := index-(longest-1)/2, index+longest/2
	results := s[start : end+1]
	return results
}

func getPalindrome(s string, l, r int) int {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return r - l - 1
}

// func main() {
// 	longestPalindrome("bb")
// 	longestPalindrome("abbb")
// 	longestPalindrome("abba")
// 	longestPalindrome("abbc")
// }
