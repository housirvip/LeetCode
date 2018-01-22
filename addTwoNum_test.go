package main

import (
	"fmt"
	"testing"
)

func Test_addTwoNum(t *testing.T) {
	l1 := &ListNode{5, nil}
	l2 := &ListNode{5, nil}
	res := addTwoNumbers(l1, l2)
	fmt.Println(res)
	fmt.Println(res.Next)
}
