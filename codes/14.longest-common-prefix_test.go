package codes

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	data := []string{"flower", "flow", "flight"}
	t.Log(longestCommonPrefix(data))
}
