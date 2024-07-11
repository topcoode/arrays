package main

import (
	op "things/Operations"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/Factorial_Calculation", op.Factorial_Calculation)
	r.GET("/Fibonacci_Series", op.Fibonacci_Series)
	r.POST("/Prime_Number", op.Prime_Number)
	r.POST("/Palindrome", op.Palindrome)
	r.GET("/Add", op.Add)
	r.POST("/StringReverse", op.StringReverse)
	r.POST("/BubbleSort", op.BubbleSort)
	r.GET("/loops", op.Loops)
	r.GET("/Metrics", op.Matrics)
	r.GET("/3d-dimensional", op.ThreeD_operations)
	r.GET("/AddingSubtration", op.Adding_Subtration)
	r.POST("/linearsearch", op.Linear_Search)
	r.POST("/BinaryOperations", op.BinarySearch)
	r.GET("/acenddecen", op.Ascending_Decending)
	r.POST("/selectionsort", op.SelectionSort)
	r.POST("/subarray", op.Sub_array)
	r.POST("/subsequence", op.Subsequence)
	r.POST("/mergesort", op.Merge_sort)
	r.GET("/averagesubarray", op.AverageSubarray)
	r.GET("/Naive_Approach", op.Naive_Approach)
	r.GET("/Sliding_Window_technique", op.Sliding_Window_technique)
	r.GET("/Trailing_Zero", op.Trailing_Zero)
	r.GET("/CoutingPlaceValue", op.CoutingPlaceValue)
	r.GET("/Recursion_main", op.Recursion_main)
	r.GET("/Back_Tracking", op.Back_Tracking)
	r.Run(":8080")
}
