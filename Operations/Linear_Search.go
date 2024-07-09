package operations

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type input struct {
	Data int `json:"data"`
}

func Linear_Search(c *gin.Context) {
	var inputdata input
	if err := c.ShouldBindJSON(&inputdata); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	arr := []int{45, 56, 67, 78, 90, 12, 34}

	for i := 0; i < len(arr); i++ {
		if arr[i] == inputdata.Data {
			fmt.Println("sorted the value", arr[i])
			break
		} else {
			fmt.Println("not found")
		}
	}
}
