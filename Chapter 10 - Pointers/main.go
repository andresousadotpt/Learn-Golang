package main

import "fmt"

/*

x := 5
y := x -- We say that y copies the value of x but lives in another address in RAM


A pointer is a variable that stores a memory address

z := &x -- z stores a reference to x

We create a new variable which lives in another address in ram
and instead of saving a copy of x we save a reference in memory to x
and z will be a pointer type

If we want to de-reference to get the value of z we can do


(we use the de-reference operator `*`)
w := *z;


A method iwth a pointer receivers can modify the value to which the receiver points.

*/

func main() {
	x := 5
	z := &x
	w := *z
	fmt.Println(z)
	fmt.Println(w)
	w = 1
	fmt.Println(z)
	fmt.Println(*z)
	fmt.Println(w)
	*z = 10 // we have to use the * because z is a pointer type which needs to be de-referenced
	fmt.Println(*z)

}
