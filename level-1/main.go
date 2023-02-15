package main

import "fmt"

// package level scope
var x int = 42

var y string = "James Bond"

var z bool = true

// short declaration operator ":=" cant be used in package scope, just in block scope

// user def type
type myType int // underlying type->int

var y2 int

func main() {
	// exercise 1 --- fmt pkg
	/*
		x := 42
		y := "James Bond"
		z := true

		fmt.Printf("%d\t%s\t%t\n", x, y, z)
		fmt.Println(x)
		fmt.Println(y)
		fmt.Println(z)
	*/
	//========================================================================

	// exercise 2 --- package level scope, zro values, var keyword
	/*
		fmt.Println(x, y, z) // --- will print default compiler assigned value aka "zero values"
	*/

	//========================================================================

	// exercise 3 --- fmt.Sprintf
	/*
		s := fmt.Sprintf("%d\t%s\t%t", x, y, z)
		// or ... s := fmt.Sprintf("%v \t %v \t %v", x, y, z)
		fmt.Println(s)
	*/

	//========================================================================

	// exercise 4 --- user def. types, %T formatting, underlying type

	var x myType

	fmt.Println(x)        //print value of x
	fmt.Printf("%T\n", x) // prints the data type of var x

	x = 42

	fmt.Println(x)
	fmt.Printf("%T\n", x) // type remains the same
	// x = "hi" ... gives error because user def type "myType has underlying type "int"

	//=========================================================================

	// exercise 5 ---
	y2 = int(x) // not using := bcz it will "re-declare" the variable
	fmt.Println(y2)
	fmt.Printf("%T\t%T\n", x, y2) //print var type of x and y2
}
