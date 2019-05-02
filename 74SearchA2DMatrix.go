/*Write an efficient algorithm that searches for a 
value in an m x n matrix. This matrix has the following properties:

Integers in each row are sorted from left to right.
The first integer of each row is greater than the l
ast integer of the previous row.
Example 1:

Input:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 3
Output: true
Example 2:

Input:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 13
Output: false
*/

package main
import("fmt")

func main () {
	var matrix = [][]int {{1,  3,  5,  7},
                         {10, 11, 16, 20},
                         {23, 30, 34, 50}}
    var target int = 3

    fmt.Println(searchMatrix(matrix, target))
}

func searchMatrix(matrix [][]int, target int) bool {

    var m int = len(matrix)
    if m==0 {
        return false
    }
    
    var n int = len(matrix[0])
    if n==0 {
        return false
    }

    if m>n {
        var i int = m-1
        var j int = 0
        for j<n && i>=0 {
            if matrix[i][j]==target {
                return true
            } else if matrix[i][j]>target {
                i--
            } else if matrix[i][j]<target {
                j++
            }
        }
    } else {
        var i int = 0
        var j int = n-1
        for j>=0 && i<m {
            if matrix[i][j]==target {
                return true
            } else if matrix[i][j]>target {
                j--
            } else if matrix[i][j]<target {
                i++
            }
        }
    }

    return false

}







