package operations

import "github.com/gin-gonic/gin"

func Insertion_sort(c *gin.Context) {
	arr := []int{78, 23, 45, 2, 98, 5}
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
}

/*
78,|23,45,2,98,5 ----temp-23-----23 > 78 - flase
23,78|,45,2,98,5 ----temp-45-----45 >78 -false,45 >23 -true
23,45,78|,2,98,5 ----temp-2------2 >78 -false , 2 >45 -false ,2 >23 -false
2,23,45,78,|98,5 ----temp-98-----98>78 -true
2,23,45,78,98,|5 ----temp-5------5>98-false,5>78 -flase,5>45-flase,5>23-false,5>2-true
2,5,23,45,78,98-sorted

*/
