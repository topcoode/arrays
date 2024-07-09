package operations

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Factorial_Calculation(c *gin.Context) {
	number := 5
	result := factorial(number)
	fmt.Printf("The factorial of %d is %d\n", number, result)
}
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
