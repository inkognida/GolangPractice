package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func ListLen(head *ListNode) int {
	l := 0
	for head != nil {
		l++
		head = head.Next
	}

	return l
}

func printNodeLists(nodes []*ListNode) {
	for i := 0; i < len(nodes); i++ {
		fmt.Printf("List %d: ", i)
		curr := nodes[i]
		for curr != nil {
			fmt.Printf("%d ", curr.Val)
			curr = curr.Next
		}
		fmt.Println()
	}
}

func splitList(len int, k int, head *ListNode) []*ListNode {
	var r []*ListNode

	if len < k {
		r = make([]*ListNode, k)
		for i, _ := range r {
			if head != nil {
				r[i] = &ListNode{
					Val:  head.Val,
					Next: nil,
				}
				head = head.Next
			} else {
				r[i] = &ListNode{}
			}
		}
	} else if len%k == 0 {
		r = make([]*ListNode, k)
		fmt.Println(r, len/k)
		for i, _ := range r {
			for j := 0; j < len/k; j++ {
				r[i] = &ListNode{
					Val:  head.Val,
					Next: head.Next,
				}
				head = head.Next

			}

		}
		printNodeLists(r)
	}
	//else {
	//	r = make([]*ListNode, k)
	//
	//	subSum := 0
	//	for i := 0; i < k; i++ {
	//		r[i] = make([]int, len/k)
	//		subSum += len / k
	//	}
	//	for i := 0; i < k; i++ {
	//		if subSum == len {
	//			break
	//		}
	//		r[i] = append(r[i], 0)
	//		subSum++
	//	}
	//
	//}
	//printNodeLists(r)
	return r
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	_ = ListLen(head)
	splitList(9, 3, head)

	return nil
}

func initListNode(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	curr := head
	for i := 1; i < len(vals); i++ {
		curr.Next = &ListNode{Val: vals[i]}
		curr = curr.Next
	}
	return head
}

func printListNode(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Printf("%d ", curr.Val)
		curr = curr.Next
	}
	fmt.Println()
}

func main() {
	k := 3
	list := initListNode([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})

	splitListToParts(list, k)
}
