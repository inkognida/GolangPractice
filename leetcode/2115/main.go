package main

import "fmt"

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	reps := make(map[string]struct{}, 0)
	sups := make(map[string]struct{}, 0)
	for _, s := range supplies {
		sups[s] = struct{}{}
	}

	for i, r := range recipes {
		c := 0
		for _, ing := range ingredients[i] {
			if _, ok := sups[ing]; ok {
				c++
			}
		}
		if c == len(ingredients[i]) {
			reps[r] = struct{}{}
		}
	}

	for i, r := range recipes {
		c := 0
		for _, ing := range ingredients[i] {
			if _, ok := sups[ing]; ok {
				c++
			} else if _, ok := reps[ing]; ok {
				c++
			}
		}
		if c == len(ingredients[i]) {
			reps[r] = struct{}{}
		}
	}

	ans := make([]string, 0)
	for k, _ := range reps {
		ans = append(ans, k)
	}

	return ans
}

func main() {

	r := []string{"bread", "sandwich", "burger"}
	var ing [][]string
	ing = [][]string{
		{"yeast", "flour"},
		{"bread", "meat"},
		{"sandwich", "meat", "bread"},
	}
	s := []string{"yeast", "flour", "meat"}

	fmt.Println(findAllRecipes(r, ing, s))
}
