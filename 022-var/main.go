package main

import "fmt"

// Global scope
var a = 1;

// Variable with zer value.
var c int;

func main() {
	// Local scope variable
	b := 2;
	fmt.Println(a, b, c)
}