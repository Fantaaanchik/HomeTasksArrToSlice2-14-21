package inoutPut

import "fmt"

var (
	NArr21 int
	K      int
	L      int
)

func InputArr21() {

	fmt.Printf("Array21. Дан массив размера N и целые числа K и L (1 ≤ K ≤ L ≤ N). Найти среднее арифметическое " +
		"элементов массива с номерами от K до L включительно\n")
	fmt.Printf("Введите три числа:\n")
	_, err := fmt.Scan(&NArr21, &K, &L)
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}
