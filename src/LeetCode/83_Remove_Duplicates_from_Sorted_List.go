/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    
    prev := head
    curr := head.Next

    for curr != nil {
        if curr.Val == prev.Val {
            prev.Next = curr.Next
        } else {
            prev = curr
        }   
        curr = curr.Next
    }

    return head
}
