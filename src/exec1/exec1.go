//become familiar writing Go programs and running them on your machine

package main

import "fmt"

func main() {
	// use the main function for testing your functions
	fmt.Println("Hello, world!")

	fmt.Println("Fizzbuzz() test")
	fmt.Printf("Fizzbuzz(%v) = %v\n", 27, Fizzbuzz(27))
	fmt.Printf("Fizzbuzz(%v) = %v\n", 25, Fizzbuzz(25))
	fmt.Printf("Fizzbuzz(%v) = %v\n", 105, Fizzbuzz(105))

	fmt.Println("IsPrime() test")
	fmt.Printf("IsPrime(%v) = %v\n", 1, IsPrime(1))
	fmt.Printf("IsPrime(%v) = %v\n", 2, IsPrime(2))
	fmt.Printf("IsPrime(%v) = %v\n", 22, IsPrime(22))
	fmt.Printf("IsPrime(%v) = %v\n", 32, IsPrime(32))

	fmt.Println("IsPalindrome() test")
	fmt.Printf("IsPalindrome(%v) = %v\n", "abccba", IsPalindrome("abccba"))
	fmt.Printf("IsPalindrome(%v) = %v\n", "abccbb", IsPalindrome("abccbb"))
	fmt.Printf("IsPalindrome(%v) = %v\n", "abcbba", IsPalindrome("abcbba"))
}

// Fizzbuzz is a classic introductory programming problem.
// If n is divisible by 3, return "Fizz"
// If n is divisible by 5, return "Buzz"
// If n is divisible by 3 and 5, return "FizzBuzz"
// Otherwise, return the empty string
func Fizzbuzz(n int) string {
	// TODO
	var ch string
	if n%3 == 0 {
		ch = "Fizz"
	}
	if n%5 == 0 {
		ch = ch + "Buzz"
	}
	return ch
}

// IsPrime checks if the number is prime.
// You may use any prime algorithm, but you may NOT use the standard library.
func IsPrime(n int) bool {
	// TODO
	var lim int
	lim = (n / 2) - 1
	if n == 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i <= lim; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// IsPalindrome checks if the string is a palindrome.
// A palindrome is a string that reads the same backward as forward.
func IsPalindrome(s string) bool {
	// TODO
	for i := 0; i < len(s); i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
