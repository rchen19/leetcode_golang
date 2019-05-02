/*
Given a 2D board and a word, find if the word exists in the grid.

The word can be constructed from letters of sequentially adjacent cell, 
where "adjacent" cells are those horizontally or vertically neighboring. 
The same letter cell may not be used more than once.

Example:

board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

Given word = "ABCCED", return true.
Given word = "SEE", return true.
Given word = "ABCB", return false.
*/

package main

import("fmt")


func main() {

    var board = [][]byte {{'A','B','C','E'},
                          {'S','F','C','S'},
                          {'A','D','E','E'}}
    var word = "ABCCED" 

    var result = exist(board, word)
    fmt.Println(result)

}

type Key struct {
    X, Y int
}

func exist(board [][]byte, word string) bool {
    if len(board) == 0 {
        return false
    }
    if len(board[0]) == 0 {
        return false
    }
    var h, w = len(board), len(board[0])
    var checked = make(map[Key]bool)
    var find = false
    for i:=0; i<h; i++ {
        for j:=0; j<w; j++ {
            if board[i][j] == word[0] {
                checked[Key{i,j}] = true

                find = search(board, [2]int{i,j}, word, 1, checked)
                delete(checked, Key{i,j})
                if find {
                    break
                }

            }
            
        }
        if find {
            break
        }
    }
    return find
    
}

func search(board [][]byte, curr_pos [2]int, word string, word_pos int, checked map[Key]bool) bool {
    var moves = [4][2]int{{0,1}, {1,0} ,{-1,0}, {0,-1}}
    var output = false
    if word_pos == len(word) {
        output = true
    } else {
        var dead_end_cout = 0
        for _, move := range moves {
            var next_pos = [2]int{curr_pos[0]+move[0], curr_pos[1]+move[1]}
            var next_pos_key = Key{curr_pos[0]+move[0], curr_pos[1]+move[1]}
            var _, ok = checked[next_pos_key]
            if ! ok && next_pos[0]>=0 && next_pos[0]<len(board) && 
                next_pos[1]>=0 && next_pos[1]<len(board[0]) && 
                board[next_pos[0]][next_pos[1]]==word[word_pos] {
                checked[next_pos_key] = true
                output = search(board, next_pos, word, word_pos+1, checked)
                delete(checked, next_pos_key)
            } else {
                dead_end_cout += 1
            }
            if output {
                break
            }

        }
        if dead_end_cout==4 {
            output = false
        }
    }

    return output

}











