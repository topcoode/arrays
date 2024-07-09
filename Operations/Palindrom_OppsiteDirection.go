package operations

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Palindrom_OppsiteDirection(c *gin.Context) {
	arr := []int{6, 0, 5, 6}
	start := 0
	end := len(arr) - 1
	for start < end {
		if arr[start] == arr[end] {
			start++
			end--
		} else {
			fmt.Println("invalid palindome")
			break
		}

		fmt.Println("its a palindrome")

	}
}
