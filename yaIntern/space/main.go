package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	val, _ := strconv.Atoi(scanner.Text())
	return val
}

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	countN := readInt(scanner)
	minInc := make(map[int]int, countN)
	hie := make([]int, 0)
	minIncOutHie := make([]int, 0)
	childBos := make(map[int]struct{}, countN)

	minhi := 0
	minnoHi := 0
	for i := 0; i < countN; i++ {
		minInc[i+1] = readInt(scanner)
	}
	for i := 0; i < countN; i++ {
		h := readInt(scanner)
		if h != 0 {
			if minhi == 0 {
				minhi = i + 1
			}
			hie = append(hie, i+1)
		} else {
			if minnoHi == 0 {
				minnoHi = i + 1
			}
			minIncOutHie = append(minIncOutHie, i+1)
		}
	}
	for i := 0; i < countN; i++ {
		cb := readInt(scanner)
		if cb != 0 {
			childBos[i+1] = struct{}{}
		}
	}
	matesCount := readInt(scanner)
	matesInc := make([]int, matesCount)
	matesHi := make([]int, matesCount)
	matesBos := make([]int, matesCount)

	for i := 0; i < matesCount; i++ {
		matesInc[i] = readInt(scanner)
	}
	for i := 0; i < matesCount; i++ {
		matesHi[i] = readInt(scanner)
	}
	for i := 0; i < matesCount; i++ {
		matesBos[i] = readInt(scanner)
	}

	res := make([]int, matesCount)
Main:
	for i := 0; i < matesCount; i++ {
		if _, ok := childBos[matesBos[i]]; ok {
			res[i] = matesBos[i]
			if res[i] < minhi && res[i] < minnoHi {
				continue Main
			}
		}
		for j := 0; j < len(hie); j++ {
			if matesHi[i] == 1 {
				id := hie[j]
				if minInc[id] <= matesInc[i] && (id < res[i] || res[i] == 0) {
					res[i] = id
					break
				}
			}
		}
		if res[i] < minnoHi && res[i] != 0 {
			continue Main
		}

		for j := 0; j < len(minIncOutHie); j++ {
			id := minIncOutHie[j]
			v := minInc[id]
			if v <= matesInc[i] && (id < res[i] || res[i] == 0) {
				res[i] = id
				break
			}
		}
	}
	for i := 0; i < matesCount-1; i++ {
		fmt.Print(res[i], " ")
	}
	fmt.Println(res[matesCount-1])
}
