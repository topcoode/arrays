package operations

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Input struct {
	Data string `json:"data"`
}

func StringReverse(c *gin.Context) {
	var input Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	intdata := []rune(input.Data)
	fmt.Println("before:", intdata)

	for i, j := 0, len(intdata)-1; i < j; i, j = i+1, j-1 {
		intdata[i], intdata[j] = intdata[j], intdata[i]
	}

	fmt.Println("after:", intdata)
	stringdata := string(intdata)
	fmt.Println(stringdata)

	c.JSON(http.StatusOK, gin.H{"reversed_data": stringdata})
}
