package main

import (
	"container/list"
	"fmt"
	"regexp"
	"strconv"
	"sync"
	"time"

	"github.com/inancgumus/myhttp"
)

// double-linked list block

func listElems(n int) *list.List {
	lst := list.New()
	for i := 0; i <= n; i++ {
		lst.PushBack(i)
	}
	return lst
}

// regular expressions block

func regularExp() {
	searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	pattern := "[0-9]+.[0-9]+"
	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v*2, 'f', 2, 32)
	}
	if ok, _ := regexp.Match(pattern, []byte(searchIn)); ok {
		fmt.Println("Pattern found")
	}
	re, _ := regexp.Compile(pattern)
	str := re.ReplaceAllString(searchIn, "##.#")
	fmt.Println("Str:", str)
	strF := re.ReplaceAllStringFunc(searchIn, f)
	fmt.Println("Str using f func:", strF)
}

type Info struct {
	mutex_ sync.Mutex
	str_   string
}

func syncMutex(info *Info) {
	info.mutex_.Lock()
	info.str_ = "Hello, mutex World!"
	info.mutex_.Unlock()
}

func PackagaeImportUsageExample() {
	mh := myhttp.New(time.Second)
	response, _ := mh.Get("https://jsonip.com/")
	fmt.Println("HTTP status code: ", response.StatusCode)
}

func main() {
	lst := listElems(5)
	for i := lst.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	regularExp()

	var info_ Info
	info_.str_ = "Hello, world!"
	fmt.Println(info_.str_)
	syncMutex(&info_)
	fmt.Println(info_.str_)
	PackagaeImportUsageExample()

}
