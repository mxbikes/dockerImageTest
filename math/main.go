package main

import "fmt"

func main() {
	var value1 = 5
	var value2 = 7
	var result = Add(value1, value2)

	fmt.Printf("Result: %v + %v = %v", value1, value2, result)
}
