package operations

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type inputdata struct {
	Data   []int `json:"data"`
	Target int   `json:"target"`
}

func BinarySearch(c *gin.Context) {
	var input inputdata
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	// input := []int{98,65,6,70,12}
	length := len(input.Data)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if input.Data[i] < input.Data[j] {
				input.Data[i], input.Data[j] = input.Data[j], input.Data[i]
			}
		}
	}
	fmt.Println("ascending:", input)
	low := 0
	high := len(input.Data) - 1
	target := input.Target
	fmt.Println(target)
	for low <= high {
		mid := (low + high) / 2
		if input.Data[mid] == target {
			fmt.Println("tha value @ :", input.Data[mid])
		} else if input.Data[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
}
