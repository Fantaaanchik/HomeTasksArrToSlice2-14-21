package arrFunc

import (
	"fmt"
	"homeTasksArrToSlice/inoutPut"
)

func FuncArr2() {
	var (
		step = 1
	)
	inoutPut.InPutArr2()

	for i := 1; len(inoutPut.Array2) < inoutPut.NArr2; i++ {
		step = step * 2
		inoutPut.Array2 = append(inoutPut.Array2, step)
	}
	fmt.Println(inoutPut.Array2)

	return
}
