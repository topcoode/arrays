package operations

import (
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Inputdata struct {
	Data []int `json:"data"`
}

func JumpSearch(c *gin.Context) {
	var In Inputdata
	err := c.ShouldBindJSON(&In)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	target := 5
	JumpSearch_main(In.Data, target)
}
func JumpSearch_main(In []int, target int) int {

	length := len(In)
	var blockSize = int(math.Sqrt(float64(length)))

	var i = 0

	for {

		if In[i] >= target {
			break
		}

		if i > len(In) {
			break
		}

		i += blockSize
	}

	for j := i; j > 0; j-- {
		if In[j] == target {
			return j
		}
	}

	return -1
	/*
		target := 7
		arr := In.Data
		n := len(arr)
		if n == 0 {
			c.JSON(http.StatusOK, gin.H{"result": "Array is empty"})
			return
		}

		// Finding the block size to be jumped
		step := int(math.Sqrt(float64(n)))

		// Finding the block where the element is present
		prev := 0
		for arr[int(math.Min(float64(step), float64(n)))-1] < target {
			prev = step
			step += int(math.Sqrt(float64(n)))
			if prev >= n {
				c.JSON(http.StatusOK, gin.H{"result": "Element not found"})
				return
			}
		}

		// Doing a linear search for target in the block beginning with prev
		for arr[prev] < target {
			prev++
			// If we reached next block or end of array, element is not present
			if prev == int(math.Min(float64(step), float64(n))) {
				c.JSON(http.StatusOK, gin.H{"result": "Element not found"})
				return
			}
		}

		// If element is found
		if arr[prev] == target {
			c.JSON(http.StatusOK, gin.H{"result": fmt.Sprintf("Element found at index %d", prev)})
			return
		}

		c.JSON(http.StatusOK, gin.H{"result": "Element not found"})
	*/
}
