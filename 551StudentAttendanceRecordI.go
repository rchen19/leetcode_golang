/*
You are given a string representing an attendance record 
for a student. The record only contains the following three characters:
'A' : Absent.
'L' : Late.
'P' : Present.
A student could be rewarded if his attendance record 
doesn't contain more than one 'A' (absent) or more than 
two continuous 'L' (late).

You need to return whether the student could be 
rewarded according to his attendance record.

Example 1:

Input: "PPALLP"
Output: True
Example 2:

Input: "PPALLL"
Output: False
*/

package main
import("fmt")

func checkRecord(s string) bool {
    var countA int = 0
    var continuousL int = 0
    for _, val := range s {
        if val=='A' {
            countA++
            continuousL = 0
        } else if val=='L' {
            continuousL++
        } else {
            continuousL = 0
        }

        if countA>1 || continuousL>2 {
            return false
        }

    }
    return true

}


func main() {
    var s string = "PPALLP"
    fmt.Println(checkRecord(s))
}
