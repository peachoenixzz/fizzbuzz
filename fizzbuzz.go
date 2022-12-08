package fizzbuzz

import "fmt"

func isFizzBuzz(n int) bool {
	if n%3 == 0 && n%5 == 0 {
		return true
	}
	return false
}

func FizzBuzz(n int) string {
	if isFizzBuzz(n) {
		return "FizzBuzz"
	}

	if n%3 == 0 {
		return "Fizz"
	}

	if n%5 == 0 {
		return "Buzz"
	}

	return fmt.Sprintf("%v", n)
}
