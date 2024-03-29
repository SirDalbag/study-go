package main

import (
	"errors"
	"fmt"
)

func _if() {
	if true {
		fmt.Println("True")
	}
}

func if_else() {
	if true {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}

func if_else_if(number int) {
	if number == 0 {
		fmt.Println("Zero")
	} else if number > 0 {
		fmt.Println("Positive")
	} else {
		fmt.Println("Negative")
	}
}

func swtich_case(status bool) {
	switch status {
	case true:
		fmt.Println("True")
	case false:
		fmt.Println("False")
	default:
		fmt.Println("Default")
	}
}

func raise_error(number int) {
	if number >= 20 {
		fmt.Println(true, nil)
	} else {
		fmt.Println(false, errors.New("Error"))
	}
}

func conditions() {
	_if()
	if_else()
	if_else_if(1)
	swtich_case(true)
	raise_error(20)
}
