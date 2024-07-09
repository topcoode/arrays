package operations

import (
	"fmt"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PrimeNumber struct {
	Num int `json:"num" binding:"required"`
}

func Prime_Number(c *gin.Context) {
	var requestbody PrimeNumber
	if err := c.ShouldBindJSON(&requestbody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("requestbody--->", requestbody)
	primeornot := Prime(requestbody.Num)
	if primeornot == true {
		fmt.Println("requestbody is prime number")
	} else {
		fmt.Println("requestbody is not prime number")
	}
}

func Prime(n int) bool {
	if n <= 1 { // Check for numbers less than or equal to 1
		return false
	}
	if n == 2 { //Check if the number is 2
		return true
	}
	if n%2 == 0 { //Check if the number is even
		return false
	}
	squaredata := int(math.Sqrt(float64(n)))
	fmt.Println("squaredata--->", squaredata)
	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			return false
		}
	} //Loop to Check for Divisors
	return true
}
