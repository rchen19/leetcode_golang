/*
Given two strings A and B of lowercase letters, return true if and only if we can swap two letters in A so that the result equals B.

 

Example 1:

Input: A = "ab", B = "ba"
Output: true
Example 2:

Input: A = "ab", B = "ab"
Output: false
Example 3:

Input: A = "aa", B = "aa"
Output: true
Example 4:

Input: A = "aaaaaaabc", B = "aaaaaaacb"
Output: true
Example 5:

Input: A = "", B = "aa"
Output: false
 

Note:

0 <= A.length <= 20000
0 <= B.length <= 20000
A and B consist only of lowercase letters.
*/

package main
import("fmt")


func buddyStrings(A string, B string) bool {
    if len(A)!=len(B) {
        return false
    }

    if A==B {
        //check if repeating char exists
        var L = make(map[rune]int)
        for _, char := range A {
            if _, ok := L[char]; ok {
                return true
            } else {
                L[char] += 1
            }
        }
        return false
    }

    var L []int

    for i:=0;i<len(A);i++ {
        if A[i]!=B[i] {
            L = append(L, i)
        }
    }

    if len(L)==2 && A[L[0]]==B[L[1]] && A[L[1]]==B[L[0]] {
        return true

    } else {
        return false
    }
}

func main() {
    A := "aaaaaaa"
    B := "aaaaaaa"
    fmt.Println(buddyStrings(A, B))
}