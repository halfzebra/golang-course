package main

import "fmt"

// ... - variadic argument
// interface{} - any
func print(a ...interface{}) {
	fmt.Println(a)
}

func main() {
	print(1, 2, 3)
}
