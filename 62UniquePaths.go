/*
A robot is located at the top-left corner of a m x n grid 
(marked 'Start' in the diagram below).

The robot can only move either down or right at any point in time. 
The robot is trying to reach the bottom-right corner of the grid 
(marked 'Finish' in the diagram below).

How many possible unique paths are there?

Note: m and n will be at most 100.

Example 1:

Input: m = 3, n = 2
Output: 3
Explanation:
From the top-left corner, there are a total of 3 ways to reach 
the bottom-right corner:
1. Right -> Right -> Down
2. Right -> Down -> Right
3. Down -> Right -> Right
Example 2:

Input: m = 7, n = 3
Output: 28

*/

package main
import("fmt")

func main () {

    var m, n int = 3, 2
    
    var result int = uniquePaths(m,n)
    fmt.Println(result)

}

//using recursion
func uniquePaths1(m int, n int) int {
    if m == 1 || n == 1{
        return 1
    }
    return uniquePaths(m,n-1) + uniquePaths(m-1,n)
}

//usng DP
func uniquePaths(m int, n int) int {
    var L = make([][]int, m)

    for i:=0;i<m;i++ {
        L[i] = make([]int, n)
        L[i][0] = 1
    }

    for j:=0;j<n;j++ {
        L[0][j] = 1
    }

    for i:=1;i<m;i++ {
        for j:=1;j<n;j++ {
            L[i][j] = L[i-1][j] + L[i][j-1]
        }
    }

    return L[m-1][n-1]
}








