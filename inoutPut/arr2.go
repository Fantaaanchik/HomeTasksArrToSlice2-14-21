package inoutPut

import "fmt"

var (
	NArr2  int
	Array2 = []int{}
)

func InPutArr2() {

	fmt.Printf("Array2. Дано целое число N (> 0). Сформировать и вывести целочисленный массив размера N, " +
		"содержащий степени двойки от первой до N-й: 2, 4, 8, 16, . . . .\n")
	fmt.Printf("Введите челое число:\n")
	_, err := fmt.Scan(&NArr2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return
}
