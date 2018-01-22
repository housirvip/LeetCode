package leetcode

import "testing"

func Test_KMP(t *testing.T) {
	model := "aba"
	t.Log(getNext(model))
	t.Log(KMP(model, "abdabac"))
}
