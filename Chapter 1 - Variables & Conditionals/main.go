package main

import "fmt"

/*
Basic data types

bool

string

int  int8  int16  int32  int64  (pos., neg.) the number after is the number of bits of the variable
uint uint8 uint16 uint32 uint64 uintptr (no neg.) the number after is the number of bits of the variable

in case of int and uint, either 32 or 64 bits, depends on the system arch.
uintptr is an integer type that is large enough to contain the bit pattern of any pointer

byte // under the hood is an alias for uint8

rune // under the hood is an alias for int32, repesents a Unicode code point

float32 float64

complex64 complex128 // represent the concept of imaginary numbers

*/

func main() {
	// Different ways to initialize variables:

	// with :=
	a := 1
	fal := false
	hey, world := "Hello", "World"

	// with var
	var b int = 1

	// if we will have a lot of variables we can do like this for better syntax and readability
	var (
		c int     = 2
		d float32 = 3.3
		f bool    // default is false
		g int     // default is 0
		h float32 // default is 0
	)

	// Some type convertion
	intValue := 88
	floatValue := float32(88)

	// Const aka immutable variables (we can't modify it but can compute constants), in GO we don't use like const VALUE_OF_SOMETHING
	const constVariable = "constant variable"

	// Formatting strings
	// printf - prints a formatted string to stdout
	// sprintf - returns the formattes tring
	fmt.Printf("I am %v\n", a) // %v is representation of value. Use this if we are unsure what else to use. Normally just the type-specific variant
	// we have %s for string or %d for integer in decimal form or %f for a decimal (%.1f prints the first number after decimal point)

	// Conditionals is the same as java or c# but we can do a magic thing which is:
	// Initial Statement of an if block
	/*
		length := getLength(email)
		if length < 1 {
			fmt.Println("Email is invalid")
		}
	*/
	// We can do a syntatic sugar which is:
	// if INITIAL_STATEMENT; CONDITION {}
	// But the initialized variable is only accessable inside the if block
	// is a safety hack to ensure that the test variable is never reused in other code if we do not intend to
	if test := 1; test > 0 {
		fmt.Printf("%v is higher than 0\n", test)
	}

	fmt.Println(hey, world)
	fmt.Println(a, fal, b, c, d, f, g, h)
	fmt.Printf("Type of a -> %T\n", a) // %T is the type of something
	fmt.Println(intValue, floatValue)
	fmt.Println(constVariable)
}
