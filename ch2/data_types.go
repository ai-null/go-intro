package main

import (
	"fmt"
)

func main() {
	fmt.Println("\n Rune and strings")
	var x rune = 'A'  // good
	var y int32 = 'J' // legal but confusing
	var z = x + y
	fmt.Printf("string convention %d\n", z) // string convention 139%

	// Explicit Type Conversion
	floatAndInt()
	byteAndInt()

	var i int = 20
	var f float64 = float64(i)
	fmt.Printf("Solution 1, %f\n", f)

	const value = 20
	i = value
	f = value
	fmt.Printf("Solution 2, i=%d\n", i)
	fmt.Printf("Solution 2, f=%f\n", f)

	// assign max values to each type then add 1
	var b byte = 255
	var small int32 = 2147483647
	var bigI uint64 = 18446744073709551615

	b += 1
	fmt.Printf("\nSolution 3, max value byte + 1 = %d", b)

	small += 1
	fmt.Printf("\nSolution 3, max value int32 + 1 = %d", small)

	bigI += 1
	fmt.Printf("\nSolution 3, max value uint64 + 1 = %d", bigI)
}

func floatAndInt() {
	fmt.Println("\n Explicit Type Conversion")
	var x int = 10
	var y float64 = 30.2
	var sum1 float64 = float64(x) + y
	var sum2 int = int(y) + x

	fmt.Println(sum1, sum2)
}

func byteAndInt() {
	var x int = 10
	var y byte = 100
	var sum3 int = x + int(y)
	var sum4 byte = byte(x) + y

	fmt.Println(sum3, sum4)
}
