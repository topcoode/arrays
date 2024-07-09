package operations

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Slicedata struct {
	Data []int `json:"data"`
}

func BubbleSort(c *gin.Context) {
	var slicecdata Slicedata
	if err := c.ShouldBindJSON(&slicecdata); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("before-->", slicecdata)
	n := len(slicecdata.Data)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if slicecdata.Data[j] > slicecdata.Data[j+1] {
				slicecdata.Data[j], slicecdata.Data[j+1] = slicecdata.Data[j+1], slicecdata.Data[j]

			}
			fmt.Print(slicecdata)
		}
	}
	fmt.Println("after-->", slicecdata)
	c.JSON(http.StatusOK, slicecdata)
}
