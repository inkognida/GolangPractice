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
func subSlice(nums []int) int {
	nm := make(map[int]struct{}, len(nums))
	max := 0
	//Main:
	//	for i := 0; i < len(nums); i++ {
	//		nm = make(map[int]struct{}, len(nums))
	//		nm[nums[i]] = struct{}{}
	//
	//		for j := i + 1; j < len(nums); j++ {
	//			fmt.Println(nm, nums[j])
	//			if _, ok := nm[nums[j]]; !ok {
	//				nm[nums[j]] = struct{}{}
	//			} else {
	//				if len(nm) > max {
	//					max = len(nm)
	//					continue Main
	//				}
	//			}
	//		}
	//		if len(nm) > max {
	//			max = len(nm)
	//			continue Main
	//		}
	//	}

	for i, v := range nums {
		if _, ok := nm[v]; !ok {
			nm[v] = struct{}{}
			if i == len(nums)-1 && len(nm) > max {
				max = len(nm)
			}
		} else {
			if len(nm) > max {
				max = len(nm)
				nm = make(map[int]struct{}, 0)
				nm[v] = struct{}{}
			}
		}
	}

	return max
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

	fmt.Println(subSlice([]int{1, 2, 3, 4, 4, 5, 6, 7, 8, 9, 10}))
}
