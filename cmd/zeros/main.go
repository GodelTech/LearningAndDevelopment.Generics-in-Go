// package main contains example of work with zero values with generic T type
package main

import "fmt"

// returnZero demonstrates how nil-values work with generic methods and types
func returnZero[T any](s ...T) T {
	var zero T
	return zero
}

func main() {
	fmt.Println(returnZero(5)) // prints "0"

	fmt.Println(returnZero("string")) // prints ""

	fmt.Println(returnZero(true)) // prints "false"

}
