package main

import "fmt"

type luckyNumber int

var x luckyNumber = 7

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 3
	fmt.Println(x)
}
