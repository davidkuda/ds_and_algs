package main

 type ListNode struct {
    Val int
    Next *ListNode
 }

 // --- reverse linked list

 // reverse by creating a new Node
 func reverseList1(head *ListNode) *ListNode {
    var n *ListNode
    for head != nil {
        n = &ListNode{Val: head.Val, Next: n}
        head = head.Next
    }
    return n
}

// reverse by storing head as tempNext
func reverseList2(head *ListNode) *ListNode {
    var r *ListNode
    for head != nil {
        tempNext := head.Next
        head.Next = r
        r = head
        head = tempNext
    }
    return r
}
