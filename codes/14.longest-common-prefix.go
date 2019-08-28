package codes

/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	n := -1
	cont := true
	for cont {
		n++
		for _, str := range strs {
			// fmt.Println(len(str), n)
			if len(str) == n {
				return str
			} else if str[n] != strs[0][n] {
				cont = false
				break
			}
		}
	}
	// fmt.Println(string(strs[0][:n]))
	return string(strs[0][:n])
}
