package main

import (
	"bufio"
	"fmt"
	"github.com/umpc/go-sortedmap"
	"os"
	"strconv"
	"strings"
)

type Condition struct {
	I       int
	R       *int
	A       *int
	Servers []int
}

func main() {
	var r *bufio.Reader

	f, _ := os.Open("input.txt")
	r = bufio.NewReader(f)

	mnqInput, _ := r.ReadString('\n')
	mnq := strings.Split(strings.TrimSpace(mnqInput), " ")
	m, _ := strconv.Atoi(mnq[0])
	n, _ := strconv.Atoi(mnq[1])
	q, _ := strconv.Atoi(mnq[2])

	less := func(i, j interface{}) bool {
		return i.(Condition).I < j.(Condition).I
	}

	dataCenters := sortedmap.New(n, less)
	for i := 0; i < n; i++ {
		servers := make([]int, m)
		for j, _ := range servers {
			servers[j] = 1
		}

		Rtmp := 0
		Atmp := m

		dataCenters.Insert(i+1, Condition{
			I:       i + 1,
			R:       &Rtmp,
			Servers: servers,
			A:       &Atmp,
		})
	}

	GETMIN := 0
	GETMAX := 0
	fmt.Println(*dataCenters.Map()[1].(Condition).A)
	fmt.Println(*dataCenters.Map()[2].(Condition).A)
	fmt.Println(*dataCenters.Map()[3].(Condition).A)

	for i := 0; i < q; i++ {
		infoInput, _ := r.ReadString('\n')
		info := strings.Split(strings.TrimSpace(infoInput), " ")
		if info[0] == "DISABLE" {
			centerId, _ := strconv.Atoi(info[1])
			serverId, _ := strconv.Atoi(info[2])
			serverId--
			val, _ := dataCenters.Get(centerId)
			if val.(Condition).Servers[serverId] != 0 {
				val.(Condition).Servers[serverId] = 0
				*(val.(Condition).A)--
				*(val.(Condition).R) *= *(val.(Condition).A)
			}
		}
	}

	fmt.Println(*dataCenters.Map()[1].(Condition).A)
	fmt.Println(*dataCenters.Map()[2].(Condition).A)
	fmt.Println(*dataCenters.Map()[3].(Condition).A)

	fmt.Println(GETMAX, GETMIN)
}
