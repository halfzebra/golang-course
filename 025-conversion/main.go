package main

import "fmt"

var a int = 1;

type hotdog int

var b hotdog = 2;

func main() {
	a := 1;
	fmt.Printf("%T %v", a, a)
}