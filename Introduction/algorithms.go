package main

import "fmt"

func adding() []int {
	arr := []int{3, 0, 2, 0, 0, 0, 4, 1}
	for i := len(arr)-1; i >= 0; i-- {
		if arr[i] == 0 {
			arr = append(arr[:i], arr[i+1:]...)
			arr = append(arr, append([]int{0}, arr[len(arr):]...)...)
		}
	}
	return arr
}

/* Ответ, конечно, n**3.
Получаем его так:
1. строки нумеруются с 1
2. строка номер n начинается с числа n*n - n +1, в ней n чисел, значение последнего числа n*n - n +1 + (n - 1)*2
3. Среднее значение числа в строке (n*n - n +1 + n*n - n +1 + (n - 1)*2)/2 = n**2
4. Сумма в строке: n (чисел) * n**2 = n**3
*/

func rowSum(n int) int{
	return n * n * n
}
func rowSumv2(n int) int {
	rows := []int{}
	nRow := []int{}

	if n > 1 {
		for i := 1; i < n*(n-1); i += 2 {
			rows = append(rows, i)
		}
		for i := rows[len(rows)-1]+2; i < (rows[len(rows)-1]+(n*2))+2; i += 2 {
			nRow = append(nRow, i)
		}
		sum := 0
		for i := 0; i < len(nRow); i++ {
			sum += nRow[i]
		}
		return sum
	} else {
		return 1
	}
}


func main() {
	addArray := adding()
	fmt.Println(addArray)

	sumN := rowSum(2)
	sumNv2 := rowSumv2(2)
	fmt.Println(sumN, sumNv2)

}