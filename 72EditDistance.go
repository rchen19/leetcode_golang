/*
Given two words word1 and word2, find the minimum number of 
operations required to convert word1 to word2.

You have the following 3 operations permitted on a word:

Insert a character
Delete a character
Replace a character
Example 1:

Input: word1 = "horse", word2 = "ros"
Output: 3
Explanation: 
horse -> rorse (replace 'h' with 'r')
rorse -> rose (remove 'r')
rose -> ros (remove 'e')
Example 2:

Input: word1 = "intention", word2 = "execution"
Output: 5
Explanation: 
intention -> inention (remove 't')
inention -> enention (replace 'i' with 'e')
enention -> exention (replace 'n' with 'x')
exention -> exection (replace 'n' with 'c')
exection -> execution (insert 'u')
*/

package main
import("fmt")

// return the min of two integers
func min2(a, b int) int {
    if a<=b {
        return a
    } else {
        return b
    }
}

// DP
func minDistance(word1 string, word2 string) int {
    var m int = len(word1)
    var n int = len(word2)
    if m==0 {
        return n
    }

    if n==0 {
        return m
    }

    var L = make([][]int, m+1)
    for i:=0; i<m+1; i++ {
        L[i] = make([]int, n+1)
    }

    for k:=0; k<m+1; k++ {
        L[k][0] = k
    }
    for l:=0; l<n+1; l++ {
        L[0][l] = l
    }


    for k:=1; k<m+1; k++ {
        for l:=1; l<n+1; l++ {
            //fmt.Println(k,l)
            if word1[k-1]==word2[l-1] {

                L[k][l] = L[k-1][l-1]
            } else {

                L[k][l] = min2(L[k-1][l-1]+1, min2(L[k-1][l], L[k][l-1]) + 1)

            }

        }
    }
    //fmt.Println(L[m-1][n-1])

    return L[m][n]
 
}

func main() {

    var word1 string = "intention"
    var word2 string = "execution"
 
    var result int = minDistance(word1, word2)
    fmt.Println(result)

}














