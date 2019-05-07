/*
Given an array of words and a width maxWidth, 
format the text such that each line has exactly 
maxWidth characters and is fully (left and right) justified.

You should pack your words in a greedy approach; 
that is, pack as many words as you can in each line. 
Pad extra spaces ' ' when necessary so that each line has exactly maxWidth characters.

Extra spaces between words should be distributed 
as evenly as possible. If the number of spaces on a 
line do not divide evenly between words, the empty slots 
on the left will be assigned more spaces than the slots on the right.

For the last line of text, it should be left justified 
and no extra space is inserted between words.

Note:

A word is defined as a character sequence consisting of non-space characters only.
Each word's length is guaranteed to be greater than 0 and not exceed maxWidth.
The input array words contains at least one word.
Example 1:

Input:
words = ["This", "is", "an", "example", "of", "text", "justification."]
maxWidth = 16
Output:
[
   "This    is    an",
   "example  of text",
   "justification.  "
]
Example 2:

Input:
words = ["What","must","be","acknowledgment","shall","be"]
maxWidth = 16
Output:
[
  "What   must   be",
  "acknowledgment  ",
  "shall be        "
]
Explanation: Note that the last line is "shall be    " instead of "shall     be",
             because the last line must be left-justified instead of fully-justified.
             Note that the second line is also left-justified becase it contains only one word.
Example 3:

Input:
words = ["Science","is","what","we","understand","well","enough","to","explain",
         "to","a","computer.","Art","is","everything","else","we","do"]
maxWidth = 20
Output:
[
  "Science  is  what we",
  "understand      well",
  "enough to explain to",
  "a  computer.  Art is",
  "everything  else  we",
  "do                  "
]
*/

package main
import("fmt"; "strings")

func main() {
    var words []string = []string {"Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we","do"}
    var maxWidth int = 20
    var result []string = fullJustify(words, maxWidth) 
    fmt.Println("[")
    for i, s := range result {
        if i==len(result)-1 {
            fmt.Println("\"" + s + "\"")
        } else {
            fmt.Println("\"" + s + "\",")
        }
        
    }
    fmt.Println("]")
}


func fullJustify(words []string, maxWidth int) []string {
    var out []string
    var i int = 0
    var j int
    var numChar int
    var numSpacing int
    var totalSpaces int
    var line string
    var numSpacesPerSpacing int
    var residue int

    for i<len(words) {
        j = i + 1
        numChar = len(words[i])
        for j<len(words) {
            if numChar+len(words[j])+j-i<=maxWidth {
                numChar += len(words[j])
                j++
            } else {
                break
            }
        }
        
        numSpacing = j - i - 1
        totalSpaces = maxWidth - numChar
        line = ""

        //last line
        if j == len(words) {
            for k:=i;k<j;k++ {
                if k==j-1 {
                    line += words[k] + strings.Repeat(" ", totalSpaces - numSpacing)
                } else {
                    line += words[k] + " "
                }
            }
        } else if numSpacing>0 {
            numSpacesPerSpacing = totalSpaces / numSpacing
            residue = totalSpaces % numSpacing
            for k:=i;k<j;k++ {
                if k-i+1<=residue {
                    line += words[k] + strings.Repeat(" ", numSpacesPerSpacing+1)
                } else if k==j-1 {
                    line += words[k]
                } else {
                    line += words[k] + strings.Repeat(" ", numSpacesPerSpacing)
                }
            }
        } else {
            line += words[i] + strings.Repeat(" ", totalSpaces)
        }

        out = append(out, line)
        i = j
    }
    return out
}






























