package leetcode

import (
	"testing"
)

func Test_addTwoNum(t *testing.T) {
	l1 := createLinkList([]int{1,3,5,8,5})
	l2 := createLinkList([]int{1,3,5,8})
	res := addTwoNumbers(l1, l2)
	t.Log(showLinkList(res))
}
