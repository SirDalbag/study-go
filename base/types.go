package main

import "fmt"

var _global any
var _ any

func numbers() {
	var _int int = 0
	var _int8 int8 = 127
	var _int16 int16 = 32767
	var _int32 int32 = 2147483647
	var _int64 int64 = 9223372036854775807

	fmt.Println("Int:", _int, _int8, _int16, _int32, _int64)

	var _uint uint = 0
	var _uint8 uint8 = 255
	var _uint16 uint16 = 65535
	var _uint32 uint32 = 4294967295
	var _uint64 uint64 = 18446744073709551615

	fmt.Println("Uint:", _uint, _uint8, _uint16, _uint32, _uint64)

	var _float32 float32 = 0.1
	var _float64 float64 = 0.1

	fmt.Println("Float:", _float32, _float64)

	var _complex64 complex64 = complex(0, 0)
	var _complex128 complex128 = complex(0, 0)

	fmt.Println("Complex:", _complex64, _complex128)
}

func chars() {
	var _string string = "Hello, World!"

	fmt.Println("String:", _string)

	var _rune rune = 'H'
	fmt.Println("Rune:", _rune)

	var _byte byte = 'H'
	fmt.Println("Byte:", _byte)
}

func booleans() {
	var _bool bool = true
	fmt.Println("Bool:", _bool)
}

func arrays() {
	var _array [5]int = [5]int{1, 2, 3, 4, 5}

	fmt.Println("Array:", _array)

	var _slice []int = []int{1, 2, 3, 4, 5}

	fmt.Println("Slice:", _slice)
}

func maps() {
	var _map map[string]int = map[string]int{"a": 1, "b": 2, "c": 3}

	fmt.Println("Map:", _map)
}

func structs() {
	type Struct struct {
		Field string
	}

	var _user Struct = Struct{Field: "Field"}

	fmt.Println("Struct:", _user)
}

func pointers() {
	var _int int = 0

	var _pointer *int = &_int

	fmt.Println("Pointer:", *_pointer)
}

func types() {
	fmt.Println("Global:", _global)
	numbers()
	chars()
	booleans()
	arrays()
	maps()
	structs()
	pointers()
}
