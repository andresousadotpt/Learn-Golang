package main

/*

In a function we can accept a function as parameter


func add(x, y int) int {
	return x + y
}

func mul(x, y int) int {
	return x * y
}

func aggregate(a, b, c int, arithmetic func(int, int) int) int {
	return arithmetic(arithmetic(a, b), c)
}

func main() {
	fmt.Println(aggregate(2, 2, 2, add))
	fmt.Println(aggregate(2, 2, 2, mul))

}


======= Cyrrying ========

Functions currying is a practice of writing a function that takes a function (or functions) as input, and returns a new function

func add(x, y int) int {
	return x + y
}

func mul(x, y int) int {
	return x * y
}

func selfMath(mathFunc func(int, int) int) func(int) int {
	return func(x int) int {
		return mathFunc(x, x)
	}
}

func main() {
	sumFunc := selfMath(add)
	mulFunc := selfMath(mul)
	fmt.Println(sumFunc(2, 2))
	fmt.Println(mulFunc(2, 2))

}


======== Defer keyword ========

defer keyword is fairly unique feature of go. It allows a function to be executed automatically just before its enclosing function returns
(or in another words allows a function to be executed automatically just before the function returns or the function ends)




*/

func main() {

}
