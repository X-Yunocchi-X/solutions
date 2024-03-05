/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package leetcode

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	merge := func(a *ListNode, b *ListNode) *ListNode {
		dummy := &ListNode{}
		curr := dummy
		for a != nil && b != nil {
			if a.Val < b.Val {
				curr.Next = a
				a = a.Next
			} else {
				curr.Next = b
				b = b.Next
			}
			curr = curr.Next
		}
		if a == nil {
			curr.Next = b
		}
		if b == nil {
			curr.Next = a
		}
		return dummy.Next
	}

	var mergeSort func([]*ListNode) *ListNode

	mergeSort = func(list []*ListNode) *ListNode {
		if len(list) == 0 {
			return nil
		}
		if len(list) == 1 {
			return list[0]
		}
		mid := len(list) / 2
		return merge(mergeSort(list[:mid]), mergeSort(list[mid:]))
	}

	return mergeSort(lists)
}
