/*
Given a sorted array nums, remove the duplicates in-place such that duplicates appeared at most twice and return the new length.

Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

Example 1:

Given nums = [1,1,1,2,2,3],

Your function should return length = 5, with the first five elements of nums being 1, 1, 2, 2 and 3 respectively.

It doesn't matter what you leave beyond the returned length.
Example 2:

Given nums = [0,0,1,1,1,1,2,3,3],

Your function should return length = 7, with the first seven elements of nums being modified to 0, 0, 1, 1, 2, 3 and 3 respectively.

It doesn't matter what values are set beyond the returned length.
Clarification:

Confused why the returned value is an integer but your answer is an array?

Note that the input array is passed in by reference, which means modification to the input array will be known to the caller as well.

Internally you can think of this:

// nums is passed in by reference. (i.e., without making a copy)
int len = removeDuplicates(nums);

// any modification to nums in your function would be known by the caller.
// using the length returned by your function, it prints the first len elements.
for (int i = 0; i < len; i++) {
    print(nums[i]);
}
*/

package main
import("fmt")

func removeDuplicates1(nums []int) int {
    var n, i, j, count = len(nums), 0, 1, 1
    for j<n {
        if nums[i]==nums[j] {
            count += 1
            if count == 3 {
                count -= 1
                j += 1
            } else {
                nums[i+1] = nums[j]
                i += 1
                j += 1
            }
        } else {
            fmt.Println(i ,j)
            nums[i+1] = nums[j]
            count = 1
            i += 1
            j += 1
        }

    }
    fmt.Println(nums)
    return i+1
}

func removeDuplicates(nums []int) int {
    var i int = 0
    for _, n := range nums {
        if i < 2 || n > nums[i-2] {
            //nums[i] = n
            i++
        }
    }
    return i
}


func main() {

    var nums = []int {0,0,1,1,1,1,2,3,3}
 
    var result = removeDuplicates(nums)
    fmt.Println(result)

}