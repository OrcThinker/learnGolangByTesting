package main

import (
	"fmt"
)

func main() {
	fmt.Println(hello("world"))
}

const engHelloPrefix = "Hello, "

func hello(name string) string {
	if name == "" {
		name = "World"
	}
	return engHelloPrefix + name
}
