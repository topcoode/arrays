package operations

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ThreeD_operations(c *gin.Context) {
	values := [2][2][3]int{{{1, 2, 3}, {4, 5, 6}}, {{7, 8, 9}, {10, 11, 12}}}
	fmt.Println(values)
	//length := len(values)
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 3; k++ {
				fmt.Print(values[i][j][k])
			}
		}
	}
}
