package inoutPut

import "fmt"

var (
	NArr14      int
	ArrayOdd14  = []int{}
	ArrayEven14 = []int{}
)

func InputArr14() {

	fmt.Printf("Array14. Дан массив A размера N. Вывести вначале его элементы с четными номерами (в порядке " +
		"возрастания номеров), а затем — элементы с нечеными номерами (также в порядке возрастания номеров): A2, A4, " +
		"A6, . . ., A1, A3, A5, . . . .\nУсловный оператор не использовать.\n")
	fmt.Printf("Введите число N:\n")
	_, err := fmt.Scan(&NArr14)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return
}
