//package main
//
//import (
//	"bufio"
//	"context"
//	"fmt"
//	"log"
//	"os"
//	"strconv"
//	"strings"
//	"tfs-grpc/3_grpc_candles/internal/candlespb"
//
//	"google.golang.org/grpc"
//)
//
//func main() {
//	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
//	if err != nil {
//		log.Fatalf("can't connect to server: %v", err)
//	}
//	client := candlespb.NewCandlesServiceClient(conn)
//	fmt.Printf("created client: %v", client)
//
//	client.Candles(context.Background(), &candlespb.CandleRequest{
//		Instrument: "",
//		Period:     0,
//	})
//
//	conn.Close()
//}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SplitBySeparator(s string, sep rune) []string {
	return strings.FieldsFunc(s, func(r rune) bool {
		return r == sep
	})
}

type Country struct {
	MinInc int
	HighEd string
	ChildB string
}

type Mate struct {
	Inc    int
	HighEd string
	ParB   int // parent boost
}

func relocate(mate Mate, country Country, countryPosition int) bool {
	if (country.ChildB == "1" && mate.ParB == countryPosition) ||
		mate.Inc >= country.MinInc && mate.HighEd >= country.HighEd {
		return true
	}
	return false
}

func main() {
	// Open the input file
	var r *bufio.Reader

	if file, err := os.Open("input.txt"); err != nil {
		r = bufio.NewReader(os.Stdin)
	} else {
		r = bufio.NewReader(file)
	}

	// country count
	cN, _ := r.ReadString('\n')
	cN_ := SplitBySeparator(strings.TrimSpace(cN), ' ')
	countN, _ := strconv.Atoi(cN_[0])

	countries := make([]Country, countN)

	// min income to move to contry
	mInc, _ := r.ReadString('\n')
	minInc_ := SplitBySeparator(strings.TrimSpace(mInc), ' ')
	for i, v := range minInc_ {
		pv, _ := strconv.Atoi(v)
		countries[i].MinInc = pv
	}

	// higher education if 0 <= i <= 1
	hi, _ := r.ReadString('\n')
	hi_ := SplitBySeparator(strings.TrimSpace(hi), ' ')

	for i, v := range hi_ {
		countries[i].HighEd = v
	}

	// child boost, if 1 - we can skip higher education and min income
	childB, _ := r.ReadString('\n')
	childB_ := SplitBySeparator(strings.TrimSpace(childB), ' ')
	for i, v := range childB_ {
		countries[i].ChildB = v
	}

	// classmates
	clsQ, _ := r.ReadString('\n')
	clsQ_ := SplitBySeparator(strings.TrimSpace(clsQ), ' ')
	matesCount, _ := strconv.Atoi(clsQ_[0])

	classmates := make([]Mate, matesCount)

	// classmates income
	matesIn, _ := r.ReadString('\n')
	matesInc_ := SplitBySeparator(strings.TrimSpace(matesIn), ' ')
	for i, v := range matesInc_ {
		pv, _ := strconv.Atoi(v)
		classmates[i].Inc = pv
	}

	// classmates higher education
	matesH, _ := r.ReadString('\n')
	matesH_ := SplitBySeparator(strings.TrimSpace(matesH), ' ')
	for i, v := range matesH_ {
		classmates[i].HighEd = v
	}

	// classmates parents boost
	matesBo, _ := r.ReadString('\n')
	matesBo_ := SplitBySeparator(strings.TrimSpace(matesBo), ' ')
	for i, v := range matesBo_ {
		pv, _ := strconv.Atoi(v)
		classmates[i].ParB = pv
	}

	res := make([]int, 0)
Main:
	for _, m := range classmates {

		for j, c := range countries {
			if relocate(m, c, j+1) {
				res = append(res, j+1)
				continue Main
			}

			if j == countN-1 {
				res = append(res, 0)
			}

		}
	}

	for i := 0; i < len(res)-1; i++ {
		fmt.Print(res[i], " ")
	}
	fmt.Println(res[len(res)-1])
}

	cN, _ := r.ReadString('\n')
	cN_ := spl(cN[:len(cN)-1], ' ')
	countN, _ := strconv.Atoi(cN_[0])

	mInc, _ := r.ReadString('\n')
	minInc_ := spl(mInc[:len(mInc)-1], ' ')
	minInc := make([]int, countN)

	hi, _ := r.ReadString('\n')
	hi_ := spl(hi[:len(hi)-1], ' ')
	hie := make([]int, countN)

	childB, _ := r.ReadString('\n')
	childB_ := spl(childB[:len(childB)-1], ' ')
	childBos := make([]int, countN)

	for i := 0; i < countN; i++ {
		mi, _ := strconv.Atoi(minInc_[i])
		minInc[i] = mi

		h, _ := strconv.Atoi(hi_[i])
		hie[i] = h

		cb, _ := strconv.Atoi(childB_[i])
		childBos[i] = cb

	}

	clsQ, _ := r.ReadString('\n')
	clsQ_ := spl(strings.TrimSpace(clsQ), ' ')
	matesCount, _ := strconv.Atoi(clsQ_[0])

	matesIn, _ := r.ReadString('\n')
	matesInc_ := spl(strings.TrimSpace(matesIn), ' ')
	matesInc := make([]int, matesCount)

	matesH, _ := r.ReadString('\n')
	matesH_ := spl(strings.TrimSpace(matesH), ' ')
	matesHi := make([]int, matesCount)

	matesBo, _ := r.ReadString('\n')
	matesBo_ := spl(strings.TrimSpace(matesBo), ' ')
	matesBos := make([]int, matesCount)

	for i := 0; i < matesCount; i++ {
		matesInc[i] = mi

		mh, _ := strconv.Atoi(matesH_[i])
		matesHi[i] = mh

		mb, _ := strconv.Atoi(matesBo_[i])
		matesBos[i] = mb
	}

	res := make([]int, matesCount)
Main:
	for i := 0; i < matesCount; i++ {

		for j := 0; j < countN; j++ {
			if (childBos[j] == 1 && matesBos[i] == j+1) ||
				(matesInc[i] >= minInc[j] && matesHi[i] >= hie[j]) {
				res[i] = j + 1
				continue Main
			}
		}
		res[i] = 0
	}

	for i := 0; i < matesCount-1; i++ {
		fmt.Print(res[i], " ")
	}
	fmt.Println(res[matesCount-1])
