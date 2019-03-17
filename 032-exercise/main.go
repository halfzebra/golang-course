package main

import "fmt"

type luckyNumber int

var x luckyNumber = 7

var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	y = int(x)
	fmt.Println(x)
}
