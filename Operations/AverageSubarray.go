package operations

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func AverageSubarray(c *gin.Context) {
	arr := []int{75, 86, 84, 5, 74, 1, 86, 5}
	averagedata := 0
	lengthdata := len(arr[:4])
	average := 0
	for i := 0; i < lengthdata; i++ {
		averagedata = averagedata + arr[i]
		fmt.Println("averagedata  inner", averagedata)
	}
	fmt.Println("averagedata outer;", averagedata)
	average = averagedata / lengthdata
	fmt.Println(average)

}
