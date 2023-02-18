// pointers
package main

import (
	"fmt"
)

func foo(y *int) {
	fmt.Println("y befor", y)
	fmt.Println("y befor", *y)
	*y = 43
	fmt.Println("y after", y)
	fmt.Println("y after", *y)
}

type person struct {
	first string
	last  string
	age   int
}

func (p *person) changeMe() {
	// just using *p.<something> will give ERROR, acceptable shorthand is p.<something>
	(*p).first = "forgor"
	(*p).last = "forgor"
	(*p).age = -90
}

func (p *person) getPersonDetails() {
	fmt.Printf("I am %v %v, aged %v\n", (*p).first, (*p).last, (*p).age)
}

func main() {
	//--------------------------------
	a := 42
	fmt.Println(a)
	fmt.Println(&a) // & gives you the address

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	b := &a
	fmt.Println(b)
	fmt.Println(*b) // * gives you the value stored at an address when you have the address
	fmt.Println(*&a)

	*b = 43
	fmt.Println(a)
	//----------------------------------
	x := 0
	fmt.Println("x befor", &x)
	fmt.Println("x befor", x)
	foo(&x)
	fmt.Println("x after", &x)
	fmt.Println("x after", x)
	//------------------METHOD SETS--------------------
	// !!! refer notes

	/*
		● a NON-POINTER RECEIVER
			○ works with values that are POINTERS or NON-POINTERS.
		● a POINTER RECEIVER
			○ only works with values that are POINTERS.

		Receivers 	|		Values
		------------+----------------------------------
		(t T) 		|		T and *T
		(t *T) 		|		*T


		IMPORTANT: “The method set of a type determines the INTERFACES
					that the type implements.....”
															~ golang spec
	*/

	//--------------------------------------------------

	// exercise - 1------------------------------------
	val := 's'
	vari := val
	fmt.Println(&val, &vari) // getting addresses of both the variable...0xc000010250 0xc000010260

	// exercise - 2 -----------------------------------
	per := person{
		first: "capitane",
		last:  "spok",
		age:   66,
	}
	per.getPersonDetails()
	per.changeMe()
	per.getPersonDetails()
}
