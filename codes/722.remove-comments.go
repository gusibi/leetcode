/*
 * @lc app=leetcode id=722 lang=golang
 *
 * [722] Remove Comments
 */
package main

import (
	"fmt"
	"log"
	"strings"
)

func removeComments(source []string) []string {
	isBlock := false
	isInline := false
	i := 0
	newSource := make([]string, len(source))
	for _, line := range source {
		lcIndex, bcsIndex, bceIndex := strings.Index(line, "//"), strings.Index(line, "/*"), strings.Index(line, "*/")
		// fmt.Println(line, lcIndex, bcsIndex, bceIndex, isBlock)
		if isBlock {
			// 如果是在块注释中，找*/ 如果没有则忽略
			if bceIndex > -1 && bcsIndex < bceIndex-1 {
				// strings.Index(line, "/*") < strings.Index(line, "*/") -1  说明 不是 /*/
				isBlock = false
				if bceIndex+2 < len(line) {
					if isInline {
						newCode := fmt.Sprintf("%s%s", newSource[i-1], string([]rune(line)[bceIndex+2:len(line)]))
						newSource[i-1] = newCode
					} else {
						newSource[i] = string([]rune(line)[bceIndex+2 : len(line)])
						i++
					}
				}
				isInline = false
			}
			continue
		} else {
			// 找 // 和 /*
			if lcIndex > -1 && (lcIndex < bcsIndex || bcsIndex < 0) {
				if strings.Index(line, "//") == 0 {
					continue
				} else {
					// “//” 之后的删掉
					// log.Println("append: ", string([]rune(line)[0:lcIndex]))
					newSource[i] = string([]rune(line)[0:lcIndex])
					i++
				}
			} else if bcsIndex > -1 && (lcIndex > bcsIndex || lcIndex < 0) {
				isBlock = true
				if bceIndex > -1 && bcsIndex < bceIndex-1 {
					isBlock = false
				}
				if bcsIndex == 0 {
					continue
				} else {
					// “/*” 之后的删掉
					log.Println("append: ", string([]rune(line)[0:bcsIndex]))
					newSource[i] = string([]rune(line)[0:bcsIndex])
					i++
					isInline = true
					if bceIndex > -1 && bcsIndex < bceIndex-1 {
						isInline = false
					}
				}
				// fmt.Println(isBlock)
			} else {
				// log.Println("append: ", line)
				newSource[i] = line
				i++
			}
		}
	}
	newSource = newSource[0:i]
	fmt.Println("new source", newSource)
	return newSource
}

func main() {
	// source := []string{"/*Test program */", "int main()", "{ ", "  // variable declaration ", "int a, b, c;", "/* This is a test", "   multiline  ", "   comment for ", "   testing */", "a = b + c;", "}"}
	// removeComments(source)
	// source := []string{"a/*comment", "line", "more_comment*/b"}
	// removeComments(source)
	source := []string{"struct Node{", "    /*/ declare members;/**/", "    int size;", "    /**/int val;", "};"}
	removeComments(source)
}
