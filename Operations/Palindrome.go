package operations

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Palindrome_str struct {
	Num interface{}
}

func Palindrome(c *gin.Context) {
	var requestbody Palindrome_str
	if err := c.ShouldBindJSON(&requestbody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	bool := Palindrome_data(requestbody.Num)
	fmt.Println(bool)
}
func Palindrome_data(value interface{}) bool {
	//  := value
	//var Revereddata int
	// Convert the interface to string
	strValue, ok := value.(string)
	if !ok {
		return false
	}
	fmt.Println("strvalue--->", strValue)
	runes := []rune(strValue)
	fmt.Println("runes--->", runes)
	n := len(runes)
	fmt.Println("n", n)
	for i := 0; i < n/2; i++ {
		fmt.Println(runes[i])
		fmt.Println(runes[n-1-i])
		if runes[i] != runes[n-1-i] {
			return false
		}
	}
	return true
}
