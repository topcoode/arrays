package operations

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Fibonacci_Series(c *gin.Context) {
	values := []int{0, 1}
	data := Fibonacci(values)
	c.JSON(http.StatusOK, data)
}

func Fibonacci(value []int) []int {
	if len(value) >= 10 {
		return value
	}
	last := value[len(value)-1]
	secondLast := value[len(value)-2]
	sum := last + secondLast
	data := append(value, sum)
	return Fibonacci(data)
}
