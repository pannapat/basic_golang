package main

import (
	"fmt"
	"hi/echo"
)

func init() {
	fmt.Println("Initial...")
}

func add(x int, y int) int {
	return x + y
}

func minus(x int, y int) int {
	return x - y
}

func multiply(x int, y int) int {
	return x * y
}

func divide(x int, y int) int {
	return x / y
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	echo.Say()
	fmt.Println("add", add(42, 13))
	fmt.Println("minus", minus(42, 13))
	fmt.Println("multiply", multiply(42, 13))
	fmt.Println("divide", divide(42, 12))
	a, b := swap("hello", "world");
	fmt.Println(a,b);
}
