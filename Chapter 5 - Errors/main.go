package main

import (
	"fmt"
	"strconv"
)

/*


	When something can go wrong in a function that function
	should return an error as its last return value.

	Errors are just values and specifically they are just interfaces


	Because errors are just interfaces, you can build
	your own custom type that implements the error interface

	type userError struc {
		name string
	}

	func (e userError) Error() string {
		return fmt.Println("HELLO THIS IS CUSTOM ERROR TYPE")
	}


	The go std library provides an errors package. We can do

	var err error = errors.New("Somethin went wrong")

*/

func main() {

	i, err := strconv.Atoi("42b")
	if err != nil {
		fmt.Println("Couldn't convert: ", err)
		return
	}

	fmt.Println(i)

}
