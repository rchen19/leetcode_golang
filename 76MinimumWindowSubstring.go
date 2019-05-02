/*
Given a string S and a string T, find the minimum window in S 
which will contain all the characters in T in complexity O(n).

Example:

Input: S = "ADOBECODEBANC", T = "ABC"
Output: "BANC"
Note:

If there is no such window in S that covers all characters in T, 
return the empty string "".
If there is such window, you are guaranteed that 
there will always be only one unique minimum window in S.
*/

package main
import("fmt")

func main () {
    var s string = "ADOBECODEBANC"
    var t string = "ABC"
    var result string = minWindow(s, t)
    fmt.Println(result)
}

func minWindow(s string, t string) string {
    if s==t {
        return s
    }
    
    var result string = ""

    var memory = make(map[byte]int)
/*
    for _, char in range t {
        if _, ok := memory[char]; ok {
            memory[char]++
        } else {
            memory[char] = 1
        }
    }
*/
    for _, char := range t {
        memory[byte(char)]++
    }

    var seen =make(map[byte]int)
    var count int = len(t)
    var i, j int = 0, 0

    for j < len(s) {
        if _, ok := memory[s[j]]; ok {

            seen[s[j]]++

            if seen[s[j]] <= memory[s[j]] {
                count--
            }

            if count == 0 {
                if len(result) == 0 || j+1-i < len(result) {
                    result = s[i:j+1]
                } 

                for i<=j {
                    i++
                    if _, ok :=memory[s[i-1]]; ok {
                        seen[s[i-1]]--
                        if seen[s[i-1]] < memory[s[i-1]] {
                            count++
                            j++
                            break
                        } else if j+1-i < len(result) {
                            
                            result = s[i:j+1]
                        }
                    } else if j+1-i < len(result) {
                        result = s[i:j+1]
                    }
                }
            } else {
                j++
            }

        } else {
            j++
        }

    }

    return result
}


























