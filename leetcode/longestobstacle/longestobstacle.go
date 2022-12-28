package main

import (
	"errors"
	"fmt"
	"log"
	"sort"
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

var sizes = []int{}

func largestSubtree(n *Node) int {

	if n == nil {
		return 0
	}

	if n.Left == nil && n.Right == nil {
		return 1
	}

	left := largestSubtree(n.Left)
	right := largestSubtree(n.Right)

	fmt.Println(left, right)

	return 0
}

func treeSolution(obstacles []int) {
	tree := &MinTree{}

	for i := len(obstacles) - 1; i >= 0; i-- {
		tree.Insert(obstacles[i])
	}

	tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Value, " ") })
	fmt.Println()

	_ = largestSubtree(tree.Root)

	fmt.Println(sizes)

	fmt.Println("					", tree.Root.Value)                            // 4
	fmt.Println("				", tree.Root.Left.Value)                        // 4
	fmt.Println("			", tree.Root.Left.Left.Value)                    // 3
	fmt.Print("		", tree.Root.Left.Left.Left.Value)                  // 1
	fmt.Println("		", tree.Root.Left.Left.Right.Value)               // 2
	fmt.Println("			", tree.Root.Left.Left.Right.Left.Value)         // 2
	fmt.Println("		", tree.Root.Left.Left.Right.Left.Left.Value)     // 2
	fmt.Println("	", tree.Root.Left.Left.Right.Left.Left.Left.Value) // 2
}

/*












 */

func reverseSlice(slice *[]int) {
	for i, j := 0, len(*slice)-1; i < j; i, j = i+1, j-1 {
		(*slice)[i], (*slice)[j] = (*slice)[j], (*slice)[i]
	}
}

func longestObstacleCourseAtEachPosition(obstacles []int) []int {
	result := make([]int, len(obstacles))

	for i := 0; i < len(obstacles); i++ {
		//tmp := make([]int, len(obstacles[:i+1]))
		//
		//copy(tmp, obstacles[:i+1])
		//reverseSlice(&tmp)
		//
		//result[i] = LDS(tmp)
		result[i] = LDS(obstacles[:i+1])
	}

	return result
}

// LDS FIRST SOLUTION
//include := nums[0]
//
//if len(nums) == 0 {
//	return 0
//}
//
//subSolutions := make([]int, len(nums))
//
//subSolutions[0] = 1
//
//for i := 1; i < len(nums); i++ {
//	for j := 0; j < i; j++ {
//		if nums[j] >= nums[i] && subSolutions[j] >= subSolutions[i] && nums[j] <= include {
//			subSolutions[i] = subSolutions[j]
//		}
//	}
//	subSolutions[i] = subSolutions[i] + 1
//}
func LDS(a []int) int {
	include := a[len(a)-1]
	c := make([]int, len(a))

	c[len(a)-1] = 1

	for i := len(a) - 2; i >= 0; {
		for j := len(a) - 1; j > i; {
			if a[j] >= a[i] && c[j] >= c[i] && a[j] <= include {
				c[i] = c[j]
			}
			j--
		}
		c[i] = c[i] + 1
		i--
	}

	max := func(slice []int) int {
		if len(slice) == 0 {
			log.Fatalln("Empty array")
		}

		max := slice[0]
		for _, v := range slice {
			if v > max {
				max = v
			}
		}
		return max
	}

	return max(c)
}

func longestObstacleCourseAtEachPosition_(obstacles []int) []int {
	var stack []int

	answer := make([]int, len(obstacles))

	for i, obstacle := range obstacles {
		fmt.Println(obstacle)
		if len(stack) == 0 || obstacle >= stack[len(stack)-1] {

			stack = append(stack, obstacle)
			answer[i] = len(stack)
			fmt.Println("stack cont", stack, i, obstacle)

			continue
		}
		index := sort.Search(len(stack), func(j int) bool {
			return stack[j] > obstacle
		})
		fmt.Println("index", stack, "index:", index, i)

		answer[i] = index + 1
		stack[index] = obstacle

		fmt.Println("stack", stack, i)
	}
	return answer
}

// 5 s = [5], a = [1]
// 1 s = [1], a = [1]
// 5 s = [1,5], a = [2]
// 5 s = [1,5,5], a = [3]
// 1 s = []

func main() {

	r := longestObstacleCourseAtEachPosition_([]int{5, 1, 5, 5, 1, 3, 4, 5, 1, 4}) //[1 1 2 3 2 3 4 5 3 4] (5)

	fmt.Println("R:", r) //LDS([]int{4, 1, 5, 4, 3, 1, 5, 5, 1, 5})

	//arr := []int{5, 1, 5, 5, 1, 3, 4, 5, 1, 4} //[1,1,2,3,2,3,4,5,3,5]
	//var ans []int
	//
	//for i := 0; i < len(arr); i++ {
	//	tmp := make([]int, len(arr[:i+1]))
	//	copy(tmp, arr[:i+1])
	//	reverseSlice(&tmp)
	//
	//	var sol []int
	//
	//	for _, v := range tmp {
	//
	//		if len(sol) == 0 || sol[len(sol)-1] >= v {
	//			fmt.Println("if", tmp, sol, v)
	//			sol = append(sol, v)
	//			continue
	//		}
	//		fmt.Println("after continue", tmp, sol, v)
	//
	//		idx := sort.Search(len(tmp), func(j int) bool {
	//			return tmp[j] <= v
	//		})
	//
	//		fmt.Println("IDX", idx, v)
	//		if v > sol[len(sol)-1] {
	//			continue
	//		} else {
	//			fmt.Println("else", sol)
	//			sol = append(sol, tmp[idx])
	//		}
	//		fmt.Println("Biba", tmp, v, sol)
	//
	//	}
	//	fmt.Println("solution", sol)
	//	ans = append(ans, len(sol))
	//}
	//fmt.Println(ans)

	//tmp := []int{2, 3, 2, 1}
	//
	//idx := sort.Search(len(tmp), func(j int) bool {
	//	return tmp[j] < 3
	//})
	//fmt.Println(idx, ans, tmp[idx])

	//var sol []int
	//for _, v := range tmp {
	//
	//	if len(sol) == 0 || sol[len(sol)-1] >= v {
	//		sol = append(sol, v)
	//		continue
	//	}
	//	fmt.Println("after continue", tmp, sol, v)
	//
	//	idx := sort.Search(len(tmp), func(j int) bool {
	//		return tmp[j] <= v
	//	})
	//
	//	fmt.Println("IDX", idx, v, tmp)
	//	if idx == len(tmp) {
	//		break
	//	} else {
	//		fmt.Println("else", sol)
	//		sol = append(sol, tmp[idx])
	//	}
	//	fmt.Println("Biba", tmp, v, sol)
	//
	//}
	//fmt.Println("solution", sol)
	//fmt.Println(ans)

}

/*
class Solution {
public:
    vector<int> longestObstacleCourseAtEachPosition(vector<int>& nums) {

        int n = nums.size();
        // lis store elements of longest increasing subsequence till ith
        vector<int> lis;
        // ans[i] store, no of elements satisfying the condition including ith
        vector<int> ans;
        for(int i = 0; i < n; i++)
        {
            int idx = upper_bound(lis.begin(), lis.end(), nums[i]) - lis.begin();
            ans.push_back(idx + 1);
            if(idx == lis.size())
                lis.push_back(nums[i]);
            else
                lis[idx] = nums[i];
        }
        return ans;
    }
};
*/
