package operations

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SubarrayInput struct {
	Data []int `json:"data"`
}

func Sub_array(c *gin.Context) {
	var subarrayinput SubarrayInput
	if err := c.ShouldBindJSON(&subarrayinput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i := 0; i <= len(subarrayinput.Data); i++ {
		for j := i; j <= len(subarrayinput.Data); j++ {
			fmt.Println(subarrayinput.Data[i : j+1])
		}
	}
}
