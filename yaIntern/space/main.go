package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	var in []byte
	var e error
	file, _ := os.Open("input.txt")
	in, _ = ioutil.ReadAll(file)
	if e != nil {
		in, _ = ioutil.ReadAll(os.Stdin)
	}
	// TODO IMPLEMENT IT IN BYTES
	lines := bytes.Split(in, []byte("\n"))
	countN, _ := strconv.ParseInt(string(lines[0]), 10, 16)
	minInc_ := strings.Fields(string(lines[1]))
	hi_ := strings.Fields(string(lines[2]))
	childB_ := strings.Fields(string(lines[3]))
	minInc := make([]int64, countN)
	hie := make([]int64, countN)
	childBos := make([]int64, countN)
	for i := int64(0); i < countN; i++ {
		mi, _ := strconv.ParseInt(minInc_[i], 10, 16)
		minInc[i] = mi
		h, _ := strconv.ParseInt(hi_[i], 10, 16)
		hie[i] = h
		cb, _ := strconv.ParseInt(childB_[i], 10, 16)
		childBos[i] = cb
	}
	matesCount, _ := strconv.ParseInt(string(lines[4]), 10, 16)
	matesInc_ := strings.Fields(string(lines[5]))
	matesH_ := strings.Fields(string(lines[6]))
	matesBo_ := strings.Fields(string(lines[7]))
	matesInc := make([]int64, matesCount)
	matesHi := make([]int64, matesCount)
	matesBos := make([]int64, matesCount)
	for i := int64(0); i < matesCount; i++ {
		mi, _ := strconv.ParseInt(matesInc_[i], 10, 16)
		matesInc[i] = mi
		mh, _ := strconv.ParseInt(matesH_[i], 10, 16)
		matesHi[i] = mh
		mb, _ := strconv.ParseInt(matesBo_[i], 10, 16)
		matesBos[i] = mb
	}
	res := make([]int64, matesCount)
Main:
	for i := int64(0); i < matesCount; i++ {
		for j := int64(0); j < countN; j++ {
			if (childBos[j] == 1 && matesBos[i] == j+1) || (matesInc[i] >= minInc[j] && matesHi[i] >= hie[j]) {
				res[i] = j + 1
				continue Main
			}
		}
		res[i] = 0
	}
	for i := int64(0); i < matesCount-1; i++ {
		fmt.Print(res[i], " ")
	}
	fmt.Println(res[matesCount-1])
}
