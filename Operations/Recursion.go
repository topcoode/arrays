package operations

import "github.com/gin-gonic/gin"

func Recursion(c *gin.Context) {
	arr := []int{34, 45, 7, 12, 78}
	target := 7
	index := 0
	Recursioned(arr, target, index)

}
func Recursioned(arr []int, target int, index int) bool {
	// Base case: if index is out of bounds
	if index >= len(arr) {
		return false
	}

	// Base case: if the target element is found
	if arr[index] == target {
		return true
	}

	// Recursive case: move to the next element
	return Recursioned(arr, target, index+1)
}
