/*
Given a string, find the first non-repeating character in it 
and return it's index. If it doesn't exist, return -1.

Examples:

s = "leetcode"
return 0.

s = "loveleetcode",
return 2.
Note: You may assume the string contain only lowercase letters.
*/

package main
import("fmt")

func main() {
    var s string = "loveleetcode"
    fmt.Println(firstUniqChar(s))
}

func firstUniqChar1(s string) int {
    var L [26]int
    for i:=0;i<len(s);i++ {
        L[s[i]-'a'] ++
    }

    for i:=0;i<len(s);i++ {
        if L[s[i]-'a'] == 1 {return i}
    }
    return -1
    
}

func firstUniqChar2(s string) int {
    var L [26]int
    for _, val := range s {
        L[val-'a'] ++
    }

    for i, val:= range s {
        if L[val-'a'] == 1 {return i}
    }
    return -1
    
}

func firstUniqChar(s string) int {
    var D map[rune]int = make(map[rune]int)
    for _, val := range s {
        D[val] += 1
    }

    for i, val := range s {
        if D[val] == 1 {
            return i
        }
    }
    return -1
}




