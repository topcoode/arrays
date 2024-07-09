package operations

import (
	"fmt"
	"sort"

	"github.com/gin-gonic/gin"
)

type ValueIndex struct {
	Value int
	Index int
}

func TwoSum(c *gin.Context) {
	arr := []int{3, 2, 4}
	target := 6
	output := TwoSumMain(arr, target)
	fmt.Println(output)
	c.JSON(200, gin.H{
		"result": output,
	})
}

func TwoSumMain(arr []int, target int) []int {
	valueIndexArr := make([]ValueIndex, len(arr))
	for i, v := range arr {
		valueIndexArr[i] = ValueIndex{Value: v, Index: i}
	}

	sort.Slice(valueIndexArr, func(i, j int) bool {
		return valueIndexArr[i].Value < valueIndexArr[j].Value
	})

	start := 0
	end := len(valueIndexArr) - 1
	result := []int{-1, -1}

	for start < end {
		sum := valueIndexArr[start].Value + valueIndexArr[end].Value
		if sum == target {
			result[0] = valueIndexArr[start].Index
			result[1] = valueIndexArr[end].Index
			return result
		} else if sum < target {
			start++
		} else {
			end--
		}
	}
	return result
}
