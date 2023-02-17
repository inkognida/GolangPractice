package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func len_(l *ListNode) string {
	tmp := l
	var s string

	for tmp != nil {
		s += strconv.Itoa(tmp.Val)
		tmp = tmp.Next
	}

	return s
}

func addNode(l **ListNode, value int) {
	tmp := ListNode{
		Val:  value,
		Next: nil,
	}

	iter := l
	for (*iter).Next != nil {
		*iter = (*iter).Next
	}
	(*iter).Next = &tmp
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1 := len_(l1)
	s2 := len_(l2)

	v1, _ := strconv.ParseInt(s1, 10, 64)
	v2, _ := strconv.ParseInt(s2, 10, 64)

	r := int(v1 + v2)

	var rl *ListNode
	for r > 0 {
		addNode(&rl, r%10)
		r /= 10
	}

	return rl
}

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	r := addTwoNumbers(l1, l2)
	fmt.Println(r)
}
