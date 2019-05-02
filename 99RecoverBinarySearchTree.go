/*
Two elements of a binary search tree (BST) are swapped by mistake.

Recover the tree without changing its structure.

Example 1:

Input: [1,3,null,null,2]

   1
  /
 3
  \
   2

Output: [3,1,null,null,2]

   3
  /
 1
  \
   2
Example 2:

Input: [3,1,4,null,null,2]

  3
 / \
1   4
   /
  2

Output: [2,1,4,null,null,3]

  2
 / \
1   4
   /
  3
Follow up:

A solution using O(n) space is pretty straight forward.
Could you devise a constant space solution?
*/

//Implementation for a binary tree node.
package main
import("fmt"; "strconv"; "strings")


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

//create a new struct of integer than can be nil
type NilInt struct {
    value int
    null bool
}

//func (n *NilInt) Value() interface {} {
//    if n.null {
//        return nil
//    } else {
//        return n.value
//    }
//}

func NewInt(x int) NilInt {
    return NilInt{value:x, null:false}
}

func NewNil() NilInt {
    return NilInt{value:0, null:true}
}

//convert list representation to linked nodes representation
func convertTree(nums []NilInt) []*TreeNode {
    var nodes []*TreeNode
    var pos int = 1
    
    for _, num := range nums {
        if num.null {
            nodes = append(nodes, nil)
        } else {
            nodes = append(nodes, &TreeNode{Val:num.value})
        }
    }

    for i, num := range nums {

        if !num.null && pos<len(nums) {
            nodes[i].Left = nodes[pos]
            pos++
            if pos<len(nums) {
                nodes[i].Right = nodes[pos]
                pos++
            }
        }
    }
    return nodes

}

func printNodes(nodes []*TreeNode) {

    var out string
    for _, node := range nodes {
        if node==nil {
            out += ", null"

        } else {
            if len(out)==0 {
                out += "[" + strconv.Itoa(node.Val)
            } else {
                out += ", " + strconv.Itoa(node.Val)
            }
        }
    }
    out += "]"

    fmt.Println(out)
}

func main() {
    var input string = "1,3,null,null,2"
    var inputSplits []string = strings.Split(input, ",")
    var nums []NilInt
    for _, char := range inputSplits {
        if char=="null" {
            nums = append(nums, NewNil())
        } else {
            value, _ := strconv.Atoi(char)
            nums = append(nums, NewInt(value))
        }
    }
    //var nums []NilInt = []NilInt {NewInt(1),NewInt(3),NewNil(),NewNil(),NewInt(2)}

    var nodes []*TreeNode = convertTree(nums)
    printNodes(nodes)
    recoverTree(nodes[0])
    printNodes(nodes)

}


func recoverTree(root *TreeNode)  {
    //DFS in-order traversal
    var toExplore []*TreeNode
    var currNode *TreeNode = root
    var prevNode *TreeNode = nil

    var errorPos []*TreeNode

    for true {
        for currNode!=nil {
            toExplore = append(toExplore, currNode)
            currNode = currNode.Left
        }

        if len(toExplore)==0 {
            break
        }

        currNode = toExplore[len(toExplore)-1]
        toExplore = toExplore[:len(toExplore)-1]
        if prevNode!=nil && currNode.Val<prevNode.Val {
            errorPos = append(errorPos, prevNode, currNode)
        }
        prevNode = currNode
        currNode = currNode.Right
    }

    if len(errorPos)==2 {
        errorPos[0].Val, errorPos[1].Val = errorPos[1].Val, errorPos[0].Val
    } else if len(errorPos)==4 {
        errorPos[0].Val, errorPos[3].Val = errorPos[3].Val, errorPos[0].Val
    }
    
}






