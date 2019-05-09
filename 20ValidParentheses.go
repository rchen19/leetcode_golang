/*
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.

Example 1:

Input: "()"
Output: true
Example 2:

Input: "()[]{}"
Output: true
Example 3:

Input: "(]"
Output: false
Example 4:

Input: "([)]"
Output: false
Example 5:

Input: "{[]}"
Output: true
*/

package main
import("fmt")

func isValid(s string) bool {
    var stack []rune
    for _, char := range s {
        switch char {
        case '(', '{', '[':
            stack = append(stack, char)
        case ')':
            if len(stack)<1||stack[len(stack)-1]!='(' {
                return false
            } else {
                stack = stack[:len(stack)-1]
            }
        case '}':
            if len(stack)<1||stack[len(stack)-1]!='{' {
                return false
            } else {
                stack = stack[:len(stack)-1]
            }
        case ']':
            if len(stack)<1||stack[len(stack)-1]!='[' {
                return false
            } else {
                stack = stack[:len(stack)-1]
            }
        }
    }
    if len(stack)>0 {
    	return false
    }
    return true
}

func main() {
    var s string = "()[]{}"
    fmt.Println(isValid(s))
}