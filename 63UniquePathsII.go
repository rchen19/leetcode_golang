/*
A robot is located at the top-left corner of a m x n grid 
(marked 'Start' in the diagram below).

The robot can only move either down or right at any point in time. 
The robot is trying to reach the bottom-right corner of the grid 
(marked 'Finish' in the diagram below).

Now consider if some obstacles are added to the grids. 
How many unique paths would there be?

An obstacle and empty space is marked as 1 and 0 respectively in the grid.

Note: m and n will be at most 100.

Example 1:

Input:
[
  [0,0,0],
  [0,1,0],
  [0,0,0]
]
Output: 2
Explanation:
There is one obstacle in the middle of the 3x3 grid above.
There are two ways to reach the bottom-right corner:
1. Right -> Right -> Down -> Down
2. Down -> Down -> Right -> Right
*/

package main
import("fmt")



func uniquePathsWithObstacles1(obstacleGrid [][]int) int {
    var m int = len(obstacleGrid)
    var n int = len(obstacleGrid[0])
    if obstacleGrid[0][0]==1 {
        return 0
    }
    var total_path int = uniquePaths(obstacleGrid, m, n, 0, 0)

    return total_path
}

func uniquePaths(obstacleGrid [][]int, m int, n int, x int, y int) int {
    var num_path int = 0
    if x==m-1 && y==n-1 {
        if obstacleGrid[x][y]==0 {
            num_path+=1
        }
        
    } else {
        if x+1<m && obstacleGrid[x+1][y]!=1 {
        num_path += uniquePaths(obstacleGrid, m, n, x+1, y)
        }
        if y+1<n && obstacleGrid[x][y+1]!=1 {
            num_path += uniquePaths(obstacleGrid, m, n, x, y+1)
        }
    }
    

    return num_path
}

// using dp
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    var m int = len(obstacleGrid)
    var n int = len(obstacleGrid[0])
    if obstacleGrid[0][0]==1 || obstacleGrid[m-1][n-1]==1 {
        return 0
    }
    var L = make([][]int, m)
    for k:=0;k<m;k++ {
        L[k] = make([]int, n)
    }

    for i:=m-1;i>=0;i-- {
        for j:=n-1;j>=0;j-- {
            if i==m-1 && j==n-1 {
                L[i][j] = 1
            } else if obstacleGrid[i][j]==1 {
                L[i][j]=0
            } else if i==m-1 {
                L[i][j] = L[i][j+1]
            } else if j==n-1 {
                L[i][j] = L[i+1][j]
            } else {

                L[i][j] = L[i+1][j] + L[i][j+1]
            }
            
        }
    }
    //fmt.Println(L)
    return L[0][0]
}

func main () {
    var obstacleGrid [][]int = [][]int {{0,0,0},
                                      {0,1,0},
                                      {0,0,0}}

    var result int = uniquePathsWithObstacles(obstacleGrid)
    fmt.Println(result)
}


