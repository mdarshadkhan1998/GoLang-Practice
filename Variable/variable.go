package main

import "fmt"

const (
	true  = 0 == 0 // Untyped bool.
	false = 0 != 0 // Untyped bool.
)

func main() {
	var a string
	fmt.Println("a", a)
	var b, c string = "apple", "ball"
	fmt.Println(b, c)
}
