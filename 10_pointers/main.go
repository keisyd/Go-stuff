package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	//Use * to read val from address
	fmt.Println(a, " ", *b)

	//Change value with the pointer
	*b = 10

	fmt.Println(a)

}
