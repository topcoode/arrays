package operations

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInput struct {
	Data []int `json:"data"`
}

func SelectionSort(c *gin.Context) {
	var input UserInput
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	n := len(input.Data)
	for i := 0; i < n-1; i++ {
		// Find the minimum element in unsorted array
		minIndex := i
		for j := i + 1; j < n; j++ {
			if input.Data[j] < input.Data[minIndex] {
				minIndex = j
			}
		}
		// Swap the found minimum element with the first element
		if minIndex != i {
			input.Data[i], input.Data[minIndex] = input.Data[minIndex], input.Data[i]
		}
	}

	fmt.Println(input.Data)
	c.JSON(http.StatusOK, gin.H{"sorted_data": input.Data})
}
