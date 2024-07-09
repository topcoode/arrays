package operations

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Ascending_Decending(c *gin.Context) {
	arr := []int{12, 78, 4, 70, 57, 46, 98}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	fmt.Println("ascending order:", arr)

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	fmt.Println("decending order:", arr)
}
