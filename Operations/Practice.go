package operations

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Pratice(c *gin.Context) {
	arr := []int{45, 78, 23, 57, 8, 43}
	Sorting(arr)
}
func Sorting(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}
	length := len(arr)
	mid := length / 2
	left := arr[:mid]
	right := arr[mid:]
	fmt.Println(left, right)
	leftarray := Sorting(left)
	rightarray := Sorting(right)
	fmt.Println(leftarray, rightarray)
	return Merge(leftarray, rightarray)

}
func Merge(left []int, right []int) []int {
	fmt.Println("merge function:", left, right)
	return nil
}
