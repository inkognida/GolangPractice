package main

import (
	"errors"
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
	case value == n.Value:
		return nil

	case value < n.Value:
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
