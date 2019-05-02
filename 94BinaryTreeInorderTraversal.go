/*
Given a binary tree, return the inorder traversal of its nodes' values.

Example:

Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [1,3,2]
Follow up: Recursive solution is trivial, could you do it iteratively?
*/
package main
import("fmt")

func main() {
    fmt.Println("hello world")
}


//Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}


//tree traversal using iteration instead of recursion
//this method does not modify the original tree
func inorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    var result []int
    var toExplore []*TreeNode
    var curr_node *TreeNode = root

    for true {

        //add left node to toExplore
        for curr_node!=nil {
            toExplore = append(toExplore, curr_node)
            curr_node = curr_node.Left
        }

        if len(toExplore)==0 {
            break
        }

        curr_node = toExplore[len(toExplore)-1]
        toExplore = toExplore[:len(toExplore)-1]
        result = append(result, curr_node.Val)
        curr_node = curr_node.Right

    }
    return result
}

//tree traversal using iteration instead of recursion
//this method modifies the original tree
func inorderTraversal3(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    var result []int
    var toExplore []*TreeNode = []*TreeNode {root}
    var curr_node *TreeNode

    for len(toExplore)>0 {
        //start exploring the last element in toExplore
        curr_node = toExplore[len(toExplore)-1]
        //first look at left child
        if curr_node.Left == nil {
            //current node value needs to be add to result
            result = append(result, curr_node.Val)
            //remove it from toExplore
            toExplore = toExplore[:len(toExplore)-1]
            // then start look at the right child
            if curr_node.Right != nil {
                toExplore = append(toExplore, curr_node.Right)
            } 

        } else {
            //add the left child to the toExplore list
            toExplore = append(toExplore, curr_node.Left)
            //below modifies the original tree
            //just to mark the left child as already explored
            curr_node.Left = nil
        }
    }
    return result
}