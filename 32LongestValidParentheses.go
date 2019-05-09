/*
Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed) parentheses substring.

Example 1:

Input: "(()"
Output: 2
Explanation: The longest valid parentheses substring is "()"
Example 2:

Input: ")()())"
Output: 4
Explanation: The longest valid parentheses substring is "()()"
*/

package main

import (
	"fmt"
)

//slightly improved version compared to the one below
func longestValidParentheses(s string) int {
	var stack []int
	//use a slice of length len(s)
	//to mark if each index is paired
	var flag []bool = make([]bool, len(s))

	for i, char := range s {
		if char == '(' {
			//push index of '(' into stack
			stack = append(stack, i)
		} else if len(stack) > 0 {
			//pop the stack when seeing ')'
			//(popped index, current index) represents the start and end
			//indices of paired '(' and ')
			flag[stack[len(stack)-1]] = true //index of paired ')'
			flag[i] = true                   //index of paired ')'
			stack = stack[:len(stack)-1]     //pop the stack
		}
	}

	//loop through the flag slice
	//to find longest continuously flagged indices
	var maxLength int = 0
	var curLength int = 0
	for _, yes := range flag {
		if yes {
			curLength++
		} else {
			if curLength > maxLength {
				maxLength = curLength
			}
			curLength = 0
		}
	}
	if curLength > maxLength {
		maxLength = curLength
	}
	return maxLength
}

//much faster using stack
func longestValidParentheses1(s string) int {
	var stack []int
	var segList [][2]int
	//a seg is a pair of int representing start and end

	for i, char := range s {
		if char == '(' {
			//push index of '(' into stack
			stack = append(stack, i)
		} else if len(stack) > 0 {
			//pop the stack when seeing ')'
			//also push (popped index, current index) into segList
			//(popped index, current index) represents the start and end
			//indices of paired '(' and ')'
			segList = append(segList, [2]int{stack[len(stack)-1], i})
			stack = stack[:len(stack)-1] //pop the stack
		}
	}
	if len(segList) == 0 {
		return 0
	} //there is no paired '(' and ')'

	//use a slice of length len(s)
	//to mark if each index is paired
	var flag []bool = make([]bool, len(s))
	for _, seg := range segList {
		flag[seg[0]] = true //index of paired ')'
		flag[seg[1]] = true //index of paired ')'
	}

	//loop through the flag slice
	//to find longest continuously flagged indices
	var maxLength int = 0
	var curLength int = 0
	for _, yes := range flag {
		if yes {
			curLength++
		} else {
			if curLength > maxLength {
				maxLength = curLength
			}
			curLength = 0
		}
	}
	if curLength > maxLength {
		maxLength = curLength
	}
	return maxLength
}

//this one is very slow
func longestValidParentheses2(s string) int {
	if len(s) == 0 {
		return 0
	}
	//change input string to a slice of int, '(' -> +1, ')' -> -1
	w := make([]int, len(s))
	for i, char := range s {
		switch char {
		case '(':
			w[i] = 1
		case ')':
			w[i] = -1
		}
	}

	var l int = 0
	var i, j int
	var t int

	for k := 0; k < len(s); k++ {
		i, j = k, k
		t = w[k]
		//fmt.Println(w[k:])
		for i <= j && j < len(s) {
			if t > 0 {
				j++
				if j < len(s) {
					t += w[j]
				}
			} else if t < 0 {
				if j == i {
					j++
				}
				if i < len(s) {
					t -= w[i]
				}
				i++
				if j < len(s) {
					t += w[j]
				}

			} else {
				if j-i+1 > l {
					l = j - i + 1
				}
				j++
				if j < len(s) {
					t += w[j]
				}

			}
		}
	}

	return l
}

func main() {
	s := "()(()"
	fmt.Println(longestValidParentheses(s))
}
