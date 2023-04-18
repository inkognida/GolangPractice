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

	countryNumbers := readInt(scanner)
	minimumIncomeForCountry := make(map[int]int, countryNumbers)
	higherEducationForCountry := make([]int, 0)
	minimumIncomeForCountryWithoutEducation := make([]int, 0)
	childBoostCountry := make(map[int]struct{}, countryNumbers)

	minCountryNumberWithHigherEducation := 0
	minCountryNumberWithoutEducation := 0
	for i := 0; i < countryNumbers; i++ {
		minimumIncomeForCountry[i+1] = readInt(scanner)
	}
	for i := 0; i < countryNumbers; i++ {
		h := readInt(scanner)
		if h != 0 {
			if minCountryNumberWithHigherEducation == 0 {
				minCountryNumberWithHigherEducation = i + 1
			}
			higherEducationForCountry = append(higherEducationForCountry, i+1)
		} else {
			if minCountryNumberWithoutEducation == 0 {
				minCountryNumberWithoutEducation = i + 1
			}
			minimumIncomeForCountryWithoutEducation = append(minimumIncomeForCountryWithoutEducation, i+1)
		}
	}
	for i := 0; i < countryNumbers; i++ {
		cb := readInt(scanner)
		if cb != 0 {
			childBoostCountry[i+1] = struct{}{}
		}
	}
	matesCount := readInt(scanner)
	matesIncomes := make([]int, matesCount)
	matesHigherEducation := make([]int, matesCount)
	matesParentBoost := make([]int, matesCount)

	for i := 0; i < matesCount; i++ {
		matesIncomes[i] = readInt(scanner)
	}
	for i := 0; i < matesCount; i++ {
		matesHigherEducation[i] = readInt(scanner)
	}
	for i := 0; i < matesCount; i++ {
		matesParentBoost[i] = readInt(scanner)
	}

	result := make([]int, matesCount)
Main:
	for i := 0; i < matesCount; i++ {
		if _, ok := childBoostCountry[matesParentBoost[i]]; ok {
			result[i] = matesParentBoost[i]
			if result[i] < minCountryNumberWithHigherEducation && result[i] < minCountryNumberWithoutEducation {
				continue Main
			}
		}
		for j := 0; j < len(higherEducationForCountry); j++ {
			if matesHigherEducation[i] == 1 {
				id := higherEducationForCountry[j]
				if minimumIncomeForCountry[id] <= matesIncomes[i] && (id < result[i] || result[i] == 0) {
					result[i] = id
					break
				}
			}
		}
		if result[i] < minCountryNumberWithoutEducation && result[i] != 0 {
			continue Main
		}

		for j := 0; j < len(minimumIncomeForCountryWithoutEducation); j++ {
			id := minimumIncomeForCountryWithoutEducation[j]
			v := minimumIncomeForCountry[id]
			if v <= matesIncomes[i] && (id < result[i] || result[i] == 0) {
				result[i] = id
				break
			}
		}
	}
	for i := 0; i < matesCount-1; i++ {
		fmt.Print(result[i], " ")
	}
	fmt.Println(result[matesCount-1])
}
