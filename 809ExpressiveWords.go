/*
Sometimes people repeat letters to represent extra feeling, 
such as "hello" -> "heeellooo", "hi" -> "hiiii".  
In these strings like "heeellooo", we have groups of 
adjacent letters that are all the same:  "h", "eee", "ll", "ooo".

For some given string S, a query word is stretchy if 
it can be made to be equal to S by any number of applications 
of the following extension operation: choose a group consisting 
of characters c, and add some number of characters c to the group 
so that the size of the group is 3 or more.

For example, starting with "hello", we could do an extension 
on the group "o" to get "hellooo", but we cannot get "helloo" 
since the group "oo" has size less than 3.  Also, we could do 
another extension like "ll" -> "lllll" to get "helllllooo".  
If S = "helllllooo", then the query word "hello" would be 
stretchy because of these two extension operations: 
query = "hello" -> "hellooo" -> "helllllooo" = S.

Given a list of query words, return the number of words that are stretchy. 

 

Example:
Input: 
S = "heeellooo"
words = ["hello", "hi", "helo"]
Output: 1
Explanation: 
We can extend "e" and "o" in the word "hello" to get "heeellooo".
We can't extend "helo" to get "heeellooo" because the group "ll" is not size 3 or more.
 

Notes:

0 <= len(S) <= 100.
0 <= len(words) <= 100.
0 <= len(words[i]) <= 100.
S and all words in words consist only of lowercase letters
*/

package main
import("fmt")

func expressiveWords(S string, words []string) int {
    var count int = 0
    for _, word := range words {
        //fmt.Println(word, checkWord(S, word))
        count += checkWord(S, word)
    }
    return count
}

func checkWord1(S string, word string) int {
    if S==word {
        return 1
    }
    if len(S)<len(word) {
        return 0
    }
    var i, j int = 0, 0
    for i<len(S) {
        //fmt.Println(i,j)
        if j<len(word) && S[i]==word[j] {
            i++
            j++
        } else if i<len(S)-1 && S[i]==S[i+1] {
            i+=2
        } else if (i<=1 || j==0 || S[i]!=word[j-1] || S[i]!=S[i-1] || S[i]!=S[i-2]) {
            break
        } else {
            i++
        }

    }
    if i==len(S) && j==len(word) {
        return 1
    } else {
        return 0
    }

}

func checkWord(S string, word string) int {
    if S==word {
        return 1
    }
    if len(S)<len(word) {
        return 0
    }
    var i, j int = 0, 0
    for i<len(S) {
        if i==0 && j==0 {
            if S[i]!=word[j] {
                return 0
            } else {
                i++
                j++
            }
        } else if j==len(word) {
            if i>=2 && S[i]==word[j-1] && S[i-1]==word[j-1] && S[i-2]==word[j-1] {
                i++
            } else if i+1<len(S) && S[i]==word[j-1] && S[i+1]==word[j-1] {
                i += 2
            } else {

                return 0
            }

        } else if S[i]!=word[j] {
            //fmt.Println(i,j)
            //fmt.Println(S[i-1], S[i-1], S[i], S[i+1], word[j-1])
            if i+1<len(S) && S[i]==word[j-1] && S[i+1]==word[j-1] {
                i += 2
            } else if i>=2 && S[i]==word[j-1] && S[i-1]==word[j-1] && S[i-2]==word[j-1]{
                i++
            } else {
                //fmt.Println("here")
                return 0
            }
            
        } else {
            i++
            j++
        }

    }
    if i==len(S) && j==len(word) {
        return 1
    } else {
        return 0
    }
}

func main() {
    //var S string = "dddiiiinnssssssoooo"

    //var words []string = []string {"dinnssoo","ddinso","ddiinnso","ddiinnssoo","ddiinso","dinsoo","ddiinsso","dinssoo","dinso"}
    
    var S string = "vvvppppeeezzzzztttttkkkkkkugggggbbffffffywwwwwwbbbccccddddddkkkkksssppppddpzzzzzhhhhbbbbbmmmy"

    var words []string = []string {"vvpeezttkkuggbbfywwbbccddkkspdpzhbbmmyy","vvppeeztkkugbfywwbccddkksspdppzhhbmyy","vppezzttkkugbffyywbccddksspddpzhhbmy","vvppezztkugbffyywwbbccddkssppddpzzhhbbmmy","vvppezttkuggbfyywwbbcddkspdppzhhbmy","vppeezzttkkuugbfyywwbbccdkkssppdpzzhbbmy","vpeezztkkugbbffyywwbbccddkksppdpzzhhbbmmy","vppeeztkkuuggbffywbbccddkksppdppzhhbmyy","vpeeztkkuggbfyywbbccdksppdpzhbmy","vpeezztkkugbffywwbbccdkkssppddppzzhhbbmmy","vvpeztkkugbbfyywbcdkssppddpzzhhbbmyy","vpezztkugbbffyywwbcddksppddpzzhbbmy","vvpeezztkkugbbffywwbccdkkspddpzzhbmmyy","vvpeezzttkkuuggbbffyywbbccdkspdppzhhbmy","vvppeezztkkuggbbfywbcdkspdpzhhbmyy","vvppeezzttkkuugbffyywwbbccddkkspddpzzhbmyy","vppezztkuuggbffywwbcdksspdppzhhbmyy","vvppeezzttkuuggbffywbccddkksspddppzzhhbmmy","vvppezzttkuggbffywbbccdkspddppzzhhbmy","vvpezzttkuugbbfywwbccdkssppdpzhbbmmy","vvpeezzttkuugbbffyywbccdksppddppzhhbmyy","vpeezzttkkuggbbffywbccddksppddpzhhbbmy","vvpezttkuuggbffywwbbccddkspdppzhhbmmyy","vppeezzttkkuugbffywbccddksppddpzhhbmmyy","vvpezttkkuugbbfywbccdkspddppzzhbbmmy","vppezzttkkuugbbffywwbcddkssppddpzhhbmmy","vppezzttkugbfywbbcdksppddppzzhhbmyy","vppeeztkuggbbffywbbccdkkspddppzzhbbmmy","vvpeeztkuuggbbfywbcdkksspddppzhhbbmmyy","vpezttkkuuggbbffyywwbbcdksspddppzhhbmy","vpeezzttkkuuggbffywwbccdkksspddppzzhbmyy","vpezttkkuugbffyywbccdksspddppzhbbmmyy","vvppezztkugbbffyywbbccdkksppdppzhbmyy","vvpeezttkuggbbfyywwbbcddkksppdpzzhbbmyy","vvpeztkuuggbffyywbbccdkksspddppzzhbbmy","vppeezzttkugbbffyywbccddksppdppzzhbmmyy","vppeezttkkuugbbfywwbccddkksspdpzhhbmmy","vpeezzttkugbbffywbbccdkksspddppzhbbmyy","vpeezttkkuugbbfywbbccddksppddppzzhhbmmy","vpeezztkuuggbbffywwbbccddksspddpzzhhbbmmyy","vppeezttkkuggbbffyywwbccdksspdpzzhbmy","vpezzttkkuugbbfyywbbcdksspdppzzhbbmyy","vvppezttkkuggbbfyywbbccdkksspddpzhbbmyy","vvpezzttkuggbbffyywbbcdkksppdpzzhbmmyy","vvpeztkugbfywwbccddkkspddpzhhbbmyy","vvppezttkuugbbfyywwbcddkksspdppzhhbbmy","vvpeeztkkuuggbbfywwbcdkspddpzzhhbmmy","vvpeezttkugbffywbbccdkkssppddppzhhbbmyy","vpeztkuuggbbfyywwbcddksppddpzhbbmy","vppeztkuggbbfyywbcdksspdppzzhhbmy","vppeezttkkugbbffyywbccddkksppdpzhhbmy","vvppeeztkugbfyywbcdkksppdppzhbmyy","vpezttkuugbbffywbcdksppddpzzhhbbmmy","vppezzttkuugbfyywbcddkksspdpzhbbmmy","vppezzttkkuggbffywbbcdksspdpzzhhbbmmyy","vpezzttkuggbfyywbbccdksspdpzhhbbmmy","vvppezttkkugbffyywbcdkssppdpzzhbmy","vvpeezttkkuuggbbfyywbbccdkspdppzhhbmy","vpeezttkkuugbfywbccddkksppddpzzhhbmmy","vvppezttkuuggbbffywbbccdkksppdpzzhhbbmmy","vvppeeztkuggbbffyywbccdksspddppzzhbmmyy","vvppeezztkuggbfywwbccddkkspddpzhbbmy","vpezttkuuggbfyywwbcdkkspdpzhhbbmmyy","vppezzttkuggbffywbbcdkkssppddppzhhbmyy","vppeztkuuggbffyywbccdkkspdppzzhhbmmyy","vppeezztkuuggbfywbccddkksspddppzhhbbmyy","vvppeztkuugbfywwbccdkkspddppzzhhbmmy","vvpezztkuugbbffyywwbbccddksppdpzhbbmmyy","vvpezzttkkuuggbffyywwbbcdkspdpzhbmmyy","vvppeztkkuuggbbfyywbbccdksppdppzzhbmmyy","vvppezztkuggbffyywwbcddkkssppdpzhbmmyy","vvpezzttkkuggbbffywwbcddkksspdpzzhhbbmmy","vpezztkkuuggbfyywwbccddkssppdppzhhbbmmy","vvppezztkuugbffywwbccdkkspdppzhhbmmy","vpeztkugbfyywwbcdkksspdppzzhbmmy","vvpeezzttkkugbbfywwbcdkkspdpzzhhbmmy","vpezzttkuuggbbfywbccdkspddppzzhhbbmmy","vppeztkkuugbffyywwbbcddksspddppzhbbmyy","vpeztkkuggbffyywbbccddkssppdppzhbmyy","vvppeezztkuggbffyywwbcddkksppdppzhbmyy","vpeezztkugbfyywbbccdkkspdppzhbmmyy","vvppeezttkugbfywwbcddkkssppdppzhbmmyy","vpeeztkuggbffywwbbccddksspdppzzhhbmmy","vvppeeztkuugbfywbcddkssppddppzzhhbbmyy","vpezzttkuggbbffyywwbbccdkssppddppzhbbmy","vpezttkugbfywbbcddkksspddppzhbbmy","vpeezzttkkuggbbffyywwbccddkspddppzhbmyy","vppeezzttkugbffywbccdkkspddpzhhbbmyy","vpezzttkuggbbfyywbbccdkksspddpzzhhbmmy","vvppezttkugbfywwbbcdkksspddpzzhhbbmyy","vppezztkkuggbffyywbcddkkssppddpzzhhbbmmy","vppeztkkuggbfywwbccdkksppdppzhhbmmy","vvpeezzttkugbffyywwbbcddkssppddpzzhbmmy","vvpezztkkuuggbfyywbccdkksspddpzhhbbmyy","vpezttkuuggbffywbbccdksppdpzhbmmyy","vvpezzttkuggbbfywbccddksspdpzzhhbmmy","vvpeezzttkkugbbfywwbcdkksppddpzhbmy","vppeezttkkuugbbfyywwbcddkkspdpzhhbbmmyy","vvppeeztkkuugbbfyywwbbcddkksspdppzhbbmyy","vvpeezzttkkuugbfywwbbcddkspdpzzhbbmyy"}
    fmt.Println(expressiveWords(S, words))
}




