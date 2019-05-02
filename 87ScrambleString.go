/*
Given a string s1, we may represent it as a binary tree by 
partitioning it to two non-empty substrings recursively.

Below is one possible representation of s1 = "great":

    great
   /    \
  gr    eat
 / \    /  \
g   r  e   at
           / \
          a   t
To scramble the string, we may choose any non-leaf node 
and swap its two children.

For example, if we choose the node "gr" and swap its two children, 
it produces a scrambled string "rgeat".

    rgeat
   /    \
  rg    eat
 / \    /  \
r   g  e   at
           / \
          a   t
We say that "rgeat" is a scrambled string of "great".

Similarly, if we continue to swap the children of nodes 
"eat" and "at", it produces a scrambled string "rgtae".

    rgtae
   /    \
  rg    tae
 / \    /  \
r   g  ta  e
       / \
      t   a
We say that "rgtae" is a scrambled string of "great".

Given two strings s1 and s2 of the same length, determine 
if s2 is a scrambled string of s1.

Example 1:

Input: s1 = "great", s2 = "rgeat"
Output: true
Example 2:

Input: s1 = "abcde", s2 = "caebd"
Output: false
*/

package main
import("fmt"; "sort"; "strings")

func main() {
    var s1, s2 string = "great", "rgeat"
    fmt.Println(isScramble(s1, s2))

}

//without recursion
func isScramble(s1 string, s2 string) bool {
    //if s1==nil || s2==nil {return false}

    if s1==s2 {return true}

    var m, n int = len(s1), len(s2)
    if m!=n {return false}
    //fmt.Println(m,n)

    var L = make([][][]bool, m)
    for i:=0;i<m;i++ {
        L[i] = make([][]bool, m)
        for j:=0;j<m;j++ {
            //fmt.Println(i,j)
            L[i][j] = make([]bool, m+1)
            L[i][j][1] = (s1[i]==s2[j])
        }
    }

    for k:=2;k<=m;k++ {
        for i:=0;i<=m-k;i++ {
            for j:=0;j<=m-k;j++ {
                L[i][j][k] = false
                for p:=1;p<k;p++ {
                    if ((L[i][j][p] && L[i+p][j+p][k-p]) || 
                                (L[i][j+k-p][p] && L[i+p][j][k-p])) {
                            L[i][j][k] = true
                        }
                }
            }
        }
    }
    return L[0][0][m]


}

//with memoization
func isScramble1(s1 string, s2 string) bool {
    var m = make(map[[2]string]bool)
    return checkScramble(s1, s2, m)

}

//keep record in m
func checkScramble(s1 string, s2 string, m map[[2]string]bool) bool {
    if s1==s2 {
        m[[2]string{s1,s2}] = true
        return true
    }

    //if len(s1)!=len(s2) || sortString(s1)!=sortString(s2) {
    if !containSameChar(s1, s2) {
        m[[2]string{s1,s2}] = false
        return false
    }

    //var n int = len(s1)

    if len(s1)==1 && s1!=s2 {
        m[[2]string{s1,s2}] = false
        return false
    }

    if val, ok := m[[2]string{s1,s2}]; ok {
        return val
    }

    if val, ok := m[[2]string{s1,s2}]; ok {
        return val
    }

    var output bool = false

    for i:=1;i<len(s1);i++ {

        output = 
                ((isScramble(s1[:i], s2[:i])&&isScramble(s1[i:], s2[i:])) || 
                (isScramble(s1[:i], s2[len(s2)-i:])&&isScramble(s1[i:], s2[:len(s2)-i])))
    
        
        if output {
            break
        }
    }

    m[[2]string{s1,s2}] = output
    return output

}

//without memoization
func isScramble2(s1 string, s2 string) bool {

    if s1==s2 {
        return true
    }

    //if len(s1)!=len(s2) || sortString(s1)!=sortString(s2) {
    if !containSameChar(s1, s2) {

        return false
    }

    //var n int = len(s1)
    
    if len(s1)==1 && s1!=s2{
        return false
    }

    var output bool = false

    for i:=1;i<len(s1);i++ {

        output = 
                ((isScramble(s1[:i], s2[:i])&&isScramble(s1[i:], s2[i:])) || 
                (isScramble(s1[:i], s2[len(s2)-i:])&&isScramble(s1[i:], s2[:len(s2)-i])))
    
        
        if output {
            break
        }
    }

    
    return output

}

func sortString(w string) string {
    var s []string = strings.Split(w, "")
    sort.Strings(s)
    return strings.Join(s, "")
}

func containSameLetters(w1 string, w2 string) bool {
    if len(w1)!=len(w2) {
        return false
    }
    var count [26]int
    for i:=0;i<len(w1);i++ {
        count[w1[i]-'a']++
        count[w2[i]-'a']--
    }
    for i:=0;i<26;i++ {
        if count[i]!=0 {
            return false
        }
    }
    return true
}

func containSameChar(w1 string, w2 string) bool {
    if len(w1)!=len(w2) {
        return false
    }
    var count = make(map[byte]int)
    for i:=0;i<len(w1);i++ {
        count[w1[i]]++
        count[w2[i]]--
    }

    for _, v := range count {
        if v!=0 {
            return false
        }
    }
    return true
}











