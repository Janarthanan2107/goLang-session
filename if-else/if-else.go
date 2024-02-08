// package name should be declare at first as main
package main

// packages need to import at second
import (
	"fmt"
	"strconv"
)

// main function also a mandatory thing to use
func main() {
	var num int

	// printf ---> like console.log in javascript
	fmt.Printf("Enter a Number: ")

	// scanf ---> prompt message in javascript
	fmt.Scanf("%d", &num)

	// ifelse are same as javascript

	// example : calculating the times of the number given using shorthand method
	// NOTE: declaring the variable like this available for the both if and else & for the whole block
	noOfNum := num / 2

	// calculating the times of the number given on the same ifelse method
	// NOTE: declaring the variable on the  ifelse method is only available for the both if and else only not for the whole block
	if nNum := num / 2; num%2 == 0 {
		// noOfNum is declared outside the condition
		fmt.Printf("%d is Even, %d times,", num, noOfNum)
	} else {
		// nNum is declared on the condition
		fmt.Printf("%d is Odd, %d times,", num, nNum)
	}

	fmt.Printf("%s", fizzBuzz(num))
}

// example program
func fizzBuzz(n int) string {
	if n%3 == 0 && n%5 == 0 {
		return "FizzBuzz"
	} else if n%3 == 0 {
		return "Fizz"
	} else if n%5 == 0 {
		return "Bizz"
	} else {
		return strconv.Itoa(n)
	}
}
