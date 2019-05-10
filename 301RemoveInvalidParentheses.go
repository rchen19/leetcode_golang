/*
hard
Remove the minimum number of invalid parentheses
in order to make the input string valid.
Return all possible results.

Note: The input string may contain letters
other than the parentheses ( and ).

Example 1:

Input: "()())()"
Output: ["()()()", "(())()"]
Example 2:

Input: "(a)())()"
Output: ["(a)()()", "(a())()"]
Example 3:

Input: ")("
Output: [""]
*/

package main

import (
	"fmt"
)

func removeInvalidParentheses(s string) []string {
	var out []string

	genValid(s, &out, 0, 0, []byte{'(', ')'})

	return out
}

func genValid(s string, out *[]string, last_i int, last_j int, par []byte) {
	var count int = 0
	for i := last_i; i < len(s); i++ {
		if s[i] == par[0] {
			count++
		} else if s[i] == par[1] {
			count--
		}
		if count < 0 {
			for j := last_j; j <= i; j++ {
				if s[j] == par[1] && (j == last_j || s[j-1] != par[1]) {
					genValid(s[:j]+s[j+1:], out, i, j, par)
				}
			}
			return
		}
	}
	var sReversed string = reverseString(s)
	if par[0] == '(' {
		genValid(sReversed, out, 0, 0, []byte{')', '('})
	} else {
		*out = append(*out, sReversed)
	}

}

func reverseString1(s string) string {
	var newStr = make([]rune, len(s))
	for i, char := range s {
		switch char {
		case '(':
			//fmt.Println(len(s), i)
			newStr[len(s)-1-i] = ')'
		case ')':
			newStr[len(s)-1-i] = '('
		default:
			newStr[len(s)-1-i] = char
		}
	}
	return string(newStr)
}

func reverseString(s string) string {
	if len(s) < 2 {
		return s
	}
	var newStr = make([]rune, len(s))
	for i, char := range s {
		newStr[len(s)-1-i] = char
	}
	return string(newStr)
}

func main() {
	var s string = "(a)())()"
	fmt.Println(removeInvalidParentheses(s))
}
