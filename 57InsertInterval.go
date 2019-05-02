/*
Given a set of non-overlapping intervals, 
insert a new interval into the intervals (merge if necessary).

You may assume that the intervals were 
initially sorted according to their start times.

Example 1:

Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
Output: [[1,5],[6,9]]
Example 2:

Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
Output: [[1,2],[3,10],[12,16]]
Explanation: Because the new interval [4,8] 
overlaps with [3,5],[6,7],[8,10].
NOTE: input types have been changed on April 15, 2019. 
Please reset to default code definition to get new method signature.
*/

package main
import("fmt")

func main () {
    var intervals = [][]int {{1,2},{3,5},{6,7},{8,10},{12,16}}
    var newInterval = []int {4,8}
    fmt.Println(insert(intervals, newInterval))
}


func insert(intervals [][]int, newInterval []int) [][]int {
    if len(intervals)==0 {
        return append([][]int {}, newInterval)
    }

    var newList [][]int

    for _, intvl := range intervals {

        if newInterval==nil || intvl[0] < newInterval[0] {
            newList = append(newList, intvl)
        } else {
            newList = append(append(newList, newInterval), intvl)
            newInterval = nil
        }
    }
    if newInterval!=nil {
        newList = append(newList, newInterval)
    }
    //fmt.Println(newList)
    return merge(newList)
}

func merge(intervals [][]int) [][]int {
    var result [][]int
    //fmt.Println(len(result))
    //sort.Slice(intervals, func(i,j int) bool {
    //    return intervals[i][0]<intervals[j][0]
    //})
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