package operations

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Input_array struct {
	Data []int
}

func Merge_sort(c *gin.Context) {
	var input Input_array
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	length := len(input.Data)
	for i := 0; i <= length; i++ {
		subarrays := length / 2
		fmt.Println(subarrays)
	}

}
