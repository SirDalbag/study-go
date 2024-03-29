package main

import "fmt"

var a int = 2
var b int = 4

var x bool = true
var y bool = false

func arithmetic() {
	fmt.Println("Addition:", a+b)

	fmt.Println("Subtraction:", a-b)

	fmt.Println("Multiplication:", a*b)

	fmt.Println("Division:", a/b)

	fmt.Println("Residue from division:", a%b)
}

func comparison() {
	fmt.Println("Equal:", a == b)

	fmt.Println("Not equal:", a != b)

	fmt.Println("Greater than:", a > b)

	fmt.Println("Less than:", a < b)

	fmt.Println("Greater than or equal to:", a >= b)

	fmt.Println("Less than or equal to:", a <= b)
}

func logical() {
	fmt.Println("AND:", x && y)

	fmt.Println("OR:", x || y)

	fmt.Println("NOT:", !x)
}

func operators() {
	arithmetic()
	comparison()
	logical()
}
