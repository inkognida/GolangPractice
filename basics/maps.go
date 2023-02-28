package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func initMap() {
	map_ := map[int]string{0: "Tom", 1: "Bob"}
	map__ := make(map[string]int)
	map__["Kate"] = 1
	map__["Dim"] = 2
	fmt.Println("Init map:", map_, map__, map__["Kate"])

	map___ := map[int]func() int{
		0: func() int { return 10 },
		1: func() int { return 11 },
	}
	fmt.Println("Function map:", map___[0])

	values := []int{1, 2, 3}
	multiValuesMap := make(map[int][]int)
	multiValuesMap[0] = values
	fmt.Println("Multi values map:", multiValuesMap)
}

func deleteCheck() {
	map_ := make(map[int]int, 3)
	for i := 0; i < 3; i++ {
		map_[i] = rand.Intn(10)
	}
	if value, inMap := map_[1000]; inMap { // does not exist
		fmt.Println("Value:", value)
	} else {
		fmt.Println("Value doesn't exist")
	}
	delete(map_, 0) // delete first key - values pair using key
	fmt.Println(map_)
}

var Days = map[int]string{
	1: "Monday",
	2: "Tuesday",
	3: "Wednesday",
	4: "Thursday",
	5: "Friday",
	6: "Saturday",
	7: "Sunday",
}

func findDay(n int) string {
	if day, state := Days[n]; state {
		return day
	}
	return ""
}

func sliceMap() {
	sliceOfMap := make([]map[int]int, 5)
	for i := range sliceOfMap {
		sliceOfMap[i] = make(map[int]int, 1)
		sliceOfMap[i][1] = 0
		sliceOfMap[i][2] = 1
	}
	fmt.Println("Slice of map:", sliceOfMap)
}

var (
	barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23, "delta": 87,
		"echo": 56, "foxtrot": 12, "golf": 34, "hotel": 16, "indio": 87, "juliet": 65, "kilo": 43, "lima": 98}
)

func sortKeys() {
	keys := make([]string, len(barVal))
	i := 0

	for key, _ := range barVal {
		keys[i] = key
		i++
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println("key:", key, "value:", barVal[key])
	}
}

func invertMap() {
	invMap := make(map[int]string, len(barVal))
	for k, v := range barVal {
		invMap[v] = k // key becomes value and value becomes key
	}
	for k, v := range invMap {
		fmt.Println(k, v)
	}
}

func main() {
	initMap()
	// deleteCheck()
	// fmt.Println("Day finder:", findDay(1))
	// sliceMap()
	// sortKeys()
	// invertMap()
}
