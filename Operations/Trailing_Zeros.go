package operations

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Trailing_Zero(c *gin.Context) {
	factorial := 25
	factdata := Trailing_Zero_Main(factorial)
	fmt.Println(factdata)

}
func Trailing_Zero_Main(factorial int) int {
	/*
		if factorial == 0 {
			return 1
		}
		result := factorial * Trailing_Zero_Main(factorial-1)
		fmt.Println(result)
		return result
	*/
	n := factorial
	res := 0
	Powerof5 := 5
	//
	for Powerof5 < n {
		res = res + n/Powerof5
		Powerof5 = Powerof5 + 5
	}
	return res
}
