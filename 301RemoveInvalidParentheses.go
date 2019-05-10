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

//using BFS iteration
func removeInvalidParentheses(s string) []string {
	var out []string
	var currNode node
	var count int
	var toPaseRightToLeft bool
	var goodToGo bool
	//use a queue to keep the nodes to be explored
	//first in first out to achieve BFS
	var queue []node = []node{node{s, 0, 0, true}}
	for len(queue) > 0 {
		//pop the last node from the stack
		currNode = queue[0]
		queue = queue[1:]
		if currNode.leftToRight {
			//parse left->right, check for excess ')'
			toPaseRightToLeft = true
			count = 0
			for i := currNode.i; i < len(currNode.s); i++ {
				switch currNode.s[i] {
				case '(':
					count++
				case ')':
					count--
				}
				//when never count<0, need to remove a ')'
				if count < 0 {
					toPaseRightToLeft = false
					//since there exist excess ')', no need to parse right->left
					for j := currNode.j; j <= i; j++ {
						if currNode.s[j] == ')' && (j == currNode.j || currNode.s[j-1] != ')') {
							queue = append(queue, node{currNode.s[:j] + currNode.s[j+1:], i, j, currNode.leftToRight})
						}
					}
					break
				}
			}
			//if looks good from left to right
			//push right to left node to the stack
			//finished parsing from left to right
			if toPaseRightToLeft {
				queue = append(queue, node{currNode.s, len(currNode.s) - 1, len(currNode.s) - 1, false})
			}
		} else {
			goodToGo = true
			count = 0
			//parse from right to left
			//check for excess '('
			for i := currNode.i; i >= 0; i-- {
				switch currNode.s[i] {
				case '(':
					count--
				case ')':
					count++
				}
				//when never count<0, need to remove a ')'
				if count < 0 {
					goodToGo = false
					for j := currNode.j; j >= i; j-- {
						if currNode.s[j] == '(' && (j == currNode.j || currNode.s[j+1] != '(') {
							queue = append(queue, node{currNode.s[:j] + currNode.s[j+1:], i - 1, j - 1, currNode.leftToRight})
						}
					}
					break
				}

			}
			//finished parsing from right to left
			if goodToGo {
				out = append(out, currNode.s)
			}

		}

	}

	return out
}

//using DFS iteration
func removeInvalidParentheses1(s string) []string {
	var out []string
	var currNode node
	var count int
	var toPaseRightToLeft bool
	var goodToGo bool
	//use a stack to keep the nodes to be explored
	//last in first out to achieve DFS
	var stack []node = []node{node{s, 0, 0, true}}
	for len(stack) > 0 {
		//pop the last node from the stack
		currNode = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if currNode.leftToRight {
			//parse left->right, check for excess ')'
			toPaseRightToLeft = true
			count = 0
			for i := currNode.i; i < len(currNode.s); i++ {
				switch currNode.s[i] {
				case '(':
					count++
				case ')':
					count--
				}
				//when never count<0, need to remove a ')'
				if count < 0 {
					toPaseRightToLeft = false
					//since there exist excess ')', no need to parse right->left
					for j := currNode.j; j <= i; j++ {
						if currNode.s[j] == ')' && (j == currNode.j || currNode.s[j-1] != ')') {
							stack = append(stack, node{currNode.s[:j] + currNode.s[j+1:], i, j, currNode.leftToRight})
						}
					}
					break
				}
			}
			//if looks good from left to right
			//push right to left node to the stack
			//finished parsing from left to right
			if toPaseRightToLeft {
				stack = append(stack, node{currNode.s, len(currNode.s) - 1, len(currNode.s) - 1, false})
			}
		} else {
			goodToGo = true
			count = 0
			//parse from right to left
			//check for excess '('
			for i := currNode.i; i >= 0; i-- {
				switch currNode.s[i] {
				case '(':
					count--
				case ')':
					count++
				}
				//when never count<0, need to remove a ')'
				if count < 0 {
					goodToGo = false
					for j := currNode.j; j >= i; j-- {
						if currNode.s[j] == '(' && (j == currNode.j || currNode.s[j+1] != '(') {
							stack = append(stack, node{currNode.s[:j] + currNode.s[j+1:], i - 1, j - 1, currNode.leftToRight})
						}
					}
					break
				}

			}
			//finished parsing from right to left
			if goodToGo {
				out = append(out, currNode.s)
			}

		}

	}

	return out
}

//define a custom type to store strings and indices and direction to push to stack
type node struct {
	s           string
	i           int
	j           int
	leftToRight bool
}

//using DFS recursion
func removeInvalidParentheses2(s string) []string {
	var out []string

	//genValid(s, &out, 0, 0, []byte{'(', ')'})
	genValid1(s, &out, 0, 0, true)

	return out
}

func genValid1(s string, out *[]string, last_i int, last_j int, leftToRight bool) {
	//set a counter, ++ when seeing '(', and -- when seeing ')'
	var count int = 0
	for i := last_i; i < len(s); i++ {
		if s[i] == '(' {
			count++
		} else if s[i] == ')' {
			count--
		}
		//when never count<0, need to remove a ')'
		if count < 0 {
			for j := last_j; j <= i; j++ {
				//find ')' between index last_j->i
				//only try ')' whose previous character is not ')'
				//to avoid duplicate results
				if s[j] == ')' && (j == last_j || s[j-1] != ')') {
					//recursively find valid on s with ')' removed
					genValid1(s[:j]+s[j+1:], out, i, j, leftToRight)
				}
			}
			return
		}
	}

	//if there do not exist excess ')'
	//check to see if there exist excess '('
	//by traversing s from right to left
	var sReversed string = reverseString1(s)

	if leftToRight {
		//finished left to right
		//now start right to left
		genValid1(sReversed, out, 0, 0, false)
	} else {
		//finished right to left
		*out = append(*out, sReversed)
	}
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

//this function reverse the string
//also swap '(' and ')'
func reverseString1(s string) string {
	var newStr = make([]rune, len(s))
	for i, char := range s {
		switch char {
		case '(':
			newStr[len(s)-1-i] = ')'
		case ')':
			newStr[len(s)-1-i] = '('
		default:
			newStr[len(s)-1-i] = char
		}
	}
	return string(newStr)
}

//this function reverse the string
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
	var s string = "(()a"
	fmt.Println(removeInvalidParentheses(s))
}
