import "fmt"

/*
 * @lc app=leetcode id=12 lang=golang
 *
 * [12] Integer to Roman
 */

var SymbolValue = map[int]string{
	1:    "I",
	5:    "V",
	10:   "X",
	50:   "L",
	100:  "C",
	500:  "D",
	1000: "M",
}

func divMod(a, b int) (int, int) {
	return a / b, a % b
}

func stringMultiply(s string, n int) string {
	if n == 0 {
		return ""
	} else if n == 1 {
		return s
	} else if n == 2 {
		return fmt.Sprintf("%s%s", s, s)
	} else if n == 3 {
		return fmt.Sprintf("%s%s%s", s, s, s)
	} else {
		result := ""
		for i := 0; i < n; i++ {
			result = fmt.Sprintf("%s%s", result, s)
		}
		return result
	}
}

func intToRoman(num int) string {
	p := 1
	v, m := divMod(num, 10)
	var result string
	for v > 0 || m > 0 {
		r := ""
		if m >= 1 && m <= 3 {
			r = stringMultiply(SymbolValue[p], m)
		} else if m == 4 {
			r = fmt.Sprintf("%s%s", SymbolValue[p], SymbolValue[p*5])
		} else if m == 5 {
			r = SymbolValue[p*5]
		} else if m > 5 && m < 9 {
			r = fmt.Sprintf("%s%s", SymbolValue[p*5], stringMultiply(SymbolValue[p], (m-5)))
		} else if m == 9 {
			r = fmt.Sprintf("%s%s", SymbolValue[p], SymbolValue[p*10])
		}
		result = fmt.Sprintf("%s%s", r, result)
		v, m = divMod(v, 10)
		p = p * 10
	}
	return result
}

