package operations

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CoutingPlaceValue(c *gin.Context) {
	value := 4528
	res := CoutingPlaceValue_main(value)
	fmt.Println("final result:", res)
}

func CoutingPlaceValue_main(n int) int {
	if n == 0 {
		return 0
	}
	return 1 + CoutingPlaceValue_main(n/10)
}
