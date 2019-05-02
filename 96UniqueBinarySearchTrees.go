 /*
Given n, how many structurally unique BST's 
(binary search trees) that store values 1 ... n?

Example:

Input: 3
Output: 5
Explanation:
Given n = 3, there are a total of 5 unique BST's:

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3

*/

package main
import("fmt")

func main() {
    var n int = 3
    fmt.Println(numTrees(n))
}

//dp
func numTrees(n int) int {

    if n == 0 || n == 1 {
        return 1
    } else {
        var nums = make([]int, n+1)
        nums[0] = 1
        nums[1] = 1

        for i:=2;i<=n;i++ {
            for j:=1;j<=i;j++ {
                //fmt.Println(i,j)
                nums[i] += (nums[j-1] * nums[i-j])
            }
        }
        //fmt.Println(nums)
        return nums[n]
    }
    
}

func relu1(a int) int {
    if a <= 1 {
        return 1
    } else {
        return a
    }
}
 
func numTrees1(n int) int {

    if n == 0 {
        return 0
    } else if n == 1 {
        return 1
    } else {
        var num int = 0
        for i:=1;i<=n;i++ {
            var numLeft = numTrees(i-1)
            var numRight = numTrees(n-i)
            num += (relu1(numLeft) * relu1(numRight))
        }

        return num
    }
    
}

//recursion
func numTrees2(n int) int {

    if n == 0 || n == 1 {
        return 1
    } else {
        var num int = 0
        for i:=1;i<=n;i++ {

            num += numTrees(i-1) * numTrees(n-i)
        }

        return num
    }
    
}


