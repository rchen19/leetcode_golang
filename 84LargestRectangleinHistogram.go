/*
Given n non-negative integers representing the histogram's 
bar height where the width of each bar is 1, find the area of 
largest rectangle in the histogram.

[img]

Above is a histogram where width of each bar is 1, given height = [2,1,5,6,2,3].
[img]

The largest rectangle is shown in the shaded area, which has area = 10 unit.

Example:

Input: [2,1,5,6,2,3]
Output: 10
*/

package main
import("fmt")

func main() {
    var heights []int = []int {2,1,5,6,2,3}
    fmt.Println(largestRectangleArea(heights))
}

func largestRectangleArea(heights []int) int {
    var maxA int = 0
    var stack []int = []int {-1}
    var last int
    var prev int
    var area int
    //var index int
    //var i int = 0
    //fmt.Println(len(stack))
    for i:=0;i<=len(heights);i++ {

        for stack[len(stack)-1]!=-1  && (i==len(heights) || heights[i]<heights[stack[len(stack)-1]]) {
            last = stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            if len(stack)>0 {
                prev = stack[len(stack)-1]
            } else {
                prev = -1
            }
            area = heights[last] * (i - 1 - prev)
            if area>maxA {
                maxA = area
            }
        }
        stack = append(stack, i)
    }
    return maxA
}

func largestRectangleArea1(heights []int) int {
    var stack []int
    var maxA int = 0
    var last int
    var prev int
    var area int

    for i, h := range heights {
        for len(stack)>0 && h<heights[stack[len(stack)-1]] {
            last = stack[len(stack)-1]
            stack = stack[:len(stack)-1]

            if len(stack)>0 {
                prev = stack[len(stack)-1]
            } else {
                prev = -1
            }

            area = heights[last] * (i - 1 - prev)
            if area>maxA {
                maxA = area
            }
        }

        stack = append(stack, i)
    }

    var j = len(heights)
    for len(stack)>0 {
        last = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if len(stack)>0 {
            prev = stack[len(stack)-1]
        } else {
            prev = -1
        }

        area = heights[last] * (j - 1 - prev)
        if area>maxA {
            maxA = area
        }
    }
    return maxA
}