package main

/*
Interfaces

Interfaces are collections of method signatures. A type "implements" an interface if
	it has all of the methods of the given interface defined on it.

If ANY type that implements all of the methods inside the interface we can say that, that type implement the interface.
And we can treat all of that types as the interface, easier showing example

type shape interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

func (r rect) area() float64{
	return r.width * r.height
}

func (r rect) perimeter() float64{
	return 2*r.width + 2*r.height
}

We can say that `rect` implements all of the methods, so we can treat `rect` as a `shape`

A powerfull things its if a func takes an interface, we can use all of the types that implement that interface for example

func printShapeWidth(s shape){
	fmt.Println(s.width)
}

Because that rect is an implementation of the shape interface
we can use it like this:

printShapeWidth(rect{
	width: 1,
	height: 2,
})

 which returns "1"


We can have multiple interfaces implemented
A type can implement any number of interfaces in Go

We can name our interface arguments and return types
so we guarantee that the type that is implementing an interface has the correct arguments so it can be viewed or treated as an implementation of that interface

type Copier interface {
	Copy(string, string) int
}

BTW: Normally in interfaces we use the "er" in the final...


Type assertions

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

// "c"  is a new circle cast from "s"
// which is an instance of a shape
// "ok" is a bool that is true if s was a circle
// or false if s isn't a circle

c, ok := s.(circle) // we can cast an interface to a specific type




Interface type switches:

func printNumericValue(num interface{}){
	// we are casting the num interface into a "type" which we will do a pattern match
	// but we can receive from parameters like `s shape` for example, this is a boilerplate code so we can change either the interface{} or the "type"
	switch v := num.(type){
		case int:
			fmt.Printf("%v\n", v)
		case string:
			fmt.Printf("%v\n", v)
		default:
			fmt.Printf("%v\n", v)
	}
}


Interfaces should have a "minmal necessary behavior we need" so we should have as frew methods as possible

*/

func main() {

}
