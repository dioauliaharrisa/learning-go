package main

import (
	"fmt" 
	"strconv"
)

// you can declare on package level, cannot use short hand here
var i0 int = 11

// this variable is set to be global because it is uppercased
var I6 int = 86

// you can declare block of variables
var (
	i5name string = "i5"
	i5pronunciation string = "aifaif"
)

func main() {
	// printing hello world
	fmt.Println("Hello Dio!")

	fmt.Println("i0 is", i0)
	// declaring a variable
	var i1 int
	i1 = 42
	i1 = 27
	fmt.Println("i1 is",i1)

	// declaring and assigning can be one line
	var i2 int
	i2 = 42
	fmt.Println("i2 is",i2)

	// short hand for integer i guess?
	i3 := 42
	fmt.Println("i3 is",i3)

	// printing the type and value
	// := can only be inferred as int or float64 type
	i4 := 20
	fmt.Printf("i4 is %v, the type is %T", i4, i4)

	// shadowing
	// when variable is redeclared, the inner scope takes place>>
	fmt.Println("i0 is",i0)
	var i0 int = 12
	fmt.Println("i0 is",i0)

	// to convert int into string
	var i5 int = 42
	fmt.Printf("%v %T", i5, i5)
	var j5 string
	j5 = strconv.Itoa(i5)
	fmt.Printf("%v %T", j5, j5)

}