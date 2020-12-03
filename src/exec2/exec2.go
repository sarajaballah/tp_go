//Create more sophisticated but still simple programs using Go and data structure usage.
package main

import (
	"fmt"
	"unicode"
	"strings"
)

func main() {
	fmt.Println("Hello, world!")

	fmt.Println("ParsePhone() test")
	fmt.Printf("ParsePhone(%q) = %q\n", "123-456-7890", ParsePhone("123-456-7890"))
	fmt.Printf("ParsePhone(%q) = %q\n", "1 2 3 4 5 6 7 8 9 0", ParsePhone("1 2 3 4 5 6 7 8 9 0"))

	fmt.Println("Anagram() test")
	fmt.Printf("Anagram(%q, %q) = %v\n", "12345", "52314", Anagram("12345", "52314"))
	fmt.Printf("Anagram(%q, %q) = %v\n", "21435", "53241", Anagram("21435", "53241"))
	fmt.Printf("Anagram(%q, %q) = %v\n", "12346", "52314", Anagram("12346", "52314"))
	fmt.Printf("Anagram(%q, %q) = %v\n", "123456", "52314", Anagram("123456", "52314"))

	fmt.Println("FindEvens() test")
	fmt.Printf("FindEvens(%v) = %v\n", []int{1, 2, 3, 4}, FindEvens([]int{1, 2, 3, 4}))

	fmt.Println("SliceProduct() test")
	fmt.Printf("SliceProduct(%v) = %v\n", []int{5, 6, 8}, SliceProduct([]int{5, 6, 8}))

	fmt.Println("Unique() test")
	fmt.Printf("Unique(%v) = %v\n", []int{1, 2, 3, 4, 4, 5, 6, 6}, Unique([]int{1, 2, 3, 4, 4, 5, 6, 6}))

	fmt.Println("InvertMap() test")
	fmt.Printf("InvertMap(%v) = %v\n", map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}, InvertMap(map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}))
}

// ParsePhone parses a string of numbers into the format 06 22 14 33 44.
// This function should handle any number of extraneous spaces and dashes.
// All inputs will have 10 numbers and maybe extra spaces and dashes.
// For example, ParsePhone("123-456-7890") => "12 34 56 78 90"
//              ParsePhone("1 2 3 4 5 6 7 8 9 0") => "12 34 56 78 90"
func ParsePhone(phone string) string {
	// TODO
	var num string
	var c int
	c = 1
	for _, char := range phone {
		if unicode.IsDigit(char) {
			num = num + string(char)
			if len(num) == (3*c - 1) {
				num = num + " "
				c++
			}
		}
	}
	return num
}

// Write a function to check whether two given strings are anagram of each other or not.
// An anagram of a string is another string that contains same characters,
// only the order of characters can be different. For example, “abcd” and “dabc” are anagram of each other.
// This function is NOT case sensitive and should handle UTF-8
func Anagram(s1, s2 string) bool {
	// TODO
	for _, char := range s1 {
		if !strings.Contains(s2,string(char)) {
			return false
		}
	}
	return true
}

// FindEvens filters out all odd numbers from input slice.
// Result should retain the same ordering as the input.
func FindEvens(e []int) []int {
	// TODO
	var even []int
	for _,val := range e{
		if val%2 == 0 {
			even = append(even, val)
		}
	}
	return even
}

// SliceProduct returns the product of all elements in the slice.
// For example, SliceProduct([]int{1, 2, 3}) => 6
func SliceProduct(e []int) int {
	// TODO
	var result int
	result = 1
	for _,val := range e{
		result = result * val
	}
	return result
}

// Unique finds all distinct elements in the input array.
// Result should retain the same ordering as the input.
func Unique(e []int) []int {
	// TODO
	// var result []int

	// for _,val := range e{
	// 	if 
	// }
	return nil
}

// InvertMap inverts a mapping of strings to ints into a mapping of ints to strings.
// Each value should become a key, and the original key will become the corresponding value.
// For this function, you can assume each value is unique.
func InvertMap(kv map[string]int) map[int]string {
	// TODO
	new_map := make(map[int]string)
	for key,value := range kv {
		new_map[value]=key
	}
	return new_map
}
