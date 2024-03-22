package main

import (
	"fmt"
)

func main() {
	calc := InitializeCalculator()

	fmt.Println(calc.Add(4, 8))
}
