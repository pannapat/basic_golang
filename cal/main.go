package main

import "fmt"

var (
	packageScopeX int  = 5
	packageScopeY int  = 7
	checked       bool = true
	unchecked          = false // type inference

	GlobalScopeZ int // capitalized name
)

func main() {
	// fmt.Println(split(17))

	var x int // function-scope
	var y int // function-scope
	// or var x, y int
	// or
	// var (
	// 	x int
	// 	y int
	// )
	fmt.Scanln(&x)
	fmt.Scanln(&y)
	fmt.Println(myFibonacci(x, y))
	fmt.Println(myFibonacci(myFibonacci(myFibonacci(x, y))))
}
