/*
Given n pairs of parentheses, write a function to
generate all combinations of well-formed parentheses.

For example, given n = 3, a solution set is:

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
*/

package main

import (
	"fmt"
)

func generateParenthesis(n int) []string {
	out := []string{}
	str := make([]rune, n*2)
	str[0] = '('
	str[n*2-1] = ')'

	return out
}

func main() {
	n := 3
	out = generateParenthesis(n)
	fmt.Println("[")
	for _, str := range out {
		fmt.Println("\"" + str + "\",")
	}
	fmt.Println("]")
}
