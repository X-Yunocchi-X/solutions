/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	resultHead := &ListNode{}
	resultTail := resultHead

	// sorted head, sorted tail
	reverse := func(head *ListNode) (*ListNode, *ListNode) {
		if head == nil || head.Next == nil {
			return head, head
		}
		prev := (*ListNode)(nil)
		for curr := head; curr != nil; {
			next := curr.Next
			curr.Next = prev
			prev, curr = curr, next
		}
		return prev, head
	}

	curr := head
	for curr != nil {
		count := 0
		tempHead := curr
		prev := (*ListNode)(nil)
		for curr != nil && count != k {
			prev = curr
			curr = curr.Next
			count++
		}
		if count != k {
			resultTail.Next = tempHead
			break
		} else {
			prev.Next = nil
			sortedHead, sortedTail := reverse(tempHead)
			resultTail.Next = sortedHead
			resultTail = sortedTail
		}
	}
	return resultHead.Next
}
