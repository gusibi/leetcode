package codes

import "fmt"

/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */
var NumberLetter = map[string]string{
	"2": "abc", "3": "def", "4": "ghi",
	"5": "jkl", "6": "mno", "7": "pqrs",
	"8": "tuv", "9": "wxyz",
}

func _letterCombinations(combinations, digits []string) []string {
	if len(digits) == 0 {
		return combinations
	}
	var results []string
	letters := NumberLetter[digits[0]]
	// fmt.Println(letters)
	if len(combinations) == 0 {
		for _, letter := range []byte(letters) {
			results = append(results, string(letter))
		}
	} else {
		for _, combination := range combinations {
			for _, letter := range []byte(letters) {
				c := fmt.Sprintf("%s%s", combination, string(letter))
				results = append(results, c)
			}
		}
	}
	combinations = _letterCombinations(results, digits[1:])
	return combinations
}

func letterCombinations(digits string) []string {
	var digitsArray, combinations []string
	for _, digit := range []byte(digits) {
		digitsArray = append(digitsArray, string(digit))
	}
	combinations = _letterCombinations(combinations, digitsArray)
	return combinations
}

// func main(){
// 	fmt.Println(letterCombinations("1"))
// }
