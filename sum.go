package main

import "fmt"

func main() {
	fmt.Println(sum(2, 2))
}

func sum(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func multiply(a int, b int) int {
	return a * b
}

func divide(a int, b int) int {
	if b == 0 {
		return 0 // Avoid division by zero
	}
	return a / b
}

func mod(a int, b int) int {
	if b == 0 {
		return 0 // Avoid division by zero
	}
	return a % b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
