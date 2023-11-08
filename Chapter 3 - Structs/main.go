package main

/*

Struct is a collection type which is a type which contains other types

Struct is a struct with K, V pairs


** Nested Structs

Here in car and Wheel we have a nested structs

*/

type Car struct {
	make   string
	model  string
	weight float32
	wheel  Wheel
}

type Wheel struct {
	radius   int
	material string
}

/*

Anonymous structs


Its like a normal struct but we define it without a name, so we cannot reference elsewhere that struct.
we see it more when we need to nest an anonymous struct as fields whithin other struct:

type car struct {
	name string
	wheel struct {
		radius int
		material string
	}
}

but we should favor using named structs

Embeded structs

type car struct {
	make string
	model string
}

type truck struct {
	"car" is embeded, so the definition of truck
	 also contains all of the fields of the car struct
	car
	bedSize int
}

what this differs from the nested struct

if we had a nested struct and we wanted to access the trucks make
we would need to do like `truck.car.make`
but with embeded structs we only need to do `truck.make`


Struct methods
Methods that are defined as structs.

Methods are just functions that have a receiver. A receiver is a special parameter that syntactically goes before the name of the function

type rect struct {
	width int
	height int
}

func (r rect) area() int {
	return r.width * r.height
}

r := rect {
	width: 5,
	height: 10,
}

fmt.Println(r.area())


*/

func main() {
	myCar := Car{} // here we create a new empty car with all the fields initialized as their default values
	myCar.wheel.radius = 32
}
