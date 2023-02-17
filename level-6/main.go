package main

import (
	"fmt"
	"math"
)

// exercise -1 -----------------------------------------
func foo() int {
	fmt.Println("In foo")
	return 1
}
func bar() string {
	fmt.Println("In bar")
	return "string return"
}

// exercise - 2 and 3 ---------------------------------------------
func fooVariadic(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
func barVariadic(x []int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

// exercise - 4 -----------------------------------------------------
type person struct {
	first string
	last  string
	age   int
}

func (x person) speak() { // method to struct person type
	fmt.Printf("Hi I am %v %v and i am  %v years old\n", x.first, x.last, x.age)
}

// exercise - 5 ----------------------------------------------
type SQUARE struct {
	L float64
}
type CIRCLE struct {
	R float64
}

func (s SQUARE) AREA() float64 {
	return math.Pow(s.L, 2)
}
func (c CIRCLE) AREA() float64 {
	return math.Pi * math.Pow(c.R, 2)
}

type SHAPE interface {
	AREA() float64
}

func INFO(s SHAPE) {
	fmt.Println(s.AREA())
}

// exercise - 8 ----------------------------------------------
func funcOFunc(s string) func() string { // return type : a function that has a return type of string
	return func() string {
		return s
	}
}

// exercise - 9 ----------------------------------------------
func foo2(f func(xi []int) int, ii []int) int {
	sumHold := f(ii)
	return sumHold
}

// ------------------------------ main ---------------------------
func main() {
	// exercise - 1 ---------------------------------------------
	fooVal := foo()
	barVal := bar()
	fmt.Printf("foo returns...%v\n", fooVal)
	fmt.Printf("bar returns...%v\n", barVal)

	// exercise - 2 and 3 ----------------------------------------
	vals := []int{1, 2, 3}
	fooVal = fooVariadic(vals...) // unfurling the slice
	defer fmt.Printf("foo variadic returns ... %v \n", fooVal)
	barVal2 := barVariadic(vals)
	fmt.Printf("bar variadic returns ... %v \n", barVal2)

	// exercise - 4 ----------------------------------------------
	shlok := person{
		first: "shlok",
		last:  "jha",
		age:   23,
	}
	shlok.speak()

	// exercise - 5 ----------------------------------------------
	s := SQUARE{
		L: 2,
	} // square with side = 2
	c := CIRCLE{
		R: 1,
	} // unit circle
	fmt.Printf("Area of square of side: %v and circle of radius: %v are ... '%v' and '%v' respectively\n", s, c, s.AREA(), c.AREA())
	// calling method to the common interface
	INFO(s)
	INFO(c)

	// exercise - 6 ----------------------------------------------
	func() {
		fmt.Println("In an anonymous fn")
	}()

	// exercise - 7 ----------------------------------------------
	f := func() {
		fmt.Println("In yet another anonymous fn")
	}
	f()

	// exercise - 8 ----------------------------------------------
	fnf := funcOFunc("shlok from func of func")
	valFromfnf := fnf()
	fmt.Printf("%v\n", valFromfnf)

	// exercise - 9 & 10 ----------------------------------------------
	sumCallback := func(xi []int) int { // this fn encloses the scope of variable sum, v and xi
		sum := 0
		for _, v := range xi {
			sum += v
		}
		return sum
	}

	x := foo2(sumCallback, []int{1, 2, 3, 4, 5, 6})
	fmt.Println(x)
}

/*
	package main

	import "fmt"

	type detective struct {
		name     string
		codeName int
	}

	func (s detective) speak() { // defining a method for any identifier of type detective
		fmt.Printf("Hi, I am %v a.k.a codeName %v \n", s.name, s.codeName)
	}

	type human interface {
		// if a type has method speak() then it is also of type human
		// also everything is of type empty interface ... "interface{}"
		speak() // now any class that has the method speak(), is also of type human

	}

	func humanBurp(h human) {
		fmt.Printf("%v Burps\n", h) // any identifier related to human interface can call this
	}

	func funcOfFunc() func() int { // returns, a function which in itself has a return type of int
		return func() int {
			return 451
		}
	}

	//===========================================================
	func main() {
		foo()
		bar("slok")

		s1 := woo("Moneypenny")
		fmt.Println(s1)

		x, y := mouse("Ian", "Flemming")
		fmt.Println(x)
		fmt.Println(y)

		// variadic parameters----------------
		fmt.Println(fooVariadic(2, 1, 3, 4, 7, 6))

		// unfurling a slice-------------------
		xi := []int{2, 3, 4}
		fmt.Println(fooVariadic(xi...))

		// defer keyword --------------------------
		foo()        // runs first then..
		bar("shlok") // this runs

		defer foo() // runs only after bar(...) is done executing
		bar("shlok")

		// METHODS ---------------------------------
		agent1 := detective{
			name:     "James",
			codeName: 007,
		}
		agent1.speak() // calling method defined for agent1 identifier of type "detective"

		// Interface --------------------------------(interface is also treated as a type in golang)
		// a value can be of more than one type using interfaces
		// interfaces are used for polymorphism in golang
		fmt.Printf("%T\n", agent1) // main.detective
		humanBurp(agent1)          // bcz agent 1 is also of type human as it has method speak()

		// ANONYMOUS function -----------------------------------
		func() { // no function name
			fmt.Println("inside ... Anonymous function\n")
		}()
		func(x int) { // parameters
			fmt.Printf("inside ... Anonymous function with params ... %v\n", x)
		}(42) // argument

		// Function expression -------------------------- in go functions are first class citizens
		f := func() {
			fmt.Println("in a function expression!\n")
		}

		f() // calling func expression assigned to f

		f2 := func(x int) {
			fmt.Printf("in a func exp with params ... %v\n", x)
		}
		f2(2)

		// Returning func from a func
		f3 := funcOfFunc() // -> f3 has an anonymous fn now
		f3Val := f3()
		fmt.Println(f3Val)
		fmt.Println(funcOfFunc()()) //alternate

		// CALLBACKS --------------------- passing func as an argument
		// Closure ------------------------ for variable scopes
	}

	//==================================================================

	// syntax: func (r receiver) identifier(parameters) (returns(s)) {...}
	func foo() {
		fmt.Println("hello from foo")
	}

	// !!!everything in go is  "pass by value"
	func bar(s string) {
		fmt.Println("Hello, ", s)
	}

	func woo(s string) string {
		return fmt.Sprint("Hello from woo, ", s) //  return a string print
	}

	func mouse(fn string, ln string) (string, bool) {
		a := fmt.Sprint(fn, " ", ln, `, says hello!`)
		b := false
		return a, b
	}

	/// variadic parameters------------------

	func fooVariadic(x ...int) int { // "..." is an operator according to go docs used for variadic parameters
		fmt.Println(x)
		fmt.Printf("%T\n", x) // []int

		sum := 0
		for _, val := range x {
			sum += val
		}
		fmt.Println(sum)

		return sum
	}
*/
