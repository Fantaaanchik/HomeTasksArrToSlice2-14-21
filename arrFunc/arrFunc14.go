package arrFunc

import (
	"fmt"
	"homeTasksArrToSlice/inoutPut"
)

func FuncArr14() {

	inoutPut.InputArr14()

	for i := 2; len(inoutPut.ArrayEven14) < inoutPut.NArr14; i += 2 {
		inoutPut.ArrayEven14 = append(inoutPut.ArrayEven14, i)
	}
	fmt.Println(inoutPut.ArrayEven14)
	
	for i := 1; len(inoutPut.ArrayOdd14) < inoutPut.NArr14; i += 2 {
		inoutPut.ArrayOdd14 = append(inoutPut.ArrayOdd14, i)
	}
	fmt.Println(inoutPut.ArrayOdd14)
	return
}
