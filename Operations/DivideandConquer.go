package operations

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Merge function to combine two sorted sub-arrays into one sorted array
func merge(left, right []int) []int {
	//make([]int,0,0+1=1)
	result := make([]int, 0, len(left)+len(right))
	fmt.Println("result:", result)
	i, j := 0, 0
	// for 0 < 1 && 0 < 1
	for i < len(left) && j < len(right) {
		//45 <=0 i++
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append remaining elements if any
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

// MergeSort function to sort an array using merge sort
func mergeSort(arr []int) []int {
	fmt.Println("arr---->", len(arr))
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	fmt.Println("left:", left)
	fmt.Println("right:", right)
	return merge(left, right)
}

func DivideAndConquer(c *gin.Context) {
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("Original array:", arr)
	sortedArr := mergeSort(arr)
	fmt.Println("Sorted array:  ", sortedArr)
}
