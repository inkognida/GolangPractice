package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

type Tree struct {
	Root *Node
}

func (t *Tree) Insert(value int) error {
	if t.Root == nil {
		t.Root = &Node{Value: value}
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

//tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Value, " ") })
//fmt.Println()

func simpleOperation(piles []int) bool {

	//TREE INITIALIZATION
	tree := &Tree{}
	for _, v := range piles {
		err := tree.Insert(v)
		if err != nil {
			log.Fatalln(err)
		}
	}

	minElem, flag := tree.Root.findMinElement(tree.Root)
	if flag == false {
		log.Fatalln("Unable to find min element (empty tree)")
	}

	maxElem, flag := tree.Root.findMaxElement(tree.Root)
	if flag == false {
		log.Fatalln("Unable to find min element (empty tree)")
	}

	fmt.Println("Min and max elements:", minElem, maxElem)

	tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Value, " ") })
	fmt.Println()

	if err := tree.Delete(maxElem); err != nil {
		log.Fatalln("Unable to delete element", err)
	}

	tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Value, " ") })
	fmt.Println()

	if value, flag := tree.Find(minElem); flag == false {
		log.Fatalln("Unable to find element", flag)
	} else {
		fmt.Println("Found", value)
	}

	randomValue := rand.Intn(4124)
	if err := tree.Insert(randomValue); err != nil {
		log.Fatalln("Unable to insert value", err)
	}

	tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Value, " ") })
	fmt.Println()

	return true
}

//1962. Remove Stones to Minimize the Total

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7}

	_ = simpleOperation(values)

}
