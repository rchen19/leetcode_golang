 /*
Given an integer n, generate all structurally unique BST's (binary search trees) that store values 1 ... n.

Example:

Input: 3
Output:
[
  [1,null,3,2],
  [3,2,null,1],
  [3,1,null,null,2],
  [2,1,3],
  [1,null,2,null,3]
]
Explanation:
The above output corresponds to the 5 unique BST's shown below:

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3

*/

package main
import("fmt")

func main() {
    var n = 3
    fmt.Println(generateTrees(n))
}

//Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

//dp - improved2
func generateTrees(n int) []*TreeNode {

    if n==0 {
        //need to return an empty slice {}
        //not an slice of empty slice {{},}
        //all other cases the return value is a slice of slice {{},}
        return []*TreeNode {}
    } else {
        //L[i][j] is the list of root nodes for all
        //possible trees with elements from i to j
        //rows:0-->n+1 
        //cols:0-->n
        var L = make([][][]*TreeNode, n+2)
        for i:=0;i<n+2;i++ {
            L[i] = make([][]*TreeNode, n+2)
        }
        //if start>end: nil
        //if start==end: one single root for a single tree
        for i:=1;i<n+2;i++ {
            L[i][i-1] = []*TreeNode {nil}
            L[i][i] = []*TreeNode {&TreeNode{Val:i}}
        }

        //L[n+1][n] = []*TreeNode {nil}
        //fmt.Println(len(L), len(L[0]), len(L[0][0]))

        //j: end index
        for j:=2;j<n+1;j++ {
            //i: starting index
            for i:=j-1;i>0;i-- {
                //possible root node at k within range of i->j
                for k:=i;k<j+1;k++ {
                    //fmt.Println(j, k, l)
                    //var left_nodes = L[i][k-1]
                    //var right_nodes = L[k+1][j]

                    for _, left_root := range L[i][k-1] {
                        for _, right_root := range L[k+1][j] {
                            //var root = &TreeNode{Val:k, Left:left_root, Right:right_root}

                            L[i][j] = append(L[i][j], &TreeNode{Val:k, Left:left_root, Right:right_root})
                        }
                    }
                }
            }
        }
        return L[1][n]
    }
}


//dp - improved
func generateTrees_dp2(n int) []*TreeNode {

    if n==0 {
        //need to return an empty slice {}
        //not an slice of empty slice {{},}
        //all other cases the return value is a slice of slice {{},}
        return []*TreeNode {}
    } else {
        //L[i][j] is the list of root nodes for all
        //possible trees with elements from i to j
        //rows:0-->n+1 
        //cols:0-->n
        var L = make([][][]*TreeNode, n+2)
        for i:=0;i<n+2;i++ {
            L[i] = make([][]*TreeNode, n+1)
        }
        //if start>end: nil
        //if start==end: one single root for a single tree
        for i:=1;i<n+1;i++ {
            L[i][i-1] = []*TreeNode {nil}
            L[i][i] = []*TreeNode {&TreeNode{Val:i}}
        }
        //the last row
        for i:=0;i<n+1;i++ {
            //fmt.Println(n+1, i, len(L), len(L[0]))
            //fmt.Println(L[n+1])
            L[n+1][i] = []*TreeNode {nil}
        }
        //fmt.Println(len(L), len(L[0]), len(L[0][0]))

        //j: end index
        for j:=2;j<n+1;j++ {
            //i: starting index
            for i:=j-1;i>0;i-- {
                //possible root node at k within range of i->j
                for k:=i;k<j+1;k++ {
                    //fmt.Println(j, k, l)
                    var left_nodes = L[i][k-1]
                    var right_nodes = L[k+1][j]

                    for _, left_root := range left_nodes {
                        for _, right_root := range right_nodes {
                            var root = &TreeNode{Val:k, Left:left_root, Right:right_root}

                            L[i][j] = append(L[i][j], root)
                        }
                    }
                }
            }
        }
        return L[1][n]
    }
}


//dp 
func generateTrees_dp1(n int) []*TreeNode {

    if n==0 {
        //need to return an empty slice {}
        //not an slice of empty slice {{},}
        //all other cases the return value is a slice of slice {{},}
        return []*TreeNode {}
    } else {
        //L[i][j] is the list of root nodes for all
        //possible trees with elements from i to j
        //rows:0-->n+1 
        //cols:0-->n
        var L = make([][][]*TreeNode, n+2)
        for i:=0;i<n+2;i++ {
            L[i] = make([][]*TreeNode, n+1)
        }
        //if start>end: nil
        //if start==end: one single root for a single tree
        for i:=1;i<n+1;i++ {
            L[i][i-1] = []*TreeNode {nil}
            L[i][i] = []*TreeNode {&TreeNode{Val:i}}
        }
        for i:=0;i<n+1;i++ {
            //fmt.Println(n+1, i, len(L), len(L[0]))
            //fmt.Println(L[n+1])
            L[n+1][i] = []*TreeNode {nil}
        }
        //fmt.Println(len(L), len(L[0]), len(L[0][0]))

        //l: lenghth of l to look at
        for l:=2;l<n+1;l++ {
            //starting from index i
            for i:=1;i<n+2-l;i++ {
                //possible root node at k within range of i->i+l-1
                for k:=i;k<i+l;k++ {
                    //fmt.Println(j, k, l)
                    var left_nodes = L[i][k-1]
                    var right_nodes = L[k+1][i+l-1]

                    for _, left_root := range left_nodes {
                        for _, right_root := range right_nodes {
                            var root = &TreeNode{Val:k, Left:left_root, Right:right_root}

                            L[i][i+l-1] = append(L[i][i+l-1], root)
                        }
                    }
                }
            }
        }
        return L[1][n]
    }
}


//recursion 
func generateTrees0(n int) []*TreeNode {

    if n==0 {
        return []*TreeNode {}
    } else if n==1 {
        var root = &TreeNode{Val:1}
        return []*TreeNode {root}
    } else {
        return genSub(1, n)
    }
}

func genSub(left int, right int) []*TreeNode {
    var output []*TreeNode
    if left>right {
        return append(output, nil)
    } else if left==right {
        var newNode = &TreeNode{Val:left}
        return append(output, newNode)
    } else {
        for i:=left;i<=right;i++ {
            //select nums[i] to be root
            //generate all possible root node for left and right subtrees
            var left_tree []*TreeNode = genSub(left, i-1)
            var right_tree []*TreeNode = genSub(i+1, right)
            //iterate through all the possible root nodes in the left and right subtrees
            for j:=0;j<len(left_tree);j++ {
                for k:=0;k<len(right_tree);k++ {
                    //for each combination of left and right subtree root node
                    //create a root to combine the two
                    var root = &TreeNode{Val:i, Left:left_tree[j], Right:right_tree[k]}
                    output = append(output, root)
                }
            }

        }

        return output
    }
}

func generateTrees1(n int) []*TreeNode {
    //var output []*TreeNode

    if n==0 {
        return []*TreeNode {}
    } else if n==1 {
        var root = &TreeNode{Val:1}
        //root.Val = 1
        //root.Right = nil
        //root.Left = nil
        return []*TreeNode {root}
    } else {
        var nums []int = makeRange(1,n)
        return genSub1(nums, 0, n-1)
    }
}

func genSub1(nums []int, left int, right int) []*TreeNode {
    var output []*TreeNode
    if left>right {
        return append(output, nil)
    } else if left==right {
        var newNode = &TreeNode{Val:nums[left]}
        //newNode.Val = nums[left]
        //newNode.Right = nil
        //newNode.Left = nil
        return append(output, newNode)
    } else {
        for i:=left;i<=right;i++ {
            //select nums[i] to be root
            //generate all possible root node for left and right subtrees
            var left_tree []*TreeNode = genSub1(nums, left, i-1)
            var right_tree []*TreeNode = genSub1(nums, i+1, right)
            //iterate through all the possible root nodes in the left and right subtrees
            for j:=0;j<len(left_tree);j++ {
                for k:=0;k<len(right_tree);k++ {
                    //for each combination of left and right subtree root node
                    //create a root to combine the two
                    var root = &TreeNode{Val:nums[i]}
                    //root.Val = nums[i]
                    root.Left = left_tree[j]
                    root.Right = right_tree[k]
                    output = append(output, root)
                }
            }

        }

        return output
    }
}

func makeRange(start, end int) []int {
    a := make([]int, end-start+1)
    for i := range a {
        a[i] = start + i
    }
    return a
}




