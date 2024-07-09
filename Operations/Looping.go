package operations

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Loops(c *gin.Context) {
	// Bubble sort in descending order
	data := []int{56, 67, 89, 90}
	n := len(data)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if data[j] < data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}

	// Respond with the sorted data
	c.JSON(http.StatusOK, data)
}
