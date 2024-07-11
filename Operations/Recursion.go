package operations

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Recursion_main(c *gin.Context) {
	i := 1
	j := 3
	Recursion(i, j)

	//Printing the linearly 1,2,3,4
	k := 1
	f := 4
	Print_Linerally(k, f)

	// printing the back lineraly 4,3,2,1
	e, p := 4, 4
	Print_BackLineraly(e, p)

}
func Recursion(i int, j int) int {
	if i > j {
		return 0
	}
	fmt.Println("code")
	return Recursion(i+1, j)
}
func Print_Linerally(i int, j int) int {
	if i > j {
		return 0
	}
	fmt.Println("-----------------")
	/*
		1,4
		2,4
		3,4
		4,4
		5,4 - not equal
	*/
	fmt.Println(i)
	return Print_Linerally(i+1, j)
}

func Print_BackLineraly(i int, j int) {
	if i < 1 {
		return
	}
	/*
	   4,4
	*/
	fmt.Println(i)
	Print_BackLineraly(i-1, j)
}
