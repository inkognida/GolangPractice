package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

// TIMEOUT SOLUTION
/*
func maxElementIndex(slice []int) int {
	indexMax := 0
	max := 0

	for i, v := range slice {
		if v > max {
			max = v
			indexMax = i
		}
	}

	return indexMax
}

func minStoneSum(piles []int, k int) int {
	sum := 0

	if len(piles) == 1 {
		for ; k > 0; k-- {
			piles[0] = int(math.Ceil(float64(piles[0]) / float64(2)))
		}
		return piles[0]
	}

	for ; k > 0; k-- {
		index := maxElementIndex(piles)
		piles[index] = int(math.Ceil(float64(piles[index]) / float64(2)))
	}

	for _, v := range piles {
		sum += v
	}

	return sum
}

func main() {
	result := minStoneSum([]int{4122, 9928, 3477, 9942}, 6)

	fmt.Println(result)
}
*/

//type IntHeap []int
//
//func (h IntHeap) Len() int {
//	return len(h)
//}
//
//func (h IntHeap) Less(i, j int) bool {
//	return h[i] < h[j]
//}
//
//func (h IntHeap) Swap(i, j int) {
//	h[i], h[j] = h[j], h[i]
//}
//
//func (h *IntHeap) Push(x any) {
//	*h = append(*h, x.(int))
//}
//
//func (h *IntHeap) Pop() any {
//	old := *h
//	n := len(old)
//	x := old[n-1]
//	*h = old[0 : n-1]
//	return x
//}
//
//func minStoneSum(piles []int, k int) int {
//	h := &IntHeap{}
//	heap.Init(h)
//	sum := 0
//
//	for _, v := range piles {
//		heap.Push(h, v)
//		sum += v
//	}
//
//	for i := 0; i < k; i++ {
//		tmp := heap.Pop(h).(int)
//		sum -= int(math.Ceil(float64(tmp) / float64(2)))
//		heap.Push(h, int(math.Ceil(float64(tmp)/float64(2))))
//	}
//
//	return sum
//}

type Node struct {
	Value int
	Count int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(value int) error {
	if n == nil {
		return errors.New("unable to insert")
	}

	switch {
	case value == n.Value:
		return nil

	case value < n.Value:
		if n.Left == nil {
			n.Left = &Node{Value: value, Count: 1}
			return nil
		}
		return n.Left.Insert(value)

	case value > n.Value:
		if n.Right == nil {
			n.Right = &Node{Value: value, Count: 1}
			return nil
		}
		return n.Right.Insert(value)
	}
	return nil
}

func (n *Node) Find(value int) (int, bool) {
	if n == nil {
		return 0, false
	}

	switch {
	case value == n.Value:
		return n.Value, true
	case value < n.Value:
		return n.Left.Find(value)
	default:
		return n.Right.Find(value)
	}
}

func (n *Node) findMax(parent *Node) (*Node, *Node) {
	if n == nil {
		return nil, parent
	}

	if n.Right == nil {
		return n, parent
	}

	return n.Right.findMax(n)
}

func (n *Node) replaceNode(parent, replacement *Node) error {
	if n == nil {
		return errors.New("unable to replaceNode")
	}

	if n == parent.Left {
		parent.Left = replacement
		return nil
	}

	parent.Right = replacement
	return nil
}

func (n *Node) Delete(value int, parent *Node) error {
	if n == nil {
		return errors.New("unable to delete")
	}

	switch {
	case value < n.Value:
		return n.Left.Delete(value, n)
	case value > n.Value:
		return n.Right.Delete(value, n)
	default:
		if n.Left == nil && n.Right == nil {
			err := n.replaceNode(parent, nil)
			if err != nil {
				log.Fatalln(err)
			}

			return nil
		}

		if n.Left == nil {
			err := n.replaceNode(parent, n.Right)
			if err != nil {
				log.Fatalln(err)
			}
		}

		if n.Right == nil {
			err := n.replaceNode(parent, n.Left)
			if err != nil {
				log.Fatalln(err)
			}
			return nil
		}

		replacement, replParent := n.Left.findMax(n)

		n.Value = replacement.Value

		return replacement.Delete(replacement.Value, replParent)
	}
}

type Tree struct {
	Root *Node
}

func (t *Tree) Insert(value int) error {
	if t.Root == nil {
		t.Root = &Node{Value: value, Count: 1}
		return nil
	}

	return t.Root.Insert(value)
}

func (t *Tree) Find(value int) (int, bool) {
	if t.Root == nil {
		return 0, false
	}

	return t.Root.Find(value)
}

func (t *Tree) findMax(parent *Node) (*Node, *Node) {
	if t.Root == nil {
		return nil, parent
	}

	if t.Root.Right == nil {
		return t.Root, parent
	}

	return t.Root.findMax(t.Root)
}

func (t *Tree) Delete(value int) error {
	if t.Root == nil {
		return errors.New("unable to delete from empty tree")
	}

	fakeParent := &Node{Right: t.Root}
	err := t.Root.Delete(value, fakeParent)
	if err != nil {
		return err
	}

	if fakeParent.Right == nil {
		t.Root = nil
	}
	return nil
}

func (t *Tree) Traverse(n *Node, f func(*Node)) {
	if n == nil {
		return
	}

	t.Traverse(n.Left, f)
	f(n)
	t.Traverse(n.Right, f)

}

//tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Value, ":", n.Count, " ") })
//fmt.Println()

func minStoneSum(piles []int, k int) int {
	tree := &Tree{}
	for _, v := range piles {
		err := tree.Insert(v)
		if err != nil {
			log.Fatalln(err)
		}
	}

	sum := 0
	for k > 0 {
		maxNode, parentNode := tree.findMax(tree.Root)

		delta := int(math.Ceil(float64(maxNode.Value) / float64(2)))

		if maxNode.Count == 1 {
			err := tree.Delete(maxNode.Value)
			if err != nil {
				log.Fatalln(err)
			}
		} else {
			maxNode.Count--
		}

		_, flag := tree.Find(delta)
		if flag != false {
			parentNode.Count += 1
		}

		err := tree.Insert(delta)
		if err != nil {
			log.Fatalln(err)
		}

		tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Value, ":", n.Count, " ") })
		fmt.Println()

		k--
	}

	tree.Traverse(tree.Root, func(n *Node) { sum += n.Value })

	return sum
}

//1962. Remove Stones to Minimize the Total

func main() {
	values := []int{4, 3, 6, 7}
	k := 3

	result := minStoneSum(values, k)
	fmt.Println("Sum", result)
}
