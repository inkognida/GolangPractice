package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = prev
		prev = head
		head = tmp
	}
	head = prev

	return head
}

func insert(head *ListNode, val int) {
	for head.Next != nil {
		head = head.Next
	}

	head.Next = &ListNode{
		Val:  val,
		Next: nil,
	}
}

func main() {
	l := &ListNode{Val: 1, Next: nil} // start
	for i := 2; i <= 5; i++ {
		insert(l, i)
	}

	r := reverseList(l)
	for r != nil {
		fmt.Println(r.Val)
		r = r.Next
	}
}
