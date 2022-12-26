package main

import (
	"errors"
	"fmt"
	"log"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(value int) error {
	if n == nil {
		return errors.New("unable to insert")
	}

	switch {
	case n.Left != nil && n.Left.Value < value:
		if n.Right == nil {
			n.Right = &Node{Value: value}
			return nil
		}
		return n.Right.Insert(value)

	case n.Left != nil && n.Left.Value == value:
		return n.Left.Insert(value)

	case value <= n.Value:
		if n.Left == nil {
			n.Left = &Node{Value: value}
			return nil
		}
		return n.Left.Insert(value)

	case value > n.Value:
		if n.Right == nil {
			n.Right = &Node{Value: value}
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

func (n *Node) findMaxElement(parent *Node) (int, bool) {
	if n == nil {
		return 0, false
	}

	if n.Right == nil {
		return n.Value, true
	}

	return n.Right.findMaxElement(n)
}

func (n *Node) findMinElement(parent *Node) (int, bool) {
	if n == nil {
		return 0, false
	}

	if n.Left == nil {
		return n.Value, true
	}

	return n.Left.findMinElement(n)
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

type MinTree struct {
	Root *Node
}

func (t *MinTree) Insert(value int) error {
	if t.Root == nil {
		t.Root = &Node{Value: value}
		return nil
	}

	return t.Root.Insert(value)
}

func (t *MinTree) Find(value int) (int, bool) {
	if t.Root == nil {
		return 0, false
	}

	return t.Root.Find(value)
}

func (t *MinTree) findMax(parent *Node) (*Node, *Node) {
	if t.Root == nil {
		return nil, parent
	}

	if t.Root.Right == nil {
		return t.Root, parent
	}

	return t.Root.findMax(t.Root)
}

func (t *MinTree) Delete(value int) error {
	if t.Root == nil {
		return errors.New("unable to delete from empty MinTree")
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

func (t *MinTree) Traverse(n *Node, f func(*Node)) {
	if n == nil {
		return
	}

	t.Traverse(n.Left, f)
	f(n)
	t.Traverse(n.Right, f)

}

func longestObstacleCourseAtEachPosition(obstacles []int) []int {
	result := make(map[int][]int, len(obstacles))

	tree := &MinTree{}

	for i := len(obstacles) - 1; i >= 0; i-- {
		tree.Insert(obstacles[i])
	}

	tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Value, " ") })
	fmt.Println()
	fmt.Println(tree.Root.Left.Right.Left.Left.Value)
	//fmt.Println("         ", tree.Root.Value)
	//fmt.Println("       ", tree.Root.Left.Value, " ", tree.Root.Right.Value)
	//fmt.Println("     ", tree.Root.Left.Left.Value, "  ", tree.Root.Right.Value)

	for i := 0; i < len(obstacles); i++ {
		result[i] = obstacles[:i+1]
	}

	fmt.Println(result)

	return nil
}

func main() {

	r := longestObstacleCourseAtEachPosition([]int{2, 2, 2, 1, 3, 4})

	fmt.Println(r)
}

// 1 2 3 2
// 1 12 123 122

// 3 1 5 6 4 2
//

// 3 3
// 31 1
// 315 15
// 3156 356
// 31564 14
// 315642 12

// c = 0
// 1
// 5 2 2 1 2
// 1 2 3 2 2 1 2 2 2 1 2 2            (8)

// if last elem -> insert in longest subMinTree
//       2
// 	   2
//   1   2
//	1	2  2
// 1    2
//	  2
//   2
//   2
//  1
//               111111111111 1  2 1 22222222 1  2  2  3 4 4

//  	head	 2
// 	head1      2
//  head2    1   2 head3
//              2  head4
//

//if elem <= head.v && head.left == nil; insert to left; else go right
//if elem > head2.v; insert right
//if elem <= head3.v; insert left

// if elem > head[i].v && elem >= head[i - 1].v; insert right 4
