package main

import "fmt"

func main() {
	fv := func() {
		fmt.Printf("Hello World")
	}
	fv()
	fmt.Printf("fv is of type %T and has value %v", fv, fv)
}
