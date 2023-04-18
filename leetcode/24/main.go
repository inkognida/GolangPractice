package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swap(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := head.Next
	currNode := head
	var prev *ListNode

	for currNode != nil && currNode.Next != nil {
		tmp := currNode.Next

		currNode.Next = tmp.Next
		tmp.Next = currNode

		if prev != nil {
			prev.Next = tmp
		}

		prev = currNode
		currNode = currNode.Next
	}
	return newHead
}

func main() {
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	//l2 := &ListNode{
	//	Val:  1,
	//	Next: nil,
	//}

	r1 := swapPairs(l)
	for r1 != nil {
		fmt.Println(r1.Val)
		r1 = r1.Next
	}

	//fmt.Println(swapPairs(l))
	//fmt.Println(swapPairs(l2))
	//fmt.Println(swapPairs(l3))
}
