package main

/*

========== Arrays: ==========

[n]T is an array of n values of type T
arrays are fixed in size

myInts := [10]int
this is array of 10 zeros


========== Slices: ==========

[]T is a slice of T

99 out of 100 times we use slices because its better :EZ:
Slices are dynamically sized flexible view into an array. Slices are built on top of arrays

Lets imagine we have:

a := [5]int {6,3,2,6,5}
sliceOfA := a[1:4] // we wanted the 3,2,6

if we want all of the values we use:
arrayname[:]

Syntax we can use:
arrayname[lowIndex:highIndex] -- everything between lowIndex and highIndex
arrayname[lowIndex:] -- everything after lowIndex
arrayname[:highIndex] -- everything before highIndex
arrayname[:] -- all


========== make: ==========
The make is a built-in function which allocates and initializes an object of type slice, map, or chan.

Most of the time we don't need to think about the underlying array of a slice. We can create a new slice, map, or chan using the make function
This is how you create dynamically-sized arrays

In case of a slice, the size specifies its length or capacity
len -> length
cap -> capacity
 func make([]T, len, cap) []T
mySlice := make([]int, 5, 10)

we can ommit the cap and it will default to the length


we have built-in functions:
len([]T) length of a slice
cap([]T) capacity of a slice

In case of a map, an empty map is allocated with enough space to hold the specified
 number of elements. We can omit the size parameter, in which case a small starting size is allocated.

In case of a channel, the channel's buffer is initialized with the specified buffer capacity.
 If zero, or the size is omitted, the channel is unbuffered.


========== Variadic: ==========

Variadic functions can be called with any number of trailing arguments.

We use "..." syntax ins the function signature

variadic functions receives the variadic arguments as a slice

func sum(nums ...int) int {
	for i := 0; i < len(nums); i++ {
		num := nums[i]
	}
}
func main() {
	total := sum(1,2,3)
	fmt.Println(total)
}

Spread operator:
Is an operator that allows us to pass a slice into a variadic function

func main() {
	num := []int{1,2,3}
	fmt.Println(sum(nums...))
}


Built-in functions

append(slice []T, elems ...T) []T



========== Range: ==========
syntactic sugar to iterate easily over elements of a slice

for INDEX, ELEMENT := range SLICE {

}


*/

func main() {

}
