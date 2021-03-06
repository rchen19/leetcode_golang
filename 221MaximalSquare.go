/*
Given a 2D binary matrix filled with 0's and 1's,
find the largest square containing only 1's and return its area.

Example:

Input:

1 0 1 0 0
1 0 1 1 1
1 1 1 1 1
1 0 0 1 0

Output: 4
*/

package main

import (
	"fmt"
)

func main() {
	var matrix [][]byte = [][]byte{{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'}}

	fmt.Println(maximalSquare(matrix))
}

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	var m int = len(matrix)
	var n int = len(matrix[0])
	var left = make([]int, n)
	var right = make([]int, n)
	for i := range right {
		right[i] = n
	}
	var heights = make([]int, n)
	var currLeft int = 0
	var currRight int = n
	var area int = 0

	var maxA int = 0
	var width int = 0

	for i := 0; i < m; i++ {
		currLeft = 0
		currRight = n

		//parse the row from left
		//calculate accumulated value from all previous rows by:
		//look at each element, increment heights array if the element is 1
		//else set the heights array value to 0
		//so at each row, the heights array shows the accumulated value from previous rows
		//also: if an element is 0, set left[j] to be 0, otherwise set left[j] to be the starting
		//index of the continues patch of value "1" which the element belongs to
		for j := 0; j < n; j++ {
			if matrix[i][j] == byte('0') {
				heights[j] = 0
				left[j] = 0
				currLeft = j + 1
			} else {
				heights[j] += 1
				if currLeft > left[j] {
					left[j] = currLeft
				}
			}
		}

		//parse the row from right
		//if an element is 0, set right[j] to be 0, otherwise set right[j] to be the ending
		//index of the continues patch of value "1" which the element belongs to
		for j := n - 1; j >= 0; j-- {
			if matrix[i][j] == byte('0') {
				right[j] = n
				currRight = j
			} else {
				if currRight < right[j] {
					right[j] = currRight
				}
			}
		}

		//parse the row from left one more time
		for j := 0; j < n; j++ {
			if right[j]-left[j] <= heights[j] {
				width = right[j] - left[j]
			} else {
				width = heights[j]
			}

			area = width * width
			if area > maxA {
				maxA = area
			}
		}
	}
	return maxA

}
