/*
Given a linked list and a value x, partition it such that 
all nodes less than x come before nodes greater than or equal to x.

You should preserve the original relative order of the nodes 
in each of the two partitions.

Example:

Input: head = 1->4->3->2->5->2, x = 3
Output: 1->2->2->4->3->5
*/


// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}
 

func partition(head *ListNode, x int) *ListNode {
    if head == nil {
        return nil
    }
    
    if head.Next == nil {
        return head
    }
    
    var large_head *ListNode
    var small_head *ListNode
    var small_end *ListNode
    var large_end *ListNode
    var large_count = 0
    var small_count = 0
    for head != nil {
        if head.Val < x {
            if small_count==0 {
                small_head = head
                small_end = head
                small_count++
            } else {
                small_end.Next = head
                small_end = small_end.Next
                small_count++
            }
        } else {
            if large_count==0 {
                large_head = head
                large_end = head
                large_count++
            } else {
                large_end.Next = head
                large_end = large_end.Next
                large_count ++
            }
        }
        head = head.Next
    }
    if large_end != nil{
        large_end.Next = nil
    }
    if small_end == nil {
        return large_head
    } else {
        small_end.Next = large_head

        return small_head
    }
}





