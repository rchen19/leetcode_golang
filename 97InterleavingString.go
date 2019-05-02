/*
Given s1, s2, s3, find whether s3 is formed by the interleaving of s1 and s2.

Example 1:

Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
Output: true
Example 2:

Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
Output: false
*/
package main
import("fmt")

/*
func isInterleave(s1 string, s2 string, s3 string) bool {

    if len(s3) != len(s1) + len(s2) {
        return false
    } else if len(s3) == 0 {
        return true
    }


    var i, j, k int = 0, 0, 0
    var l, m, n int = len(s1), len(s2), len(s3)
    var output bool = search(s1, s2, s3, i, j, k, l, m, n)
    return output
}

func search(s1 string, s2 string, s3 string, i int, j int, k int, l int, m int, n int) bool {
    var result bool
    if k == n - 1 {
        if i < l && j == m {
            result = (s3[k] == s1[i])
        } else if i == l && j < m {
            result =  (s3[k] == s2[j])
        } else {
            result =  false
        }
    } else {

        if i < l && j < m {
            if s3[k] == s1[i] && s3[k] == s2[j] {
                result =  search(s1, s2, s3, i+1, j, k+1, l, m, n) || search(s1, s2, s3, i, j+1, k+1, l, m, n)
            } else if s3[k] == s1[i] {
                result =  search(s1, s2, s3, i+1, j, k+1, l, m, n)
            } else if s3[k] == s2[j] {
                result =  search(s1, s2, s3, i, j+1, k+1, l, m, n)
            } else {
                result =  false
            }
        } else if i < l {
            if s3[k] == s1[i] {
                result =  search(s1, s2, s3, i+1, j, k+1, l, m, n)
            }
        } else if j < m {
            if s3[k] == s2[j] {
                result =  search(s1, s2, s3, i, j+1, k+1, l, m, n)
            }
        } else {
            result = false
        }
    }

    return result
}
*/


func isInterleave(s1 string, s2 string, s3 string) bool {

    if len(s3)==0 {
        if len(s1)==0 && len(s2)==0 {
            return true
        } else {
            return false
        }
    } else if len(s3) == 1 {
        if len(s1)==1 && len(s2)==0 {
            return s3 == s1
        } else if len(s1)==0 && len(s2)==1 {
            return s3 == s2
        } else {
            return false
        }
    } else if len(s1)>0 && len(s2)==0 {
        if s3[0]==s1[0] {
            return isInterleave(s1[1:], s2, s3[1:])
        } else {
            return false
        }
    } else if len(s2)>0 && len(s1)==0 {
        if s3[0]==s2[0] {
            return isInterleave(s1, s2[1:], s3[1:])
        } else {
            return false
        }
    } else {
        if s3[0]==s1[0] && s3[0]==s2[0] {
            return isInterleave(s1[1:], s2, s3[1:]) || isInterleave(s1, s2[1:], s3[1:])
        } else if s3[0]==s1[0] {
            return isInterleave(s1[1:], s2, s3[1:])
        } else if s3[0]==s2[0] {
            return isInterleave(s1, s2[1:], s3[1:])
        } else {
            return false
        }
    }

}

func main() {

    var s1 string = "aabcc"
    var s2 string = "dbbca"
    var s3 string = "aadbbcbcac"
 
    var result = isInterleave(s1, s2, s3)
    fmt.Println(result)

}