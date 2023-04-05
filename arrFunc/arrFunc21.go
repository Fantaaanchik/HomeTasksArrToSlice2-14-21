package arrFunc

import (
	"fmt"
	"homeTasksArrToSlice/inoutPut"
)

func FuncArr21() {

	var (
		arithmeticMean int
		justNum        int
	)

	inoutPut.InputArr21()

	for i := inoutPut.K; i <= inoutPut.L; i++ {
		arithmeticMean += i
		justNum += 1
	}
	arithmeticMean = arithmeticMean / justNum
	fmt.Println(arithmeticMean)

	return
}
