/*
Given a collection of intervals, merge all overlapping intervals.

Example 1:

Input: [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
Example 2:

Input: [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.
NOTE: input types have been changed on April 15, 2019. 
Please reset to default code definition to get new method signature.
*/

package main
import("fmt"; "sort")

func main () {
    var intervals = [][]int {{1,3},{2,6},{8,10},{15,18}}
    fmt.Println(merge(intervals))
}

func merge(intervals [][]int) [][]int {
    var result [][]int
    //fmt.Println(len(result))
    sort.Slice(intervals, func(i,j int) bool {
        return intervals[i][0]<intervals[j][0]
    })
    for _, interval := range intervals {

        if len(result)==0 {
            result = append(result, interval)
        } else if result[len(result)-1][1] >= interval[0] && 
                    result[len(result)-1][1] < interval[1] {
            result[len(result)-1][1] = interval[1]
        } else if result[len(result)-1][1] < interval[0] {
            result = append(result, interval)
        }

    }
    return result
}