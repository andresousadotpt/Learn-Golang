package main

import "fmt"

func sum(x int, y int) int {
	return x + y
}

func f(x, y int, hello string) int { // we can do this x, y int its just a syntatic sugar that go offers
	return x + y
}

/*

Go supports functions as data, or if we want to call them is a callback

func f(func(int, int) int, int) int f receives a function that takes 2 integers and returns an int, then f receives another int, and returns an int

*/

/*

Passing variables by value

func main(){
	x := 5
	increment(x)
	fmt.Println(x)
}

func increment(z int){
	z++
}

When we pass a variable into increment() `z`, we pass a copy of `z` not the actual value of it in memory

for now we can do the following after we solve this issue with pointers:

func main(){
	x := 5
	x = increment(x)
	fmt.Println(x)
}

func increment(z int) int{
	z++
	return z
}

*/

/*

In GO we must adopt using early returns aka guard clauses

*/

// Ignoring return values
func getPoint() (x int, y int) { // when we want to return a lot of types we must use ()
	return 3, 4
}

func main() {
	var (
		z int = 1
		b int = 2
	)
	z = increment(z)
	fmt.Println(sum(z, b))

	// Lets use the getPoint and ignore the y value, in runtime the compiler literally ignore this _ value
	x, _ := getPoint() // we use _ to "ignore" a return value

	fmt.Println(x)
}

func increment(z int) int {
	z++
	return z
}
