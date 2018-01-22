package leetcode

func createLinkList(nums []int) *Node {
	var head, node, temp *Node
	headNotGot := true
	for i := 0; i < len(nums); i++ {
		temp = &Node{Val: nums[i]}
		if headNotGot {
			node = temp
			head = node
			headNotGot = false
		} else {
			node.Next = temp
			node = node.Next
		}
	}
	return head
}

func showLinkList(head *Node) []int {
	node := head
	res := make([]int, 0)
	for i := 0; node != nil; i++ {
		res = append(res, node.Val)
		node = node.Next
	}
	return res
}
