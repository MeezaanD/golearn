package main

import (
	// "fmt"
	// "reflect"
	// gr "hello/greeting"
	// a "hello/arrays"
	// s "hello/slices"
	m "hello/maps"
)

// LESSON 1
//  Variables Declaration

// func main() {

// 	fmt.Println("Hello, Go!")
// 	fmt.Println('A')

// 	var boolean bool = true //typed
// 	// boolean2 := true //untyped

// 	fmt.Println(boolean)
// 	// 42
// 	// 3.14
// 	// +, -, *, /
// 	// > < == != >= <=

// 	//  Gives the variable type

// 	fmt.Println(reflect.TypeOf(3.14))

// 	//  Common way to declare variables

// 	// quantity := 4
// 	// mini, max := 2.4, 3

// 	// -------------------------------------------------------------------------

// }

// LESSON 2
// Conditions and Loops

// func main() {
// 		// Conditions

// 	// if 1 == 1{
// 	// 	fmt.Println("1 equals 1")
// 	// } else if 1 != 2 && "hello" == "hello" {
// 	// 	fmt.Println("Double equals")
// 	// } else if 1 == 2 || "hello" == "hello" {
// 	// 	fmt.Println("One equal")
// 	// } else {
// 	// 	fmt.Println("Nothing equals")
// 	// }

// 	// -------------------------------------------------------------------------

// 	// Loops

// 	// i := 0
// 	// for ; i < 10; {
// 	// 	fmt.Println(i)
// 	// 	i++
// 	// }

// 	// -------------------------------------------------------------------------

// 	// continue

// 	// i := 0
// 	// for i < 10 {
// 	// 	fmt.Println(i)
// 	// 	i++
// 	// 	if i == 3 {
// 	// 		break
// 	// 	} else {
// 	// 		continue
// 	// 		fmt.Println("This won't run")
// 	// 	}
// 	// }
// }

// 	// -------------------------------------------------------------------------

// LESSON 3
// Function declaration

// func Hello() {
// 	fmt.Println("Hello")
// }

// // func Add(x int, y int) int {
// // 	return x + y
// // }

// func Add(x int, y int) (bool, int) {
// 	if x + y != 0 {
// 		return true, x + y
// 	} else {
// 		return false, x + y
// 	}
// }

// func main() {
// 	Hello()
// 	// value := Add(2, 3)
// 	// fmt.Println(value)
// 	boolean, value := Add(2, 3)
// 	fmt.Println(value)
// 	fmt.Println(boolean)
// }

// 	// -------------------------------------------------------------------------

// LESSON 4
// Pointers


// func Add(x *int,  y *int) int {
// 	*x = 0
// 	return *x + *y
// }

// func main(){
// 	//& *
// 	// *int
// 	// value := 1
// 	// fmt.Println(value)
// 	// fmt.Println(&value)

// 	// var ptr *int = &value
// 	// fmt.Println(ptr)
// 	// fmt.Println(*ptr)

// 	val1 := 2
// 	val2 := 4
// 	answer := Add(&val1, &val2)
// 	fmt.Println(answer)
// }

// 	// -------------------------------------------------------------------------


// LESSON 5
// Modules

// func main() {
// 	gr.Greet()
// }

// 	// -------------------------------------------------------------------------

// LESSON 6
// Arrays

// func main() {
// 	a.Learn()
// }

// 	// -------------------------------------------------------------------------

// LESSON 7
// Slices

// func main() {
// 	s.LearnSlices()
// }

// 	// -------------------------------------------------------------------------

// LESSON 8
// Maps

func main() {
	m.LearnMaps()
}
