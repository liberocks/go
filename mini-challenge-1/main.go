package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Print("Enter a number: ")
	fmt.Scan(&n)
	fmt.Println("read number", n, "from stdin")

	fizzBuzz(n)
}

func fizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		divisibleByThree := i%3 == 0
		divisibleByFive := i%5 == 0

		if divisibleByThree && divisibleByFive {
			fmt.Println("FizzBuzz")
		} else if divisibleByThree {
			fmt.Println("Fizz")
		} else if divisibleByFive {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
