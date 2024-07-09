package operations

import "fmt"

func N_dimensional() {
	// Define the dimensions
	const N = 5
	const size = 3 // Size for each dimension

	// Declare a 5-dimensional array
	var arr [size][size][size][size][size]int

	// Initialize values (optional)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			for k := 0; k < size; k++ {
				for l := 0; l < size; l++ {
					for m := 0; m < size; m++ {
						arr[i][j][k][l][m] = i + j + k + l + m
					}
				}
			}
		}
	}

	// Accessing elements
	fmt.Println("Element at [1][1][1][1][1] is", arr[1][1][1][1][1])
}
