package operations

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Sliding_Window_technique(c *gin.Context) {
	arr := []int{45, 85, 65, 35, 75, 96, 72, 52}

	wSum := 0
	mSum := 0
	k := 3
	for i := 0; i < k; i++ {
		wSum = wSum + arr[i]
		fmt.Println(wSum)

	}
	mSum = wSum
	fmt.Println("wSum,mSum \n", wSum, mSum)
	for i := k; i < len(arr); i++ {
		//end := i:=0 ,k:=3 end:= 0-3 = 3
		//start := a[i] = 0
		// windowsum := windowsum + arr[i-k] + a[i]
		//windowSum - arr[i] + arr[i+windowSize]
		// windowsum := previouswindow - array[starting - pointer]+array[i]
		wSum := wSum - arr[i-k] + arr[i]
		fmt.Println("wSum--->", wSum)
		if wSum > mSum {
			fmt.Println("value is greater than the msum")
			mSum = 0
			mSum = wSum
		} else {
			fmt.Println("is less than the value")
		}

	}
	fmt.Println("mSum--->", mSum)
}
