package main

import (
	"fmt"
)

func main() {
	fmt.Println(hello("world"))
}

const engHelloPrefix = "Hello, "

func hello(name string) string {
	return engHelloPrefix + name
}
