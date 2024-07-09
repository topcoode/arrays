package operations

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Naive_Approach(c *gin.Context) {
	arr := []int{75, 86, 84, 5, 74, 1, 86, 5}
	mSum := 0
	cSum := 0
	k := 3
	//Boundary
	//l:=10-3 - third last element
	//time- complexity - O(n^2)
	for i := 0; i < len(arr)-k; i++ {
		cSum = 0
		for j := i; j < k; j++ {

			cSum = cSum + arr[j]
		}
		if cSum > mSum {
			fmt.Println("value is greater than the msum")
			mSum = 0
			mSum = cSum
		} else {
			fmt.Println("is less than the value")
		}
	}
	fmt.Println(mSum)
}
