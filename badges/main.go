package main

import "fmt"

func main() {
	fmt.Println(add(1, 3))
	fmt.Println(sub(3, 1))
	fmt.Println("updated workflow ok")
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

