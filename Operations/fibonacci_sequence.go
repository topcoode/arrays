package operations

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func sum(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	mid := len(nums) / 2

	leftSum := sum(nums[:mid])
	rightSum := sum(nums[mid:])

	return leftSum + rightSum
}

func fibonacci_sequence(c *gin.Context) {
	numbers := []int{1, 2, 3, 4, 5}
	total := sum(numbers)
	fmt.Println("Sum:", total)
}
