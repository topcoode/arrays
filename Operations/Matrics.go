package operations

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Matrics(c *gin.Context) {
	//inputs
	//finding No of rows and colomns
	matrix1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	matrix2 := [][]int{
		{9, 8, 7},
		{6, 5, 4},
		{3, 2, 1},
	}
	fmt.Println("rows:", len(matrix1), len(matrix2))
	fmt.Println("coloums:", len(matrix1[0]), len(matrix2[0]))

	// Check if matrices have the same dimensions
	if len(matrix1) != len(matrix2) || len(matrix1[0]) != len(matrix2[0]) {
		fmt.Println("Matrices have different dimensions and cannot be added.")
		return
	}
	//Initialize
	rows := len(matrix1)
	cols := len(matrix1[0])
	addition := make([][]int, rows)
	for i := range addition {
		addition[i] = make([]int, cols)
	}
	subtraction := make([][]int, rows)
	for i := range subtraction {
		subtraction[i] = make([]int, cols)
	}
	fmt.Println(addition)
	// Add matrices
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			addition[i][j] = matrix1[i][j] + matrix2[i][j]

		}
	}
	// Add matrices
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			subtraction[i][j] = matrix1[i][j] - matrix2[i][j]

		}
	}
	// Print result matrix
	fmt.Println("Resultant Matrix:")
	for _, row := range addition {
		fmt.Println(row)
	}
}
