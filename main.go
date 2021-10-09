package main

import "fmt"

const myStringHello string = "Hello, World!"

func hello() string {
	return myStringHello
}

func main() {
	fmt.Println(hello())
}
