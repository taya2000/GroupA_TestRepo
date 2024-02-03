package main

import "fmt"

func main() {

	fmt.Println("====Starting the function done by Samhitha which will Print Fibonacci Series====")

	//Fibonacci Series
	// Set the limit for the Fibonacci series
	limit := 50
	// Generate and print the Fibonacci series
	result := fibonacci(limit)
	fmt.Printf("Fibonacci series up to %d: %v\n", limit, result)
	fmt.Println("")

	fmt.Println("====Starting the function done by Pamodi which will calculate the Area and Perimeter of a Rectangle====")
	// Area and Perimeter of Rectangle
	// Declare the length and width variables
	var length, width float64
	// Prompt user to enter the value of length and scan it
	fmt.Println("Enter the length of the rectangle: ")
	fmt.Scanln(&length)
	// Prompt user to enter the value of width and scan it
	fmt.Println("Enter the width of the rectangle: ")
	fmt.Scanln(&width)
	// Print user entered values
	fmt.Println("Your length is: ", length)
	fmt.Println("Your width is: ", width)
	// Call function to calculate area and perimeter
	area, perimeter := calcRectangleValues(length, width)
	// Print the values returned from the function
	fmt.Println("Area of the rectangle is: ", area)
	fmt.Println("perimeter of the rectangle is: ", perimeter)
	fmt.Println("")

	fmt.Println("====Starting the function done by Tejaswi which will calculate the factorial of a number====")
	var num int
	fmt.Println("Enter a number:")
	fmt.Scanln(&num)

	fact := factorial(num)
	fmt.Println("Factorial of", num, "is", fact)
	fmt.Println("")

	//fmt.println("====Starting the function done by Abhisheik Yadla which will print Hello world====")

	var year int

	fmt.Print("Enter a year: ")
	fmt.Scan(&year)

	if isLeapYear(year) {
		fmt.Println(year, "is a leap year.")
	} else {
		fmt.Println(year, "is not a leap year.")
	}

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

// Created by Balangoda Pamodi - 500229522
// This function will accept length and width of a rectangle and return it's area and perimeter
func calcRectangleValues(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2

	return area, perimeter
}

//Created by Tejaswi Cheripally - 500229934
//This function will calculate the factorial of a number

func factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

//Created by Abhisheik Yadla - 500219580
//This function will check whether the given year is leap year or not
func isLeapYear(year int) bool {
	// Leap year conditions: divisible by 4, not divisible by 100 unless divisible by 400
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)

}
