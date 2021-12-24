package mian

//https://leetcode-cn.com/problems/add-two-numbers/
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res *ListNode = &ListNode{}
	var root *ListNode = res
	var inc int
	for l1 != nil || l2 != nil || inc > 0 {
		var l1V, l2V int
		if l1 != nil {
			l1V = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			l2V = l2.Val
			l2 = l2.Next
		}

		num := l1V + l2V + inc
		node := &ListNode{Val: num % 10}
		inc = num / 10
		root.Next = node
		root = node
	}

	return res.Next
}
