package main

import "fmt"

func main () {
	
	fmt.Println("====Starting the function done by Samhitha which will Print Fibonacci Series====")

	//Fibonacci Series
	// Set the limit for the Fibonacci series
	limit := 50
	// Generate and print the Fibonacci series
	result := fibonacci(limit)
	fmt.Printf("Fibonacci series up to %d: %v\n", limit, result)
	fmt.Println("")
}

// Created by Samhitha Dubbaka - 500225971
// fibonacci generates a Fibonacci series up to the given limit
func fibonacci(limit int) []int {
	// Initialize the series with the first two Fibonacci numbers
	series := []int{0, 1}
	// Generate Fibonacci numbers until reaching the limit
	for i := 2; ; i++ {
		nextNum := series[i-1] + series[i-2]
		// Break if the next number exceeds the limit
		if nextNum > limit {
			break
		}
		series = append(series, nextNum)
	}
	return series
}