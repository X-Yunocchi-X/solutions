/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package leetcode

func sortList(head *ListNode) *ListNode {
	listLength := func(head *ListNode) int {
		result := 0
		for head != nil {
			result++
			head = head.Next
		}
		return result
	}

	merge := func(a *ListNode, b *ListNode) *ListNode {
		if a == nil && b == nil {
			return nil
		}
		if a == nil {
			return b
		}
		if b == nil {
			return a
		}
		aHead, bHead := a, b
		dummy := &ListNode{}
		curr := dummy
		for aHead != nil && bHead != nil {
			if aHead.Val < bHead.Val {
				curr.Next = aHead
				aHead = aHead.Next
			} else {
				curr.Next = bHead
				bHead = bHead.Next
			}
			curr = curr.Next
		}
		if aHead == nil {
			curr.Next = bHead
		}
		if bHead == nil {
			curr.Next = aHead
		}
		return dummy.Next
	}

	length := listLength(head)
	dummy := &ListNode{Next: head}

	for sep := 1; sep < length; sep <<= 1 {
		curr := dummy.Next
		prev := dummy
		for curr != nil {
			tempHead1 := curr
			for i := 1; i < sep && curr != nil; i++ {
				curr = curr.Next
			}
			if curr == nil {
				continue
			}
			tempHead2 := curr.Next
			curr.Next = nil
			curr = tempHead2
			for i := 1; i < sep && curr != nil; i++ {
				curr = curr.Next
			}
			next := (*ListNode)(nil)
			if curr != nil {
				next = curr.Next
				curr.Next = nil
			}
			newHead := merge(tempHead1, tempHead2)
			prev.Next = newHead
			for prev.Next != nil {
				prev = prev.Next
			}
			prev.Next = next
			curr = next
		}
	}
	return dummy.Next
}
