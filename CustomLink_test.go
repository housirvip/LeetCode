package leetcode

import "testing"

func Test_createLinkList(t *testing.T) {
	nums:=[]int{1,3,5,8,5}
	list:=createLinkList(nums)
	t.Log(list)
	t.Log(showLinkList(list))
}
