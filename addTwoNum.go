package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var resTemp, resNode, resHead *ListNode
	headNotGot := true
	overflow := 0
	value := 0
	for overflow != 0 || l1 != nil || l2 != nil {
		value = overflow
		if l1 != nil {
			value += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			value += l2.Val
			l2 = l2.Next
		}
		overflow = value / 10
		resTemp = &ListNode{Val: value % 10, Next: nil}
		if headNotGot {
			resHead = resTemp
			resNode = resHead
			headNotGot = false
		} else {
			resNode.Next = resTemp
			resNode = resNode.Next
		}
	}
	return resHead
}
