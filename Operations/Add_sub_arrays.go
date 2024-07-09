package operations

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Adding_Subtration(c *gin.Context) {
	arr := [4]int{54, 55, 56, 57}
	fmt.Println("before:", arr)
	adding := []int{}
	for _, value := range arr {
		adding = append(adding, value)
	}
	adding = append(adding, 56)
	fmt.Println("after", adding)
}
