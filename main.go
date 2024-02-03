package main

import (
	"fmt"
	"strings"
)

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

	fmt.Println("====Starting the function done by Kamalpreet which will Print Sum of n numbers====")
	var numbr int //Declared an integer variable represent n numbers
	fmt.Print("Enter a number upto you want the sum:")
	fmt.Scan(&numbr) //Take input from the user.
	//Using Printf to print the values in a format.
	fmt.Printf("The total of first %d number is %d", numbr, sum(numbr))

	fmt.Println("====Starting the function done by Rohit which check a string palindrome====")
	var str string //Declaring variable with data type string
	fmt.Print("Enter a string:")
	fmt.Scan(&str)       //Taking user input.
	if palindrome(str) { //Final comparions to print.
		fmt.Printf("%s is a palindrome", str)
	} else {
		fmt.Printf("%s is not a palindrome", str)
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

// Created by Abhisheik Yadla - 500219580
// This function will check whether the given year is leap year or not
func isLeapYear(year int) bool {
	// Leap year conditions: divisible by 4, not divisible by 100 unless divisible by 400
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)

}

//Created by Kamalpreet Kaur - 500218943
//This function print the sum of first n numbers.

func sum(digit int) int { //function with return type integer and 1 integer parameter
	var total int
	total = 0
	for i := 1; i <= digit; i++ { //This loop retrieve all values upto entered digit
		total = total + i // Adding all retrieved values
	}
	return total //Return sum of all numbers.
}

//Created by Rohit - 500230041
//This function will check whether provided string is a palindrome or not.

func palindrome(palin string) bool { //Function with string as parameeter and boolean data type.
	palin = strings.ToLower(palin) // string.ToLower() func will lower the case of given string.

	for i, j := 0, len(palin)-1; i < j; i, j = i+1, j-1 {
		if palin[i] != palin[j] { //Comparing the main string and reversed string.
			return false //Will return false if the reversed string don't match.
		}
	}
	return true

}
