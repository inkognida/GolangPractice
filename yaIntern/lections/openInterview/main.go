package main

import "fmt"

// first task
func reverse(nums *[]int) {
	f := 0
	l := len(*nums) - 1

	for f < l {
		tmp := (*nums)[f]
		(*nums)[f] = (*nums)[l]
		(*nums)[l] = tmp
		f++
		l--
	}
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

type Tree struct {
	Root *Node
}

// second task
func sumTree(t *Tree, ch chan int) int {

	sum1 := 0
	var walker func(t *Node)
	walker = func(n *Node) {
		if n == nil {
			return
		}
		walker(n.Left)
		ch <- n.Val

		sum1 += n.Val
		walker(n.Right)
	}
	walker(t.Root)
	close(ch)

	sum := 0
	for v := range ch {
		sum += v
	}

	fmt.Println(sum1)
	return sum
}

func totalNodes(root *Node) int {
	if root == nil {
		return 0
	}
	left := totalNodes(root.Left)
	right := totalNodes(root.Right)

	return left + right + 1
}

// third task
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func subSlice(nums []int) int {
	l := 0
	prev := 0
	mp := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		fmt.Println(mp, prev, l)
		prev = max(prev, mp[nums[i]])
		l = max(l, i-prev+1)
		mp[nums[i]] = i + 1

	}

	return l
}

func main() {
	arr := []int{3, 4, 5, 1, 8}
	reverse(&arr)
	fmt.Println(arr)

	t := &Tree{Root: &Node{
		Val: 5,
		Right: &Node{
			Val:   10,
			Right: nil,
			Left:  nil,
		},
		Left: &Node{
			Val: 20,
			Right: &Node{
				Val: 8,
				Right: &Node{
					Val:   12,
					Right: nil,
					Left:  nil,
				},
				Left: nil,
			},
			Left: nil,
		},
	}}

	fmt.Println(totalNodes(t.Root))
	ch := make(chan int, totalNodes(t.Root))
	fmt.Println(sumTree(t, ch))

	fmt.Println(subSlice([]int{1, 2, 3, 4, 4, 5, 5, 5, 6, 7, 8, 9, 10}))
	fmt.Println(subSlice([]int{5, 9, 1, 2, 3, 9, 7, 8, 5, 4, 6, 0, 9, 5, 3}))

}
